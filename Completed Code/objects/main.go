package main

import "fmt"

type Rectangle struct {
	x1, y1        float64
	width, height float64
	rotation      float64
}

type Circle struct {
	x1, y1 float64 //center
	radius float64
}

func Area(r Rectangle) float64 { //copy of input rectangle is created
	r.height = 0.0
	r.width = 0.0
	return r.height * r.width
	//copy destroyed
}

func CreateNewCircle(x, y, r float64) Circle {
	var c Circle

	c.x1 = x
	c.y1 = y
	c.radius = r
	return c
}

func main() {
	var r Rectangle
	r.width = 3.0
	r.height = 5.0
	fmt.Println(Area(r))
	fmt.Println(r.width, r.height)
}
