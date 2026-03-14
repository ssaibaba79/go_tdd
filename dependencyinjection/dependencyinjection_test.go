package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	want := "Hello, Saravanan"
	buff := bytes.Buffer{}
	Greet(&buff, "Saravanan")
	got := buff.String()

	if got != want {
		t.Errorf("want : %s, got %s", want, got)
	}

}