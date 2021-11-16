package main

import (
	"algorytmy/proj1"
)

func main() {
	arryLen := proj1.Power(2, 25)
	arry := proj1.PopulateArray(arryLen)

	// proj1.BinarySearchMax(arry)
	// proj1.BinarySearchMin(arry)
	// proj1.BinarySearchAvg(arry)
	proj1.BinarySearchAvgTime(arry)
}
