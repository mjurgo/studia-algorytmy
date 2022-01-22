package proj3

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func EstTimeCompl(fn func([]int), arry []int) time.Duration {
	iterations := 10

	var elapsedTime time.Duration
	var maxTime time.Duration
	minTime := 99 * time.Hour

	for i := 0; i < iterations+2; i++ {
		start := time.Now()
		fn(arry)
		iterationElapsed := time.Since(start)

		elapsedTime += iterationElapsed

		if iterationElapsed < minTime {
			minTime = iterationElapsed
		}
		if iterationElapsed > maxTime {
			maxTime = iterationElapsed
		}
	}

	elapsedTime -= (minTime + maxTime)
	elapsedTime = time.Duration((elapsedTime.Nanoseconds() / int64(iterations)))

	// fmt.Printf("Średni czas wykonania funkcji %s dla tablicy %d elementów to %s.\n",
	// 	runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), len(arry), elapsedTime)
	return elapsedTime
}

func TestInsertionSort(n int) {
	ascArry := AscendingArry(n)
	descArry := DescendingArry(n)
	randArry := RandomArry(n)
	constArry := ConstantArry(n, 10)
	vArry := VShapedArry(n)

	fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy rosnącej %d elementów to %s.\n",
		n, EstTimeCompl(InsertionSort, ascArry))

	fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy malejącej %d elementów to %s.\n",
		n, EstTimeCompl(InsertionSort, descArry))

	fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy losowej %d elementów to %s.\n",
		n, EstTimeCompl(InsertionSort, randArry))
	fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy stałej %d elementów to %s.\n",
		n, EstTimeCompl(InsertionSort, constArry))
	fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy V-kształtnej %d elementów to %s.\n",
		n, EstTimeCompl(InsertionSort, vArry))
}

func TestSort(fn func([]int), n int) {
	ascArry := AscendingArry(n)
	descArry := DescendingArry(n)
	randArry := RandomArry(n)
	constArry := ConstantArry(n, 10)
	vArry := VShapedArry(n)

	fmt.Printf("Średni czas wykonania %s dla tablicy rosnącej %d elementów to %s.\n",
		runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), n, EstTimeCompl(fn, ascArry))

	fmt.Printf("Średni czas wykonania %s dla tablicy malejącej %d elementów to %s.\n",
		runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), n, EstTimeCompl(fn, descArry))

	fmt.Printf("Średni czas wykonania %s dla tablicy losowej %d elementów to %s.\n",
		runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), n, EstTimeCompl(fn, randArry))

	fmt.Printf("Średni czas wykonania %s dla tablicy stałej %d elementów to %s.\n",
		runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), n, EstTimeCompl(fn, constArry))

	fmt.Printf("Średni czas wykonania %s dla tablicy V-kształtnej %d elementów to %s.\n",
		runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), n, EstTimeCompl(fn, vArry))

	// fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy losowej %d elementów to %s.\n",
	// 	n, EstTimeCompl(InsertionSort, randArry))
	// fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy stałej %d elementów to %s.\n",
	// 	n, EstTimeCompl(InsertionSort, constArry))
	// fmt.Printf("Średni czas wykonania Insertion Sort dla tablicy V-kształtnej %d elementów to %s.\n",
	// 	n, EstTimeCompl(InsertionSort, vArry))
}
