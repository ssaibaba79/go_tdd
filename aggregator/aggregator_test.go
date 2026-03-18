package aggregator

import "testing"

func TestAggregator(t *testing.T) {

	data := []int{1, 3, 2, 4, 5, 6, 7}
	t.Run("Test sum", func(t *testing.T) {
		want := 28
		got := Aggregator(SUM, data)
		assertResult(t, got, want)
	})

	t.Run("Test min", func(t *testing.T) {
		want := 1
		got := Aggregator(MIN, data)
		assertResult(t, got, want)
	})

	t.Run("Test max", func(t *testing.T) {
		want := 7
		got := Aggregator(MAX, data)
		assertResult(t, got, want)
	})

	t.Run("Test default", func(t *testing.T) {
		want := 0
		got := Aggregator("", data)
		assertResult(t, got, want)
	})

}

func assertResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%d\" != want \"%d\"", got, want)
	}
}
