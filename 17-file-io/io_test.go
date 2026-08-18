package main

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadSmallFileRejectsOversizedInput(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.txt")
	if err := os.WriteFile(path, []byte("12345"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := readSmallFile(path, 4); err == nil {
		t.Fatal("readSmallFile() error = nil, want size error")
	}
}

func TestScanLines(t *testing.T) {
	var lines []string
	err := scanLines(bytes.NewBufferString("first\nsecond\n"), 1024, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"first", "second"}; !reflect.DeepEqual(lines, want) {
		t.Fatalf("lines = %v, want %v", lines, want)
	}
}

func TestCopyFile(t *testing.T) {
	dir := t.TempDir()
	source := filepath.Join(dir, "source.bin")
	target := filepath.Join(dir, "target.bin")
	if err := os.WriteFile(source, []byte("streamed data"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := copyFile(source, target); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "streamed data" {
		t.Fatalf("target = %q, want streamed data", got)
	}
}
