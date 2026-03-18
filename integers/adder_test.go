package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	want := 5
	got := Add(3, 2)
	if got != want {
		t.Errorf("Add failed want: %d got: %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
