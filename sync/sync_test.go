package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("test counter sync", func(t *testing.T) {
		wantedIncrement := 100
		counter := NewCounter()

		for i := 0; i < wantedIncrement; i++ {
			counter.Inc()
		}

		assertCounter(t, counter, wantedIncrement)

	})

	t.Run("test concurrent counter ", func(t *testing.T) {
		wantedIncrement := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedIncrement)

		for i := 0; i < wantedIncrement; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedIncrement)

	})

}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
