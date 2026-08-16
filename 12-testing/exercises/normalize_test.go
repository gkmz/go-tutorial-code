package exercises

import "testing"

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "spaces", input: "  Go   test ", want: "Go test"},
		{name: "unicode", input: " Go\t语言 ", want: "Go 语言"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Normalize(test.input); got != test.want {
				t.Fatalf("Normalize(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// 本模块没有外部资源；保留 TestMain 作为生命周期练习答案。
	code := m.Run()
	if code != 0 {
		panic("tests failed")
	}
}

func BenchmarkConcatPlus(b *testing.B) {
	values := []string{"Go", " testing", " benchmark"}
	for i := 0; i < b.N; i++ {
		_ = ConcatPlus(values)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	values := []string{"Go", " testing", " benchmark"}
	for i := 0; i < b.N; i++ {
		_ = ConcatBuilder(values)
	}
}
