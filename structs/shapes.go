package structs

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}
