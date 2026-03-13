package composites

import "math"

type Rectangle struct {
	length float64
	width float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) Area() float64{
	return math.Floor(r.length * r.width)
}

func (r Rectangle) Perimeter() float64 {
	return math.Floor(2 *(r.length + r.width))
}

func (c Circle) Area() float64 {
	return math.Floor(math.Pi * (c.radius * c.radius))
}

func (c Circle) Perimeter() float64 {
	return math.Floor(2 * math.Pi * c.radius)
}
