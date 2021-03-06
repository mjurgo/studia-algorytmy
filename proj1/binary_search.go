package proj1

import (
	"fmt"
	"time"
)

func BinSearch(arry []int, value int) (int, int) {
	low := 0
	high := len(arry) - 1
	var mid int
	var guess int
	operationCounter := 0

	for low <= high {
		mid = (low + high) / 2
		guess = arry[mid]

		operationCounter += 1

		if guess == value {
			return guess, operationCounter
		}
		if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, operationCounter
}

func BinSearchMinTime(arry []int) {
	iterations := 10
	operationsCount := 0
	var elapsed time.Duration
	var minTime, maxTime time.Duration
	minTime = 99 * time.Hour

	for i := 0; i < iterations+2; i++ {
		start := time.Now()
		_, operations := BinSearch(arry, len(arry)/2-1)

		operationsCount += operations
		iterationElapsed := time.Since(start)
		elapsed += iterationElapsed

		if iterationElapsed < minTime {
			minTime = iterationElapsed
		}
		if iterationElapsed > maxTime {
			maxTime = iterationElapsed
		}
	}

	elapsed -= (minTime + maxTime)
	elapsed = time.Duration(int(elapsed.Nanoseconds()) / iterations)

	fmt.Printf("Średni czas wykonania dla przypadku optymistycznego: %s\n", elapsed)
	fmt.Printf("Liczba operacji w przypadku optymistycznym: %d\n", operationsCount/(iterations+2))
}

func BinSearchMaxTime(arry []int) {
	iterations := 10
	operationsCount := 0
	var elapsed time.Duration
	var minTime, maxTime time.Duration
	minTime = 99 * time.Hour

	for i := 0; i < iterations+2; i++ {
		start := time.Now()
		_, operations := BinSearch(arry, len(arry)-1)

		operationsCount += operations
		iterationElapsed := time.Since(start)
		elapsed += iterationElapsed
		// fmt.Println(iterationElapsed)

		if iterationElapsed < minTime {
			minTime = iterationElapsed
		}
		if iterationElapsed > maxTime {
			maxTime = iterationElapsed
		}
	}

	elapsed -= (minTime + maxTime)
	elapsed = time.Duration(int(elapsed.Nanoseconds()) / iterations)

	fmt.Printf("Średni czas wykonania dla przypadku pesymistycznego: %s\n", elapsed)
	fmt.Printf("Liczba operacji w przypadku pesymistycznym: %d\n", operationsCount/(iterations+2))
}

func BinSearchAvgTime(arry []int) {
	iterations := 10
	operationsCount := 0
	var elapsed time.Duration
	var minTime, maxTime time.Duration
	minTime = 99 * time.Hour

	var resultsArry []int

	for i := 0; i < len(arry); i++ {
		var iterationElapsed time.Duration
		var sumForElem time.Duration
		for j := 0; j < iterations+2; j++ {
			start := time.Now()
			_, operations := BinSearch(arry, arry[i])

			operationsCount += operations
			iterationElapsed = time.Since(start)
			elapsed += iterationElapsed
			sumForElem += iterationElapsed

			if iterationElapsed < minTime {
				minTime = iterationElapsed
			}
			if iterationElapsed > maxTime {
				maxTime = iterationElapsed
			}
		}
		// fmt.Println(time.Duration(int(sumForElem.Nanoseconds()) / (iterations + 2)))
		if i%10000000 == 0 {
			resultsArry = append(resultsArry, int(sumForElem.Nanoseconds())/(iterations+2))
		}
	}
	elapsed -= (minTime + maxTime)
	elapsed = time.Duration(int(elapsed.Nanoseconds()) / (len(arry)*(iterations+2) - 2))

	fmt.Printf("Średni czas wykonania: %s\n", elapsed)
	fmt.Printf("Średnia liczba operacji: %d\n", operationsCount/len(arry)/(iterations+2))
}
