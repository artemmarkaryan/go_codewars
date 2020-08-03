package main

import "fmt"

func DirReduc(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	flag := false
	for i := 1; i < len(arr); i++ {
		switch [2]string{arr[i-1], arr[i]} {
		case [2]string{"SOUTH", "NORTH"}:
			flag = true
		case [2]string{"NORTH", "SOUTH"}:
			flag = true
		case [2]string{"EAST", "WEST"}:
			flag = true
		case [2]string{"WEST", "EAST"}:
			flag = true
		}

		if flag {
			switch i {
			case 1 :return DirReduc(arr[2:])
			case len(arr): return DirReduc(arr[:len(arr)-1])
			default: return DirReduc(append(arr[:i-1], arr[i+1:]...))
			}
		}
	}
	return arr
}

func main() {
	fmt.Print(DirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}))
}
