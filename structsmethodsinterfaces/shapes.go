package structsmethodsinterfaces

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.length * rectangle.breadth
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shapes interface {
	Area() float64
}
