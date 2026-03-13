package composites

import "testing"

func TestArea(t *testing.T) {
  want := 50.0
	got := Rectangle{length:10.0, width:5.0}.Area()
  assertResult(t, got, want)
	
}

func TestPerimeter(t *testing.T) {
	want := 30.0
	got := Rectangle{length:10.0, width:5.0}.Perimeter()
  assertResult(t, got, want)
	
}

func TestAreaCircle(t *testing.T) {
  want := 314.00
	got := Circle{radius:10}.Area()
  assertResult(t, got, want)
	
}

func TestPerimeterCircle(t *testing.T) {
	want := 62.00
	got := Circle{radius:10}.Perimeter()
  assertResult(t, got, want)
	
}

func assertResult(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%.2f\" != want \"%.2f\"", got, want)
	}
}