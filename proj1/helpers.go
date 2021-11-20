package proj1

import (
	"fmt"
	"time"
)

func Power(x int, y int) int {
	res := 1
	for i := 1; i <= y; i++ {
		res *= x
	}
	return res
}

func PopulateArray(length int) []int {
	arry := make([]int, length)
	for i := 0; i < length; i++ {
		arry[i] = i
	}
	return arry
}

func TimeDiff(start time.Time, name string) {
	// elapsed := time.Since(start).Seconds()
	elapsed := time.Since(start)
	fmt.Printf("%s wynosi %s.\n", name, elapsed)
}
