package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountDown(t *testing.T) {

	buff := &bytes.Buffer{}
	want := fmt.Sprintf("%s\n%s\n%s\n%s", "3","2","1","Go!")
	sleeper := MockingSleeper{}
	CountDown(buff, sleeper)
	got := buff.String()
  assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want:%s, got:%s", want, got)
	}
}

type MockingSleeper struct {}

func (d MockingSleeper ) Sleep(){}


