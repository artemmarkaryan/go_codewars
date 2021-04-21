package valid_parentheses

import "fmt"

func ValidParentheses(parens string) bool {
	c := 0
	for _, v := range parens {
		switch v {
		case '(': c++
		case ')': c--
		}
		if c < 0 {return false}
	}
	return c == 0
}

func main() {
	fmt.Print(ValidParentheses(")("))
}