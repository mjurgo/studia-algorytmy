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

func BinarySearchMax(arry []int) {
	start := time.Now()
	_, operationCount := BinSearch(arry, 0)
	TimeDiff(start, "Czas wykonania w przypadku pesymistycznym")
	fmt.Printf("Liczba operacji w przypadku pesymistycznym: %d\n", operationCount)
}

func BinarySearchMin(arry []int) {
	start := time.Now()
	_, operationCount := BinSearch(arry, len(arry)/2-1)
	TimeDiff(start, "Czas wykonania w przypadku optymistycznym")
	fmt.Printf("Liczba operacji w przypadku optymistycznym: %d\n", operationCount)
}

func BinarySearchAvg(arry []int) {
	totalOperationCount := 0
	start := time.Now()
	for i := 0; i < len(arry); i++ {
		_, operationCount := BinSearch(arry, i)
		totalOperationCount += operationCount
	}
	TimeDiff(start, "Całkowity czas wykonania w przypadku średnim")
	fmt.Printf("Średnia liczba operacji: %d\n", totalOperationCount)
}

func BinarySearchAvgTime(arry []int) {
	iterations := 10
	var elapsedTime, minTime, maxTime time.Duration

	for i := 0; i < len(arry); i++ {
		for j := 0; j < iterations+2; j++ {
			start := time.Now()
			_, _ = BinSearch(arry, arry[0])
			iterationElapsed := time.Since(start)
			elapsedTime += iterationElapsed

			if iterationElapsed < minTime {
				minTime = iterationElapsed
			}
			if iterationElapsed > maxTime {
				maxTime = iterationElapsed
			}
		}
	}

	elapsedTime -= (minTime + maxTime)
	fmt.Printf("Średni czas wynosi %f.\n", float64(elapsedTime.Microseconds())/float64(len(arry)))
}
