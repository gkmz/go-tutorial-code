package initialization

import "testing"

func TestExplicitConstructorIsConfigurableAndTestable(t *testing.T) {
	service, err := NewService("test")
	if err != nil {
		t.Fatal(err)
	}
	if got := service.Name(); got != "test" {
		t.Fatalf("Name() = %q, want test", got)
	}
	if _, err := NewService(""); err == nil {
		t.Fatal("NewService() error = nil, want validation error")
	}
	if got := ImplicitValue(); got != "implicit" {
		t.Fatalf("ImplicitValue() = %q, want implicit", got)
	}
}
