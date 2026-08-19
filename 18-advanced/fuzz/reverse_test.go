package fuzz

import (
	"errors"
	"slices"
	"testing"
	"unicode/utf8"
)

func TestReverseUTF8RejectsInvalidInput(t *testing.T) {
	if _, err := ReverseUTF8(string([]byte{0xff})); !errors.Is(err, ErrInvalidUTF8) {
		t.Fatalf("ReverseUTF8() error = %v, want ErrInvalidUTF8", err)
	}
}

func TestReverseBytesCanBreakUTF8(t *testing.T) {
	input := []byte("世界")
	reversed := ReverseBytes(input)
	if utf8.Valid(reversed) {
		t.Fatalf("ReverseBytes(%q) = %x, want invalid UTF-8 demonstration", input, reversed)
	}
}

func TestReverseUTF8ReversesCodePoints(t *testing.T) {
	got, err := ReverseUTF8("Go世界")
	if err != nil {
		t.Fatal(err)
	}
	if want := "界世oG"; got != want {
		t.Fatalf("ReverseUTF8() = %q, want %q", got, want)
	}
}

// FuzzReverseBytes 验证任意字节序列反转两次后恢复原值。
func FuzzReverseBytes(f *testing.F) {
	f.Add([]byte("hello"))
	f.Add([]byte{0x00, 0xff, 0x7f})
	f.Fuzz(func(t *testing.T, input []byte) {
		got := ReverseBytes(ReverseBytes(input))
		if !slices.Equal(got, input) {
			t.Fatalf("ReverseBytes round trip failed: got %x, want %x", got, input)
		}
	})
}

// FuzzReverseUTF8 验证有效 UTF-8 的可逆性，并确认无效输入被明确拒绝。
func FuzzReverseUTF8(f *testing.F) {
	f.Add("hello")
	f.Add("世界")
	f.Add("")
	f.Fuzz(func(t *testing.T, input string) {
		reversed, err := ReverseUTF8(input)
		if !utf8.ValidString(input) {
			if !errors.Is(err, ErrInvalidUTF8) {
				t.Fatalf("invalid UTF-8 error = %v, want ErrInvalidUTF8", err)
			}
			return
		}
		if err != nil {
			t.Fatalf("ReverseUTF8(%q) error = %v", input, err)
		}
		if !utf8.ValidString(reversed) {
			t.Fatalf("ReverseUTF8(%q) returned invalid UTF-8 %q", input, reversed)
		}
		restored, err := ReverseUTF8(reversed)
		if err != nil {
			t.Fatalf("second ReverseUTF8() error = %v", err)
		}
		if restored != input {
			t.Fatalf("round trip = %q, want %q", restored, input)
		}
	})
}
