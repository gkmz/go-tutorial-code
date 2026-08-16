package exercises

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCopyAndLowercaseTee(t *testing.T) {
	dir := t.TempDir()
	sourcePath := filepath.Join(dir, "source.txt")
	targetPath := filepath.Join(dir, "target.txt")
	if err := os.WriteFile(sourcePath, []byte("Hello IO"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := CopyFile(sourcePath, targetPath); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(targetPath)
	if err != nil || string(data) != "Hello IO" {
		t.Fatalf("copied data = %q, err = %v", data, err)
	}
	var target, output bytes.Buffer
	if err := LowercaseTee(bytes.NewBufferString("ABC"), &target, &output); err != nil {
		t.Fatal(err)
	}
	if target.String() != "abc" || output.String() != "abc" {
		t.Fatalf("target/output = %q/%q", target.String(), output.String())
	}
}

func TestCountWords(t *testing.T) {
	got, err := CountWords(bytes.NewBufferString("go io go"))
	if err != nil || !reflect.DeepEqual(got, map[string]int{"go": 2, "io": 1}) {
		t.Fatalf("counts = %v, err = %v", got, err)
	}
}
