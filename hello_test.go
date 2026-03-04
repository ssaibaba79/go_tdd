package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Sarvan"
		got := Hello("Sarvan")
		assertResult(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		want := "Hello, World"
		got := Hellfiro("")
		assertResult(t, got, want)
	})
}

func assertResult(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
			t.Errorf("got \"%s\" != want \"%s\"", got, want)
		}
}
