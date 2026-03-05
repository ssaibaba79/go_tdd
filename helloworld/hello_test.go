package helloworld

import "testing"

func TestHello(t *testing.T) {

	t.Run("greet in english", func(t *testing.T) {
		want := "Hello, Sarvan"
		got := Hello("Sarvan", "English")
		assertResult(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "English")
		assertResult(t, got, want)
	})

	t.Run("greet in spanish", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", "Spanish")
		assertResult(t, got, want)
	})

	t.Run("greet in french", func(t *testing.T) {
		want := "Bonjour, Marsallie"
		got := Hello("Marsallie", "French")
		assertResult(t, got, want)
	})
}

func assertResult(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%s\" != want \"%s\"", got, want)
	}
}
