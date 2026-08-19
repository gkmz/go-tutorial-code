package config

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

const defaultWatchDebounce = 150 * time.Millisecond

// WatchOptions 描述配置文件监听的防抖和错误处理行为。
type WatchOptions struct {
	// Debounce 合并编辑器连续产生的文件事件；零值使用默认时长。
	Debounce time.Duration
	// OnError 接收读取或校验失败；回调返回后继续保留旧配置并监听后续变化。
	OnError func(error)
}

// WatchFile 监听配置文件变化，并在新配置校验通过后原子替换 Store 快照。
//
// 函数会阻塞到 ctx 被取消。为兼容编辑器通过临时文件和 Rename 保存配置的方式，
// 实现监听文件所在目录，而不是只监听当前文件句柄。
func WatchFile(ctx context.Context, filename string, store *Store, options WatchOptions) error {
	if ctx == nil {
		return errors.New("watch context must not be nil")
	}
	if store == nil {
		return errors.New("config store must not be nil")
	}
	if filename == "" {
		return errors.New("config filename must not be empty")
	}

	absoluteFilename, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("resolve config filename: %w", err)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create config watcher: %w", err)
	}
	defer watcher.Close()

	if err := watcher.Add(filepath.Dir(absoluteFilename)); err != nil {
		return fmt.Errorf("watch config directory: %w", err)
	}

	debounce := options.Debounce
	if debounce <= 0 {
		debounce = defaultWatchDebounce
	}

	var timer *time.Timer
	var timerChannel <-chan time.Time
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case watchErr, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			reportWatchError(options.OnError, fmt.Errorf("watch config: %w", watchErr))
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if !isConfigChange(event, absoluteFilename) {
				continue
			}
			// 一次保存可能产生多次 Write、Create 或 Rename，防抖后只重新加载一次。
			if timer == nil {
				timer = time.NewTimer(debounce)
			} else {
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(debounce)
			}
			timerChannel = timer.C
		case <-timerChannel:
			timerChannel = nil
			next, loadErr := Load(absoluteFilename)
			if loadErr != nil {
				reportWatchError(options.OnError, loadErr)
				continue
			}
			if replaceErr := store.Replace(next); replaceErr != nil {
				reportWatchError(options.OnError, replaceErr)
			}
		}
	}
}

func isConfigChange(event fsnotify.Event, filename string) bool {
	if filepath.Clean(event.Name) != filepath.Clean(filename) {
		return false
	}
	return event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) != 0
}

func reportWatchError(handler func(error), err error) {
	if handler != nil {
		handler(err)
	}
}
