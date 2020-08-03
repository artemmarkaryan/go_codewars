package main

import (
	"fmt"
	"unicode"
)

func toCamelCase(s string) string {
	var newString string
	flagNextUpper := false
	for _, v := range s {
		if unicode.IsLetter(v) {
			if flagNextUpper {
				newString += string(unicode.ToUpper(v))
				flagNextUpper = false
			} else {
				newString += string(v)
			}
		} else {
			flagNextUpper = true
		}
	}
	return newString
}

func main() {
	fmt.Print(toCamelCase("asdasd-asdasd-asd"))
}