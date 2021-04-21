// https://www.codewars.com/kata/54f8693ea58bce689100065f
// not solved
// todo

package egyptian_decomposition

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Solve(n, m int) (result []string) {
	if n*m == 0 {
		return []string{}
	}
	if n%m == 0 {
		return []string{strconv.Itoa(n / m)}
	}
	if n > m {
		result = append(result, strconv.Itoa(n/m))
		n = n % m
	}
	for i := 2; ; i++ {
		if i > 1<<35 {
			log.Fatal("overflow")
		}
		if i*n-m < 0 {
			continue
		}
		result = append(result, fmt.Sprintf("1/%v", i))
		if i*n-m == 0 {break}
		n, m = i*n-m, m*i
	}
	return
}

func prepareDash(s string) (aI, bI int) {
	split := strings.Split(s, "/")
	a, b := split[0], split[1]
	aI, _ = strconv.Atoi(a)
	bI, _ = strconv.Atoi(b)
	return
}

func prepareDot(s string) (aI, bI int) {
	split := strings.Split(s, ".")
	aI, _ = strconv.Atoi(split[1])
	power := len(split[0])
	bI = int(math.Pow10(power + 1))
	return
}

func prepareInt(s string) (aI, bI int) {
	aI, _ = strconv.Atoi(s)
	bI = 1
	return
}

func Decompose(s string) []string {
	var f func (string) (int, int)
	switch {
	case strings.ContainsRune(s, '/'): f = prepareDash
	case strings.ContainsRune(s, '.'): f = prepareDot
	default: f = prepareInt
	}
	return Solve(f(s))
}
