package fix_string_case

import (
	"fmt"
	"strings"
)


func SolveFixStringCase(s string) string {
	lows, highs := 0, 0
	for _, v := range s {
		if v > 90 {lows++} else {highs ++}
	}
	if highs > lows {return strings.ToUpper(s)}
	return strings.ToLower(s)
}

func main() {
	s := "aBCD"
	fmt.Printf("%v -> %v", s, SolveFixStringCase(s))
}
