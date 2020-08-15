package main

import (
	"fmt"
)

func solveReverseString(s string) (ss string) {
	spaces := []int{}
	for i := 0; i < len(s); i++ {
		if s[i:i+1] != " " {
			ss = s[i:i+1] + ss
		} else {
			spaces = append(spaces, i)
		}
	}
	for i := 0; i < len(spaces); i++ {
		ss = ss[:spaces[i]] + " " + ss[spaces[i]:]
	}
	return
}

func main() {
	s := "abc abcd"
	fmt.Printf("%v -> %v", s, solveReverseString(s))
}
