package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func demonstrateTeeReader() error {
	source := strings.NewReader("Go IO pipeline")
	var copied bytes.Buffer

	// TeeReader 只会复制真正被下游读取的数据。
	tee := io.TeeReader(source, &copied)
	data, err := io.ReadAll(tee)
	if err != nil {
		return fmt.Errorf("read tee: %w", err)
	}
	fmt.Printf("read=%q copied=%q\n", data, copied.String())
	return nil
}

func demonstrateMultiReader() error {
	combined := io.MultiReader(
		strings.NewReader("header\n"),
		strings.NewReader("body\n"),
	)
	if _, err := io.Copy(io.Discard, combined); err != nil {
		return fmt.Errorf("copy combined reader: %w", err)
	}
	return nil
}
