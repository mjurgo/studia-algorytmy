package main

import (
	"algorytmy/proj1"
	"fmt"
	"time"
)

func main() {
	fmt.Println("test")
	arryLen := proj1.Power(2, 28)
	start := time.Now()
	arry := proj1.PopulateArray(arryLen)
	proj1.TimeDiff(start, "Populating array")
	start = time.Now()
	fmt.Println(proj1.BinSearch(arry, 214542))
	proj1.TimeDiff(start, "Binary search")
}
