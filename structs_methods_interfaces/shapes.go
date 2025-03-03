package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

//rectangle methods
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//perimeter methods
func (c Circle) Perimeter() float64 {
	return c.Radius * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//triangle methods
func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height + t.Height
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
