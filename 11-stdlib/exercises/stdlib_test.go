package exercises

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"math"
	"slices"
	"strings"
	"testing"
	"time"
)

func TestNormalizeTextAndParseNumbers(t *testing.T) {
	t.Parallel()

	if got := NormalizeText("  Go\tstdlib\npractice "); got != "Go stdlib practice" {
		t.Fatalf("NormalizeText() = %q", got)
	}

	numbers, err := ParseNumbers("9223372036854775807", "3.1415")
	if err != nil {
		t.Fatalf("ParseNumbers() error = %v", err)
	}
	if numbers.Integer != math.MaxInt64 || numbers.Decimal != 3.1415 {
		t.Fatalf("ParseNumbers() = %+v", numbers)
	}

	for _, input := range []struct {
		integer string
		decimal string
	}{
		{integer: "9223372036854775808", decimal: "1.0"},
		{integer: "42", decimal: "not-a-number"},
	} {
		if _, err := ParseNumbers(input.integer, input.decimal); err == nil {
			t.Fatalf("ParseNumbers(%q, %q) should fail", input.integer, input.decimal)
		}
	}
}

func TestParseIntervalInLocation(t *testing.T) {
	t.Parallel()

	shanghai, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatalf("LoadLocation() error = %v", err)
	}
	duration, err := ParseIntervalInLocation(
		"2006-01-02 15:04",
		"2026-08-19 09:30",
		"2026-08-19 11:00",
		shanghai,
	)
	if err != nil || duration != 90*time.Minute {
		t.Fatalf("ParseIntervalInLocation() = %v, %v", duration, err)
	}

	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("LoadLocation() error = %v", err)
	}
	// 2026 年春季跳过 02:00 到 02:59，墙上时间相差两小时，实际只过去一小时。
	duration, err = ParseIntervalInLocation(
		"2006-01-02 15:04",
		"2026-03-08 01:30",
		"2026-03-08 03:30",
		newYork,
	)
	if err != nil || duration != time.Hour {
		t.Fatalf("DST duration = %v, %v", duration, err)
	}
}

func TestWaitForResult(t *testing.T) {
	t.Parallel()

	result := make(chan int, 1)
	result <- 42
	if got, err := WaitForResult(context.Background(), result, time.Second); err != nil || got != 42 {
		t.Fatalf("result first: got %d, err = %v", got, err)
	}

	if _, err := WaitForResult(context.Background(), make(chan int), time.Millisecond); !errors.Is(err, ErrWaitTimeout) {
		t.Fatalf("timeout error = %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := WaitForResult(ctx, make(chan int), time.Second); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancel error = %v", err)
	}
}

func TestJSONEncodingAndStrictDecoding(t *testing.T) {
	t.Parallel()

	event := Event{ID: 1, Name: "Go", CreatedAt: time.Unix(0, 0).UTC(), Secret: "hidden"}
	data, err := EncodeEvent(event)
	if err != nil {
		t.Fatalf("EncodeEvent() error = %v", err)
	}
	if bytes.Contains(data, []byte("hidden")) || bytes.Contains(data, []byte("description")) {
		t.Fatalf("EncodeEvent() = %s", data)
	}

	decoded, err := DecodeEventStrict(data)
	if err != nil || decoded.ID != event.ID || decoded.Name != event.Name {
		t.Fatalf("DecodeEventStrict() = %+v, %v", decoded, err)
	}
	for _, invalid := range [][]byte{
		[]byte(`{"id":1,"name":"Go","created_at":"1970-01-01T00:00:00Z","unknown":true}`),
		[]byte(`{"id":1,"name":"Go","created_at":"1970-01-01T00:00:00Z"} {"id":2}`),
	} {
		if _, err := DecodeEventStrict(invalid); err == nil {
			t.Fatalf("DecodeEventStrict(%s) should fail", invalid)
		}
	}
}

func TestStableSortAndBinarySearch(t *testing.T) {
	t.Parallel()

	records := []Record{
		{Key: 2, Sequence: 1},
		{Key: 1, Sequence: 2},
		{Key: 2, Sequence: 3},
	}
	sorted := StableSortRecords(records)
	want := []Record{
		{Key: 1, Sequence: 2},
		{Key: 2, Sequence: 1},
		{Key: 2, Sequence: 3},
	}
	if !slices.Equal(sorted, want) {
		t.Fatalf("StableSortRecords() = %+v", sorted)
	}
	if !slices.Equal(records, []Record{{Key: 2, Sequence: 1}, {Key: 1, Sequence: 2}, {Key: 2, Sequence: 3}}) {
		t.Fatalf("StableSortRecords() modified input: %+v", records)
	}
	if index := FirstRecordIndex(sorted, 2); index != 1 {
		t.Fatalf("FirstRecordIndex() = %d", index)
	}

	values := []int{3, 1, 2, 2}
	copyValues, index := SortedCopySearch(values, 2)
	if !slices.Equal(copyValues, []int{1, 2, 2, 3}) || index != 1 {
		t.Fatalf("SortedCopySearch() = %v, %d", copyValues, index)
	}
	if !slices.Equal(values, []int{3, 1, 2, 2}) {
		t.Fatalf("SortedCopySearch() modified input: %v", values)
	}
}

func TestRandomFileAndSecureToken(t *testing.T) {
	t.Parallel()

	if left, right := RandomSequence(42, 5), RandomSequence(42, 5); !slices.Equal(left, right) {
		t.Fatalf("fixed seed sequences differ: %v, %v", left, right)
	}

	content := []byte("standard library")
	readBack, err := TempFileRoundTrip(t.TempDir(), "nested.txt", content)
	if err != nil || !bytes.Equal(readBack, content) {
		t.Fatalf("TempFileRoundTrip() = %q, %v", readBack, err)
	}

	token, err := SecureToken(32)
	if err != nil {
		t.Fatalf("SecureToken() error = %v", err)
	}
	decoded, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil || len(decoded) != 32 {
		t.Fatalf("SecureToken() = %q, decode error = %v", token, err)
	}
	if strings.ContainsAny(token, "+/=") {
		t.Fatalf("SecureToken() is not raw URL-safe base64: %q", token)
	}
}
