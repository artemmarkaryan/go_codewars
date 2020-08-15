package main

import (
	"fmt"
	"sort"
)


func Solve(arr []int) int {
	sort.Ints(arr)
	sum := 0
	for _, v := range arr {
		if v > sum + 1 {
			return sum + 1
		} else {sum += v}
	}
	return sum + 1
}

func main() {
	n := []int{4,2,8,3,1}
	fmt.Printf("%v -> %v", n, Solve(n))
}