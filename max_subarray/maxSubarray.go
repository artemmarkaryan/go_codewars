package max_subarray

import "fmt"

func MaximumSubarraySum(numbers []int) int {
	s := 0
	for i := 0; i < len(numbers); i++ {
		s1 := 0
		for j := i; j < len(numbers); j++ {
			//fmt.Printf("i = %v, j = %v\n", i, j)
			s1 += numbers[j]
			if s1 > s {
				s = s1;
				//fmt.Printf("sum = %v\n", s)
			}
		}

	}
	return s
}

func main() {
	fmt.Print(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
