package iteration

import (
	"testing"
	"fmt"
)

func Repeat(character string, nums int) string {
	var repeated string
	for i := 0; i < nums; i++ {
		// repeated = repeated + character
		repeated += character
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}

