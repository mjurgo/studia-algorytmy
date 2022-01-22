package proj3

func InsertionSort(arry []int) {
	for i := 1; i < len(arry); i++ {
		j := i - 1

		for j >= 0 {
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
			}
			j--
		}
	}
}

func SelectionSort(arry []int) {
	var minIndex int

	for i := 0; i < len(arry)-1; i++ {
		minIndex = i

		for j := i + 1; j < len(arry); j++ {
			if arry[j] < arry[minIndex] {
				minIndex = j
			}
		}

		arry[minIndex], arry[i] = arry[i], arry[minIndex]
	}
}

func heapify(arry []int, n int, i int) {
	largest := i
	l := 2*i + 1 // left node
	r := 2*i + 2 // right node

	if r < n && arry[r] > arry[largest] {
		largest = r
	}
	if l < n && arry[l] > arry[largest] {
		largest = l
	}

	if largest != i {
		arry[i], arry[largest] = arry[largest], arry[i]

		heapify(arry, n, largest)
	}
}

func HeapSort(arry []int) {
	for i := len(arry)/2 - 1; i >= 0; i-- {
		heapify(arry, len(arry), i)
	}

	for i := len(arry) - 1; i > 0; i-- {
		arry[0], arry[i] = arry[i], arry[0]

		heapify(arry, i, 0)
	}
}

func CocktailSort(arry []int) {
	swapped := true
	start := 0
	end := len(arry) - 1

	for swapped {
		swapped = false

		for i := start; i < end; i++ {
			if arry[i] > arry[i+1] {
				arry[i], arry[i+1] = arry[i+1], arry[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		end--

		for i := end - 1; i >= start; i-- {
			if arry[i] > arry[i+1] {
				arry[i], arry[i+1] = arry[i+1], arry[i]
				swapped = true
			}
		}

		start++
	}
}
