package ulam_sequence_test

import (
	main "github.com/artemmarkaryan/codewars/ulam_sequence"
	"log"
	"reflect"
	"testing"
)


func TestDequeRemoveDuplicates(t *testing.T) {
	type TCase struct {
		in  []int
		out []int
	}
	cases := []TCase{
		{[]int{1, 1}, []int{}},
		{[]int{1, 2}, []int{1, 2}},
	}
	for _, c := range cases {
		d := main.NewDeque(c.in)
		d.RemoveDuplicates()
		if !reflect.DeepEqual(d.Get(), c.out) {
			log.Fatalf("%#v, %#v", d.Get(), c.out)
		}
	}
}

func TestUlamSequence(t *testing.T) {
	type TCase struct {
		u0, u1, n int
		ans       []int
	}

	cases := []TCase{
		{1, 2, 3, []int{1, 2, 3}},
		{1, 2, 4, []int{1, 2, 3, 4}},
		{1, 2, 6, []int{1, 2, 3, 4, 6, 8}},
	}
	for _, c := range cases {
		res := main.UlamSequence(c.u0, c.u1, c.n)
		if !reflect.DeepEqual(res, c.ans) {
			log.Fatal(res, " != ", c.ans)
		} else {
			log.Print(res, " == ", c.ans)
		}
	}
}
