package main

import (
	"fmt"
)

func gcd(a int, b int) int {
    if b != 0 {
		return gcd(b, a%b)
	}
	return a
}


func Solution(a []int) (r int) {
	r = a[0]
	for j := 1; j < len(a); j++ {
		if r == 1 {break}
		r = gcd(r, a[j])
	}
	return r * len(a)
}

func main() {
	n := []int{6, 9, 21, 48, 120, 405}
	fmt.Printf("%v -> %v\n", n, Solution(n))
}
