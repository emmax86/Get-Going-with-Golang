package main

import "math"

type Shape interface {
	Area() float64
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * 2 * math.Pi
}
