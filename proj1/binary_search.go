package proj1

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
