// https://www.codewars.com/kata/54f8693ea58bce689100065f
// solved

package egyptian_decomposition

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Decompose(s string) []string {
	result := []string{}
	var n, m int
	if strings.ContainsRune(s, '/') {
		split := strings.Split(s, "/")
		a, b := split[0], split[1]
		n, _ = strconv.Atoi(a)
		m, _ = strconv.Atoi(b)
	} else if strings.ContainsRune(s, '.') {
		split := strings.Split(s, ".")
		if split[0] != "0" {
			result = append(result, split[0])
		}
		n, _ = strconv.Atoi(split[1])
		m = int(math.Pow10(len(split[1])))
	} else {
		n, _ = strconv.Atoi(s)
		m = 1
	}

	if n*m == 0 {
		return result
	} else if n%m == 0 {
		return []string{strconv.Itoa(n / m)}
	} else if n > m {
		result = append(result, strconv.Itoa(n/m))
		n = n % m
	}

	var val int
	for n > 0 && m > 0 {
		val = m/n + 1
		if (val - 1) * n == m {
			val -= 1
		}
		result = append(result, fmt.Sprintf("1/%v", val))
		n, m = n*val-m, val*m
	}
	return result
}
