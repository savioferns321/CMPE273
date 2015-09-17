package testCases

import (
	"math"
	"testing"
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

/*
func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{x1: 0, x2: 5, y1: 0, y2: 3}
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())

}*/

type testRectPair struct {
	inputVal  Rectangle
	outputVal float64
}

type testCirclePair struct {
	inputVal  Circle
	outputVal float64
}

var testRectPerimeterCases = []testRectPair{
	{Rectangle{x1: 0, y1: 0, x2: 5, y2: 8}, 26},
	{Rectangle{x1: 10, y1: 10, x2: 45, y2: 26}, 102},
	{Rectangle{x1: -9, y1: -11, x2: 27, y2: 21}, 136},
	{Rectangle{x1: -10, y1: 32, x2: 27, y2: -12}, 162},
}

var testRectAreaCases = []testRectPair{
	{Rectangle{x1: 0, y1: 0, x2: 5, y2: 8}, 40},
	{Rectangle{x1: 10, y1: 10, x2: 45, y2: 26}, 560},
	{Rectangle{x1: -9, y1: -11, x2: 27, y2: 21}, 1152},
	{Rectangle{x1: -10, y1: 32, x2: 27, y2: -12}, 1628},
}

var testCircleAreaCases = []testCirclePair{
	{Circle{x: 0, y: 0, r: 5}, 78.5},
	{Circle{x: 0, y: 0, r: 7.8}, 191.1},
	{Circle{x: 0, y: 0, r: 10.3}, 333.2},
	{Circle{x: 0, y: 0, r: 0.7}, 1.5},
}

var testCirclePerimeterCases = []testCirclePair{
	{Circle{x: 0, y: 0, r: 5}, 31.4},
	{Circle{x: 0, y: 0, r: 7.8}, 49},
	{Circle{x: 0, y: 0, r: 10.3}, 64.7},
	{Circle{x: 0, y: 0, r: 0.7}, 4.3},
}

func TestRectPerimeter(t *testing.T) {
	for _, test := range testRectPerimeterCases {
		v := (test.inputVal.perimeter())
		if test.outputVal != v {
			t.Error("Expected value for input ", test.inputVal, " was ", test.outputVal, "but received ", v)
		}
	}
}

func TestRectArea(t *testing.T) {

	for _, test := range testRectAreaCases {
		v := (test.inputVal.area())
		if test.outputVal != v {
			t.Error("Expected value for input ", test.inputVal, " was ", test.outputVal, "but received ", v)
		}
	}
}

func TestCirclePerimeter(t *testing.T) {
	for _, test := range testCirclePerimeterCases {
		v := (test.inputVal.perimeter())
		if test.outputVal != float64(int(v*10))/10 {
			t.Error("Expected value for input ", test.inputVal, " was ", test.outputVal, "but received ", float64(int(v*10))/10)
		}
	}
}

func TestCircleArea(t *testing.T) {
	for _, test := range testCircleAreaCases {
		v := (test.inputVal.area())
		if test.outputVal != float64(int(v*10))/10 {
			t.Error("Expected value for input ", test.inputVal, " was ", test.outputVal, "but received ", float64(int(v*10))/10)
		}
	}
}
