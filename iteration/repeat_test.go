package iteration

import "testing"

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		// repeated = repeated + character
		repeated += character
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

