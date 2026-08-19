package fuzz

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseRecordEnforcesSizeLimit(t *testing.T) {
	input := []byte(`{"name":"orders","count":2}`)
	if _, err := ParseRecord(input, len(input)-1); !errors.Is(err, ErrInputTooLarge) {
		t.Fatalf("ParseRecord() error = %v, want ErrInputTooLarge", err)
	}
}

// FuzzRecordRoundTrip 验证成功解析的记录可以稳定编码并再次解析。
func FuzzRecordRoundTrip(f *testing.F) {
	f.Add([]byte(`{"name":"orders","count":2}`))
	f.Add([]byte(`{"name":"库存","count":0}`))
	f.Add([]byte(`null`))
	f.Fuzz(func(t *testing.T, input []byte) {
		const maxInputBytes = 4 << 10
		record, err := ParseRecord(input, maxInputBytes)
		if err != nil {
			return
		}
		encoded, err := EncodeRecord(record)
		if err != nil {
			t.Fatalf("EncodeRecord(%+v) error = %v", record, err)
		}
		reparsed, err := ParseRecord(encoded, maxInputBytes)
		if err != nil {
			t.Fatalf("ParseRecord(encoded) error = %v", err)
		}
		if !reflect.DeepEqual(reparsed, record) {
			t.Fatalf("round trip = %+v, want %+v", reparsed, record)
		}
	})
}
