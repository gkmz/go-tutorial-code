package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadEnvironmentOverridesFile(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "config.yaml")
	content := []byte("server:\n  host: file-host\n  port: 8080\ndatabase:\n  url: postgres://file\nlog:\n  level: info\n")
	if err := os.WriteFile(filename, content, 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("APP_SERVER_PORT", "9090")
	got, err := Load(filename)
	if err != nil {
		t.Fatal(err)
	}
	if got.Server.Port != 9090 || got.Server.Host != "file-host" {
		t.Fatalf("config = %+v", got)
	}
}

func TestLoadRejectsInvalidConfig(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(filename, []byte("server:\n  port: 0\ndatabase:\n  url: postgres://db\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(filename); err == nil {
		t.Fatal("Load() error = nil, want validation error")
	}
}

func TestStoreKeepsPreviousSnapshotOnFailedReplace(t *testing.T) {
	initial := Config{Server: ServerConfig{Host: "127.0.0.1", Port: 8080}, Database: DatabaseConfig{URL: "postgres://db"}, Log: LogConfig{Level: "info"}}
	store, err := NewStore(initial)
	if err != nil {
		t.Fatal(err)
	}
	invalid := initial
	invalid.Server.Port = 0
	if err := store.Replace(invalid); err == nil {
		t.Fatal("Replace() error = nil, want validation error")
	}
	if got := store.Snapshot(); got.Server.Port != 8080 {
		t.Fatalf("snapshot port = %d, want 8080", got.Server.Port)
	}
}

func TestLoadDotEnvIfPresent(t *testing.T) {
	filename := filepath.Join(t.TempDir(), ".env")
	if err := os.WriteFile(filename, []byte("APP_DATABASE_URL=postgres://dotenv\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	previous, existed := os.LookupEnv("APP_DATABASE_URL")
	if err := os.Unsetenv("APP_DATABASE_URL"); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if existed {
			_ = os.Setenv("APP_DATABASE_URL", previous)
		} else {
			_ = os.Unsetenv("APP_DATABASE_URL")
		}
	})
	if err := LoadDotEnvIfPresent(filename); err != nil {
		t.Fatal(err)
	}
	if got := os.Getenv("APP_DATABASE_URL"); got != "postgres://dotenv" {
		t.Fatalf("APP_DATABASE_URL = %q", got)
	}
}

func TestLoadWithOverridesUsesExplicitValueFirst(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "config.yaml")
	content := []byte("server:\n  port: 8080\ndatabase:\n  url: postgres://file\nlog:\n  level: info\n")
	if err := os.WriteFile(filename, content, 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("APP_SERVER_PORT", "9090")
	got, err := LoadWithOverrides(filename, map[string]any{"server.port": 10000})
	if err != nil {
		t.Fatal(err)
	}
	if got.Server.Port != 10000 {
		t.Fatalf("server port = %d, want 10000", got.Server.Port)
	}
}

func TestRedactedSummaryDoesNotExposeCredentials(t *testing.T) {
	value := Config{
		Server:   ServerConfig{Host: "127.0.0.1", Port: 8080},
		Database: DatabaseConfig{URL: "postgres://admin:secret@db.example.com:5432/orders"},
		Log:      LogConfig{Level: "info"},
	}
	summary := value.RedactedSummary()
	if strings.Contains(summary, "admin") || strings.Contains(summary, "secret") {
		t.Fatalf("summary exposes credentials: %q", summary)
	}
	if !strings.Contains(summary, "db.example.com") || !strings.Contains(summary, "orders") {
		t.Fatalf("summary misses safe fields: %q", summary)
	}
}
