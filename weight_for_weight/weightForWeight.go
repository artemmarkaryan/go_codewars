package weight_for_weight

import (
	"fmt"
	"sort"
	"strings"
)

func RuneToInt(r rune) int { return int(r) - 48 }

func CountWeight(s string) int {
	w := 0
	for _, v := range s {
		w += RuneToInt(v)
	}
	return w
}

func CompareAndSwap(i int, j int, n *[]string) {
	switch {
	case CountWeight((*n)[i]) > CountWeight((*n)[j]):
		{
			(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
		}
	case CountWeight((*n)[i]) == CountWeight((*n)[j]):
		{
			newN := []string{(*n)[i], (*n)[j]}
			sort.Strings(newN)
			(*n)[i], (*n)[j] = newN[0], newN[1]
		}
	}
}

func OrderWeight(s string) string {
	n := strings.Split(s, " ")
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			CompareAndSwap(i, j, &n)
		}
	}
	return strings.Join(n, " ")
}

func main() {
	fmt.Println(OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"))
}
