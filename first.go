package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	i := Point{1, 2}
	iAddress := &i
	(*iAddress).x = 2
	fmt.Print(i.x)

}