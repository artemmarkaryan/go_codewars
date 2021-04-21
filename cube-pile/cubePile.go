package main

import (
	"fmt"
)

func FindNb(m int) int {
  vol := 0
  for i := 1; vol <= m; i++{
  	vol += i*i*i
  	if vol == m {return i}
  }
  return -1
}

func main() {
	fmt.Print(FindNb(8))
}
