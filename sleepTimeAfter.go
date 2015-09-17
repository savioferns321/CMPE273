package main

import (
	"fmt"
	"time"
)

func sleep(n int) int {
	fmt.Println("Starting the sleep now")
	timeNow := time.Now()
	<-time.After(time.Second * time.Duration(n))
	timeLater := time.Since(timeNow)
	fmt.Println("Ending the sleep now")
	return int(timeLater.Seconds())
}

func main() {
	//	fmt.Scanln(&input)
	var input int
	fmt.Println("Enter the number of seconds that you want to sleep for")
	fmt.Scanf("%d", &input)
	fmt.Println("Sleep lasted for ", sleep(input), " seconds")

	//	var input string

}
