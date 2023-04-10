package lessons2

import "fmt"

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func init() {
	r := Rectangle{width: 10.378912, height: 5.127980}
	fmt.Println("Area of rectangle: ", area(r))
}

// In this example we have a Rectangle struct with a width and height. 
// We also have a function area() that takes a Rectangle as an argument and returns the area of the rectangle.

// But there is one problem, area function is not a method of Reactangle struct.
// What if we want to also count area of triangle, circle, etc.?
// We will have to create a function for each shape.
// Thats why we have methods.

// @Methods
// @Description: Methods are functions with a special receiver argument. The receiver appears in its own argument
// list between the func keyword and the method name.
// Using the same example, Rectangle.Area() belongs directly to rectangle, instead of as a peripheral function. More
// specifically, length , width and Area() all belong to rectangle.

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func init() {
	circle1 := Circle{radius: 5.127980};
	circle2 := Circle{radius: 10.378912};
	rectangle1 := Rectangle{width: 10.378912, height: 5.127980};
	rectangle2 := Rectangle{width: 5.127980, height: 10.378912};

	fmt.Println("Area of circle1: ", circle1.Area());
	fmt.Println("Area of circle2: ", circle2.Area());
	fmt.Println("Area of rectangle1: ", rectangle1.Area());
	fmt.Println("Area of rectangle2: ", rectangle2.Area());
}
