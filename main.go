package main

import (
	"algorytmy/proj1"
	// "algorytmy/proj2"
	// "fmt"
	"time"
)

func main() {
	// PROJ1
	arryLen := proj1.Power(2, 22)
	arry := proj1.PopulateArray(arryLen)

	proj1.BinSearchMinTime(arry)
	proj1.BinSearchMaxTime(arry)
	start := time.Now()
	proj1.BinSearchAvgTime(arry)
	proj1.TimeDiff(start, "Całość czasu wykonania funkcji dla średniego przypadku")

	// PROJ2
	// arry := []int{100913, 1009139, 10091401, 100914061, 1009140611, 10091406133,
	// 	100914061337, 1009140613399}

	// for i := 0; i < len(arry); i++ {
	// 	iterations := 10
	// 	var elapsed time.Duration
	// 	var minTime, maxTime time.Duration
	// 	minTime = 99 * time.Hour
	// 	var opCounter int
	// 	var prime bool

	// 	for j := 0; j < iterations+2; j++ {
	// 		start := time.Now()
	// 		prime, opCounter = proj2.Eratosthenes(arry[i])
	// 		iterationElapsed := time.Since(start)
	// 		elapsed += iterationElapsed

	// 		if iterationElapsed < minTime {
	// 			minTime = iterationElapsed
	// 		}
	// 		if iterationElapsed > maxTime {
	// 			maxTime = iterationElapsed
	// 		}
	// 	}
	// 	elapsed -= (minTime + maxTime)
	// 	elapsed = time.Duration((elapsed.Nanoseconds() / int64(iterations)))

	// 	fmt.Printf("%d jest liczbą pierwszą: %t | Średni czas wykonania wynosi: %d | Liczba operacji: %d\n", arry[i], prime, elapsed.Nanoseconds(), opCounter)
	// }
}
