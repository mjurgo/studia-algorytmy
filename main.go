package main

import (
	// "algorytmy/proj1"
	// "algorytmy/proj2"
	"algorytmy/proj3"
)

func main() {
	// PROJ3
	proj3.TestSort(proj3.InsertionSort, 50000)
	proj3.TestSort(proj3.SelectionSort, 50000)
	proj3.TestSort(proj3.HeapSort, 50000)
	proj3.TestSort(proj3.CocktailSort, 100000)

	// PROJ1
	// arryLen := proj1.Power(2, 20)
	// arry := proj1.PopulateArray(arryLen)

	// proj1.BinSearchMinTime(arry)
	// proj1.BinSearchMaxTime(arry)
	// start := time.Now()
	// proj1.BinSearchAvgTime(arry)
	// proj1.TimeDiff(start, "Całość czasu wykonania funkcji dla średniego przypadku")
}
