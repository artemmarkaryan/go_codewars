package high_and_low

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func highAndLow(s string) string {
	high := math.MinInt32
	low := math.MaxInt32

	for _, v := range strings.Split(s, " ") {
		integer, _ :=  strconv.Atoi(v)
		if integer < low {low = integer}
		if integer > high {high = integer}
	}
	return fmt.Sprintf("%v %v", high, low)
}

func main() {
	fmt.Print(highAndLow("1 2 3 4"))
}
