// Section 09 - Lecture 01 : What is an interface?
package main

type (
	Shape interface {
		Area() float64
	}
	Circle interface {
		Radius() float64
		Shape
	}
	Rect interface {
		Width() float64
		Length() float64
		Shape
	}
)

func main() {
	// interface embedding
	// ----

}
