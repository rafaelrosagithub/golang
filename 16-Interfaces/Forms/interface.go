package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

func writeArea(f form) {
	fmt.Printf("The shape area is %0.2f\n", f.area())
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * (c.ray * c.ray)
}

type circle struct {
	ray float64
}

func main() {
	fmt.Println("------------- Interfaces ------------")

	r := rectangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)

}
