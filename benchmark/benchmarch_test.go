package benchmark

import "testing"

func TestRepeat(t *testing.T) {
	want := "aaaa"
	got := Repeat("a", 4)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 3)
	}
}

func BenchmarkRepeatOptimized(b *testing.B) {
	for b.Loop() {
		RepeatOptimized("a", 3)
	}
}
