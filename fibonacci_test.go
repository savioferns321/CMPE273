package testCases

import (
	"fmt"
	"testing"
)

type testpair struct {
	inputVal  int
	outputVal int
}

var tests = []testpair{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
}

func fibonacciTest(x int) int {
	if x < 0 {
		fmt.Println("Cannot calculate fibonacci for value less than zero")
		return -1
	}

	if x == 0 || x == 1 {
		return x
	}

	return fibonacciTest(x-1) + fibonacciTest(x-2)

}

func TestFibonacci(t *testing.T) {
	for _, test := range tests {
		v := fibonacciTest(test.inputVal)
		if test.outputVal != v {
			t.Error("Expected value for input ", test.inputVal, " was ", test.outputVal, "but received ", v)
		}
	}
}
