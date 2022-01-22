package proj3

import (
	"math"
	"math/rand"
	"time"
)

func AscendingArry(n int) []int {
	arry := make([]int, n)

	for i := 0; i < n; i++ {
		arry[i] = i
	}

	return arry
}

func DescendingArry(n int) []int {
	arry := make([]int, n)

	for i := 0; i < n; i++ {
		arry[i] = n - i
	}

	return arry
}

func RandomArry(n int) []int {
	arry := make([]int, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		arry[i] = rand.Intn(n)
	}

	return arry
}

func ConstantArry(n int, val int) []int {
	arry := make([]int, n)

	for i := 0; i < n; i++ {
		arry[i] = val
	}

	return arry
}

func VShapedArry(n int) []int {
	arry := make([]int, n+1)
	maxVal := n / 2

	for i := 0; i < n+1; i++ {
		arry[i] = int(math.Abs(float64(maxVal - i)))
	}

	return arry
}
