package next_bigger

import (
	"fmt"
	"math"
	"sort"
)

func GetDigits(n int) (result []int) {
	for ;n > 0; n /= 10 {
		result = append(result, n%10)
	}
	return result
}

func NextBigger(n int) (result int) {
	digits := GetDigits(n)
	pivot := 0
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			pivot = i

			min := 10
			minPosition := -1
			for j := 0; j < pivot+1; j++ {
				if digits[j] > digits[pivot+1] && digits[j] < min {
					min = digits[j]
					minPosition = j
				}
			}
			digits[minPosition], digits[i+1] = digits[i+1], digits[minPosition]
			break
		}
	}
	rightReversed := digits[:pivot+1] // gotta sort
	leftReversed := digits[pivot+1:] //  stay unchanged
	sort.Ints(rightReversed)
	right := rightReversed

	var left []int
	for i := 0; i < len(leftReversed); i++ {
		left = append(left, leftReversed[len(leftReversed)-1-i])
	}
	newDigits := append(left, right...)

	for i := 0; i < len(newDigits); i++ {
		result += newDigits[i] * int(math.Pow(10, float64(len(newDigits)-i-1)))
	}
	if result == n {
		return -1
	}
	return result
}

func main() {
	n := 59884848459853

	fmt.Printf("%v -> %v", n, NextBigger(n))
}
