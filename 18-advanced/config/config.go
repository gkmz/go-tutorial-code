// Package config 提供文件、环境变量和 .env 辅助加载的强类型配置示例。
package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"sync/atomic"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config 是应用运行所需的、经过校验的配置快照。
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
}

// ServerConfig 描述 HTTP 服务配置。
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// DatabaseConfig 描述数据库连接配置。
type DatabaseConfig struct {
	URL string `mapstructure:"url"`
}

// LogConfig 描述日志级别配置。
type LogConfig struct {
	Level string `mapstructure:"level"`
}

// Validate 检查配置是否满足业务启动条件。
func (c Config) Validate() error {
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return fmt.Errorf("server.port must be between 1 and 65535, got %d", c.Server.Port)
	}
	if strings.TrimSpace(c.Server.Host) == "" {
		return errors.New("server.host must not be empty")
	}
	if strings.TrimSpace(c.Database.URL) == "" {
		return errors.New("database.url must not be empty")
	}
	switch strings.ToLower(c.Log.Level) {
	case "debug", "info", "warn", "error":
		return nil
	default:
		return fmt.Errorf("log.level must be debug, info, warn or error, got %q", c.Log.Level)
	}
}

// Load 从 YAML 文件和 APP_ 前缀环境变量加载配置。
// 环境变量优先于配置文件，配置文件优先于默认值。
func Load(filename string) (Config, error) {
	return LoadWithOverrides(filename, nil)
}

// LoadWithOverrides 加载配置，并让显式覆盖值拥有最高优先级。
// overrides 用于承接已经解析和校验过类型的命令行参数。
func LoadWithOverrides(filename string, overrides map[string]any) (Config, error) {
	reader := viper.New()
	reader.SetConfigType("yaml")
	reader.SetEnvPrefix("APP")
	reader.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	reader.AutomaticEnv()
	setDefaults(reader)
	bindEnvironment(reader)
	for key, value := range overrides {
		reader.Set(key, value)
	}

	if filename != "" {
		reader.SetConfigFile(filename)
		if err := reader.ReadInConfig(); err != nil {
			return Config{}, fmt.Errorf("read config: %w", err)
		}
	}

	var result Config
	if err := reader.Unmarshal(&result); err != nil {
		return Config{}, fmt.Errorf("decode config: %w", err)
	}
	if err := result.Validate(); err != nil {
		return Config{}, fmt.Errorf("validate config: %w", err)
	}
	return result, nil
}

// RedactedSummary 返回不包含数据库用户名和密码的配置摘要。
func (c Config) RedactedSummary() string {
	databaseHost := "invalid"
	databaseName := "unknown"
	if parsed, err := url.Parse(c.Database.URL); err == nil {
		if parsed.Hostname() != "" {
			databaseHost = parsed.Hostname()
		}
		if name := path.Base(parsed.Path); name != "." && name != "/" && name != "" {
			databaseName = name
		}
	}
	return fmt.Sprintf("server=%s:%d database_host=%s database_name=%s log_level=%s",
		c.Server.Host, c.Server.Port, databaseHost, databaseName, c.Log.Level)
}

// LoadDotEnvIfPresent 将本地 .env 文件加载到当前进程环境。
// 生产部署应使用平台环境变量或密钥管理服务，而不是依赖 .env 文件。
func LoadDotEnvIfPresent(filename string) error {
	if filename == "" {
		return errors.New("dotenv filename must not be empty")
	}
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return fmt.Errorf("stat dotenv: %w", err)
	}
	if err := godotenv.Load(filename); err != nil {
		return fmt.Errorf("load dotenv: %w", err)
	}
	return nil
}

// Store 保存可以被请求并发读取的配置快照。
type Store struct {
	value atomic.Value
}

// NewStore 创建配置存储，并拒绝未通过校验的初始值。
func NewStore(initial Config) (*Store, error) {
	if err := initial.Validate(); err != nil {
		return nil, err
	}
	store := &Store{}
	store.value.Store(initial)
	return store, nil
}

// Snapshot 返回当前完整配置快照。
func (s *Store) Snapshot() Config {
	return s.value.Load().(Config)
}

// Replace 校验并原子替换配置，失败时保留旧快照。
func (s *Store) Replace(next Config) error {
	if err := next.Validate(); err != nil {
		return err
	}
	s.value.Store(next)
	return nil
}

func setDefaults(reader *viper.Viper) {
	reader.SetDefault("server.host", "127.0.0.1")
	reader.SetDefault("server.port", 8080)
	reader.SetDefault("log.level", "info")
}

func bindEnvironment(reader *viper.Viper) {
	for _, key := range []string{
		"server.host",
		"server.port",
		"database.url",
		"log.level",
	} {
		_ = reader.BindEnv(key)
	}
}
