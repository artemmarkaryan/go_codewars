package main

import "fmt"

func Tribonacci (signature [3]float64, n int) []float64 {
	if n < 4 {return signature[:n]}
	newArr := signature[:]
	for i := n - 3; i > 0; i-- {
		newElem := newArr[len(newArr)-1] + newArr[len(newArr)-2] + newArr[len(newArr)-3]
		newArr = append(newArr, newElem)
	}
	return newArr
}

func main() {
	//fmt.Print(Tribonacci([3]float64{1.0, 2.0, 3.0}, 4))
	fmt.Printf("%v", []float64{1, 2, 3})
}
