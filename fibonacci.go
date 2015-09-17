package fibonacci

import "fmt"

func fibonacci(x int) int {
	if x < 0 {
		fmt.Println("Cannot calculate fibonacci for value less than zero")
		return -1
	}

	if x == 0 || x == 1 {
		return x
	}

	return fibonacci(x-1) + fibonacci(x-2)

}

func main() {
	var input int
	fmt.Println("Enter the no. whose Fibonacci value is to be found : ")
	fmt.Scanf("%d", &input)
	fmt.Println("The fibonacci value of this no. is ", fibonacci(input))
}
