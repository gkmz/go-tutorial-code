package exercises

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

var (
	packageResourceDir string
	benchmarkResult    string
)

func TestMain(m *testing.M) {
	directory, err := os.MkdirTemp("", "go-testing-exercise-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, "创建包级测试资源:", err)
		os.Exit(1)
	}
	packageResourceDir = directory

	code := m.Run()
	if err := os.RemoveAll(directory); err != nil {
		fmt.Fprintln(os.Stderr, "清理包级测试资源:", err)
		if code == 0 {
			code = 1
		}
	}
	os.Exit(code)
}

func TestNormalize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "spaces", input: "  Go   test ", want: "Go test"},
		{name: "unicode whitespace", input: " Go\u3000语言\t教程 ", want: "Go 语言 教程"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := Normalize(test.input); got != test.want {
				t.Fatalf("Normalize(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func TestPackageResource(t *testing.T) {
	path := filepath.Join(packageResourceDir, "ready")
	if err := os.WriteFile(path, []byte("ready"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("Stat() error = %v", err)
	}
}

type memoryRepository struct {
	names map[int]string
	err   error
}

func (repository memoryRepository) GetName(id int) (string, error) {
	if repository.err != nil {
		return "", repository.err
	}
	return repository.names[id], nil
}

func TestGreeting(t *testing.T) {
	t.Parallel()

	greeting, err := Greeting(memoryRepository{names: map[int]string{1: "Go"}}, 1)
	if err != nil || greeting != "Hello, Go" {
		t.Fatalf("Greeting() = %q, %v", greeting, err)
	}

	dependencyError := errors.New("database unavailable")
	if _, err := Greeting(memoryRepository{err: dependencyError}, 1); !errors.Is(err, dependencyError) {
		t.Fatalf("Greeting() error = %v", err)
	}
}

func TestCounterConcurrent(t *testing.T) {
	t.Parallel()

	var counter Counter
	var waitGroup sync.WaitGroup
	const goroutines = 8
	const increments = 1000

	for range goroutines {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for range increments {
				counter.Add(1)
			}
		}()
	}
	waitGroup.Wait()

	if got, want := counter.Value(), goroutines*increments; got != want {
		t.Fatalf("Counter.Value() = %d, want %d", got, want)
	}
}

func TestExpiredWithFixedClock(t *testing.T) {
	t.Parallel()

	fixed := time.Date(2026, time.August, 19, 10, 0, 0, 0, time.UTC)
	now := func() time.Time { return fixed }

	if Expired(now, fixed.Add(time.Second)) {
		t.Fatal("deadline in the future should not be expired")
	}
	if !Expired(now, fixed) {
		t.Fatal("deadline equal to now should be expired")
	}
}

func TestParallelFileIsolation(t *testing.T) {
	t.Parallel()

	for _, content := range []string{"first", "second", "third"} {
		content := content
		t.Run(content, func(t *testing.T) {
			t.Parallel()

			// 每个子测试拥有独立目录，因此可以使用相同文件名。
			path := filepath.Join(t.TempDir(), "result.txt")
			if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
				t.Fatalf("WriteFile() error = %v", err)
			}
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("ReadFile() error = %v", err)
			}
			if string(data) != content {
				t.Fatalf("file content = %q, want %q", data, content)
			}
		})
	}
}

func BenchmarkConcatPlus(b *testing.B) {
	values := []string{"Go", " testing", " benchmark"}
	b.ReportAllocs()
	for b.Loop() {
		benchmarkResult = ConcatPlus(values)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	values := []string{"Go", " testing", " benchmark"}
	b.ReportAllocs()
	for b.Loop() {
		benchmarkResult = ConcatBuilder(values)
	}
}
