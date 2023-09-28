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

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 { //copy of input rectangle is created
	return r.height * r.width
	//copy destroyed
}

func (r *Rectangle) Translate(a, b float64) {
	r.x1 += a
	r.y1 += b
}

func (c *Circle) Translate(a, b float64) {
	c.x1 += a
	c.y1 += b
}

func (c *Circle) Scale(f float64) {
	c.radius *= f
}

func (r *Rectangle) Scale(f float64) {
	r.width *= f
	r.height *= f
}

func CreateNewCircle(x, y, r float64) Circle {
	var c Circle

	c.x1 = x
	c.y1 = y
	c.radius = r
	return c
}

func main() {

	c := CreateNewCircle(-2.1, -3.7, 2.0)

	var pointerToC *Circle
	//at this point, pointerToC has a special value nil
	if pointerToC == nil {
		fmt.Println("hi")
	}

	//point it at c
	pointerToC = &c
	//if you have the key, you can unlock the door and change the fields
	//adding the star is called "dereferencing" the pointer
	//(*pointerToC).x1 = 1.7
	//(*pointerToC).y1 = 12.12

	//in Go, you don't have to dereference pointers.
	//if you have a pointer, just access fields of the object it points to.
	pointerToC.x1 = 1.7
	pointerToC.y1 = 12.12
	fmt.Println("c's center is now", c.x1, c.y1)

	fmt.Println("Let's see what we get when we print the pointer to c:", pointerToC)

	//two things:
	//1. you can have two keys unlocking one door
	//so otherPointer could point to c too
	//2. there is a one line declaration
	otherPointer := &c
	fmt.Println("This other pointer points at c too!", otherPointer.x1, otherPointer.y1)

	/*
		//this is possible (but not needed for this class)
		var x *(*(*(*(int))))
		fmt.Println(x)
	*/

	pointerToC.Translate(-3.1, 2.45)

	fmt.Println(c.x1, c.y1)

	//seeming annoyance: I need to create a pointer to a circle just to move the circle
	//this is another convenience of Go. If you have a circle object,
	//you can use it wherever there is a (c *Circle) Foo(blah) method

	c.Translate(1.4, -14.57)
	fmt.Println(c.x1, c.y1)

	/*
		c = c.Translate(2.1, 3.7)
		fmt.Println(c.Area())
		fmt.Println(c.x1, c.y1)
	*/

	var r Rectangle
	r.width = 3.0
	r.height = 5.0
	fmt.Println(r.Area())

	r.Scale(1.0 / 10.0)
	fmt.Println(r.Area())

	//one more thing
	//keyword "new" that creates an object and gives you a pointer to it
	d := new(Circle) //d has type *Circle
	d.x1 = 1.0
	d.y1 = 2.0
	d.radius = 3.0
	d.Scale(2.5)
	fmt.Println("Area of d is", d.Area())

}

func Pointers() {
	fmt.Println("Pointer stuff.")
	var b int = -14
	var a *int = &b // now a stores the location of b
	// & means "location of"
	fmt.Println("b is located at address:", a)
	fmt.Println("b's value is", b)

}
