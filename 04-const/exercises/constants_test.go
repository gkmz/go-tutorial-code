package exercises

import "testing"

func TestHTTPStatus(t *testing.T) {
	if StatusOK.String() != "OK" || HTTPStatus(999).String() != "Unknown" {
		t.Fatalf("unexpected status names")
	}
}

func TestFileMode(t *testing.T) {
	mode := ModeRead | ModeWrite
	if !mode.CanRead() || !mode.CanWrite() || mode.CanExecute() {
		t.Fatalf("unexpected permissions for %d", mode)
	}
}

func TestColor(t *testing.T) {
	if Unknown.IsValid() || Unknown.String() != "Unknown" {
		t.Fatal("unknown color should be invalid")
	}
	color, err := FromString("Green")
	if err != nil || color != Green || !color.IsValid() {
		t.Fatalf("unexpected color: %v, %v", color, err)
	}
	if _, err := FromString("Purple"); err == nil {
		t.Fatal("expected invalid color error")
	}
}

func TestOrderStatus(t *testing.T) {
	if !OrderCreated.CanTransitionTo(OrderPaid) {
		t.Fatal("created order should be payable")
	}
	if OrderCancelled.CanTransitionTo(OrderPaid) {
		t.Fatal("cancelled order must not be payable")
	}
}
