package longest_consecutive_repetition

import (
	"fmt"
)

type Result struct {r rune; rep int}

func LongestRepetition(s string) Result {
	if len(s) == 0 {return Result{}}
	var r, rMax = s[0], s[0]
	var rep, repMax = 1, 1
	for i := 1; i<len(s); i++ {
		if s[i] == r {
			rep ++
			if rep > repMax {
				repMax = rep
				rMax = r
			}
		} else {
			r = s[i]
			rep = 1
		}
	}
	return Result{rune(rMax), repMax}
}

func main() {
	fmt.Println(LongestRepetition("bbbaaabaaaa"))
}

