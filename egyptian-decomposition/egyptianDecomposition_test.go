package egyptian_decomposition_test

import (
	main "github.com/artemmarkaryan/codewars/egyptian-decomposition"
	"log"
	"testing"
)

//func TestSolve(t *testing.T) {
//	type Case struct {
//		n, m int
//		result []string
//	}
//	testCases := []Case{
//		{1, 4, []string{"1/4"}},
//		{1, 1},
//	}
//
//	for _, v := range testCases {
//		n, m := v.a, v.b
//		fs := main.Solve(n, m)
//
//		log.Printf("%v/%v -> %v", n, m, fs)
//	}
//}


func TestDecompose(t *testing.T) {
	type Case struct {
		in 	string
		out []string
	}
	testCases := []Case{
		{"21/23", []string{"1/2", "1/3", "1/13", "1/359", "1/644046"}},
		{"12/4", []string{"3"}},
		{"0.66", []string{"1/2", "1/7", "1/59", "1/5163", "1/53307975"}},
		{"0", []string{}},
		{"1", []string{"1"}},
		{"5/4", []string{"1", "1/4"}},
	}

	for _, tCase := range testCases {
		log.Print(main.Decompose(tCase.in), tCase.out)
	}
}
