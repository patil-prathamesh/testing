package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 2)
	expected := "aa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
