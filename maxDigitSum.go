package main

import (
	"fmt"
	"math"
	"strconv"

)

func solveMaxDigitSum(n int) int {
	if n%10 == 9 || n%10 == 8 || n < 10 {return n}
	return 2 * int(math.Pow(10, float64(len(strconv.Itoa(n)))-1)) - 1
}

func main() {
	n := 10
	fmt.Printf("%v -> %v", n, solveMaxDigitSum(n))
}
