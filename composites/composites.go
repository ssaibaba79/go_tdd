package composites

type Rectangle struct {
	length float64
	width float64
}

func Area(r Rectangle) float64{
	return r.length * r.width
}

func Perimeter(r Rectangle) float64 {
	return 2 *(r.length + r.width)
}
