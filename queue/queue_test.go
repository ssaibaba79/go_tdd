package queue

import (
	"sync"
	"testing"
)

func TestQueueOperations(t *testing.T) {

	t.Run("Test enqueue", func(t *testing.T) {

		n := 10
		queue := NewQueue()

		for i := 1; i <= n; i++ {
			queue.Enqueue(i)
		}

		got := queue.Size()
		assertResult(t, got, n)
	})

	t.Run("Test Dequeue", func(t *testing.T) {

		n := 10
		queue := NewQueue()

		for i := 1; i <= n; i++ {
			queue.Enqueue(i)
		}
		v := queue.Dequeue()
		assertResult(t, v, 1)

		got := queue.Size()
		assertResult(t, got, n-1)
	})

	t.Run("Test Concurrent enqueue", func(t *testing.T) {

		n := 1000
		queue := NewQueue()

		var wg sync.WaitGroup
		for i := 1; i <= n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				queue.Enqueue(i)
			}()
		}

		wg.Wait()

		got := queue.Size()
		assertResult(t, got, n)
	})

	t.Run("Test Concurrent Dequeue", func(t *testing.T) {

		n := 1000
		queue := NewQueue()

		var wg sync.WaitGroup
		for i := 1; i <= n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				queue.Enqueue(i)
			}()
		}

		wg.Wait()

		for i := 1; i <= n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				queue.Dequeue()
			}()
		}

		wg.Wait()

		got := queue.Size()
		assertResult(t, got, 0)
	})

}

func assertResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%d\" != want \"%d\"", got, want)
	}
}
