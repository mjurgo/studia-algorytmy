package proj2

import (
	// "fmt"
	"math"
)

func Eratosthenes(n int) (bool, int) {
	sqrtN := int(math.Sqrt(float64(n)))
	s := make([]bool, sqrtN+1)

	opCounter := 0

	for i := 2; i < len(s); i++ {
		s[i] = true
	}

	for j := 2; j < len(s); j++ {
		if s[j] == true {
			for k := j * 2; k <= sqrtN; k += j {
				s[k] = false
				opCounter += 1
			}
		}
	}

	for i := 2; i < len(s); i++ {
		if s[i] == true && n%i == 0 {
			return false, opCounter
		}
	}

	return true, opCounter
}
