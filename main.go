package main

import (
	"algorytmy/proj1"
	"time"
	// "fmt"
)

func main() {
	arryLen := proj1.Power(2, 25)
	arry := proj1.PopulateArray(arryLen)

	proj1.BinSearchMinTime(arry)
	proj1.BinSearchMaxTime(arry)
	start := time.Now()
	proj1.BinSearchAvgTime(arry)
	proj1.TimeDiff(start, "Całość czasu wykonania funkcji dla średniego przypadku")
}
