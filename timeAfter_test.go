package testCases

import (
	"fmt"
	"testing"
	"time"
)

type testSleepPair struct {
	inputVal  int
	outputVal int
}

var sleepTestCases = []testSleepPair{
	{2, 2},
	{4, 4},
	{5, 5},
	{0, 0},
}

func sleepTimeCalculator(n int) int {

	fmt.Println("Starting the sleep now")
	timeBefore := time.Now()
	sleep(n)
	timeDiff := time.Since(timeBefore)
	fmt.Println("Ending the sleep now")
	return int(timeDiff.Seconds())
	//}
}

func sleep(n int) {
	<-time.After(time.Second * time.Duration(n))
}

func TestSleep(t *testing.T) {
	for _, test := range sleepTestCases {
		v := sleepTimeCalculator(test.inputVal)
		if test.outputVal != v {
			t.Error("Expected value for input ", test.inputVal, " was ", test.outputVal, "but received ", v)
		}
	}
}
