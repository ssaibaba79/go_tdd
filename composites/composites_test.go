package composites

import "testing"

type Test struct {
	name  string
	shape Shape
	want  float64
}

func TestArea(t *testing.T) {

	areaTests := []Test{
		{"Rectangle", Rectangle{length: 10.0, width: 5.0}, 50.0},
		{"Circle", Circle{radius: 10}, 314.00},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got=%g want=%g", tt.shape, got, tt.want)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {

	perimeterTests := []Test{
		{"Rectangle", Rectangle{length: 10.0, width: 5.0}, 30.0},
		{"Circle", Circle{radius: 10}, 62.00},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("%#v got=%g want=%g", tt.shape, got, tt.want)
			}
		})
	}
}
