package composites

import "testing"

func TestArea(t *testing.T) {
  want := 50.0
	got := Area(Rectangle{length:10.0, width:5.0})
  assertResult(t, got, want)
	
}

func TestPerimeter(t *testing.T) {
	want := 30.0
	got := Perimeter(Rectangle{length:10.0, width:5.0})
  assertResult(t, got, want)
	
}


func assertResult(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%.2f\" != want \"%.2f\"", got, want)
	}
}