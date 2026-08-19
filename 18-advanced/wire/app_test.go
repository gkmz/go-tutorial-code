package wireexample

import (
	"context"
	"errors"
	"testing"
)

func TestInitializeAppGeneratedCode(t *testing.T) {
	app, cleanup, err := InitializeApp(Config{DSN: "memory://wire", GreetingPrefix: "Welcome"})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	message, err := app.Greet(context.Background(), "Ada")
	if err != nil {
		t.Fatal(err)
	}
	if want := "Welcome, Ada!"; message != want {
		t.Fatalf("Greet() = %q, want %q", message, want)
	}
	if got := app.GreetingCount(); got != 1 {
		t.Fatalf("GreetingCount() = %d, want 1", got)
	}
}

func TestGeneratedCleanupClosesRepository(t *testing.T) {
	app, cleanup, err := InitializeApp(Config{DSN: "memory://wire"})
	if err != nil {
		t.Fatal(err)
	}
	cleanup()

	_, err = app.Greet(context.Background(), "Ada")
	if !errors.Is(err, ErrRepositoryClosed) {
		t.Fatalf("Greet() error = %v, want ErrRepositoryClosed", err)
	}
}

func TestGeneratedAndManualInjectorsRejectInvalidConfig(t *testing.T) {
	for name, initialize := range map[string]func(Config) (*App, func(), error){
		"generated": InitializeApp,
		"manual":    InitializeAppManually,
	} {
		t.Run(name, func(t *testing.T) {
			if _, _, err := initialize(Config{}); err == nil {
				t.Fatal("initialize error = nil, want invalid DSN error")
			}
		})
	}
}
