package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (c *Circle) area() float64 {

	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {

	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func main() {

	var radius float64
	var rx1, ry1, rx2, ry2 float64
	fmt.Println("Enter the radius of the circle :")
	fmt.Scanf("%f", &radius)
	c := Circle{r: radius}
	fmt.Println("Area of circle ", c.area(), "\n Perimeter of circle : ", c.perimeter())
	fmt.Println("Enter the x co-ordinate of 1st rectangle diagonal :")
	fmt.Scanf("%f", &rx1)
	fmt.Println("Enter the y co-ordinate of 1st rectangle diagonal :")
	fmt.Scanf("%f", &ry1)
	fmt.Println("Enter the x co-ordinate of 2nd rectangle diagonal :")
	fmt.Scanf("%f", &rx2)
	fmt.Println("Enter the y co-ordinate of 2nd rectangle diagonal :")
	fmt.Scanf("%f", &ry2)
	r := Rectangle{x1: rx1, y1: ry1, x2: rx2, y2: ry2}
	fmt.Println("Area of rectangle ", r.area(), "\n Perimeter of rectangle : ", r.perimeter())

}
