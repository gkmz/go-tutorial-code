package exercises

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"testing/fstest"
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

func TestCountLongWordsAndFS(t *testing.T) {
	word := bytes.Repeat([]byte{'x'}, 2048)
	if _, err := CountLongWords(bytes.NewReader(word), 1024); err == nil {
		t.Fatal("expected scanner token limit error")
	}
	counts, err := CountLongWords(bytes.NewReader(word), 4096)
	if err != nil || counts[string(word)] != 1 {
		t.Fatalf("long word counts = %v, err = %v", counts, err)
	}
	data, err := ReadFSFile(fstest.MapFS{"config.txt": &fstest.MapFile{Data: []byte("ok")}}, "config.txt")
	if err != nil || string(data) != "ok" {
		t.Fatalf("ReadFSFile() = %q, err = %v", data, err)
	}
}

func TestLowercaseTeePreservesNonASCIIBytes(t *testing.T) {
	var target, output bytes.Buffer
	if err := LowercaseTee(bytes.NewBufferString("GO语言"), &target, &output); err != nil {
		t.Fatal(err)
	}
	if target.String() != "go语言" || output.String() != "go语言" {
		t.Fatalf("target/output = %q/%q", target.String(), output.String())
	}
}
