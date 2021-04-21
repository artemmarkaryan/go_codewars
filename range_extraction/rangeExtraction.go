package range_extraction

import (
	"fmt"
	"strconv"
	"strings"
)

func solve(list []int) (s string) {
	var out []string
	var seq = false
	var base int

	for i := 0; i < len(list); i++ {
		fmt.Printf("v: %v, seq: %v, base: %v\n", list[i], seq, base)

		if i == len(list)-1 {
			if seq == true {
				if base + 1 == list[i] {
					out = append(out, strconv.Itoa(base))
					out = append(out, strconv.Itoa(list[i]))
				} else {
					out = append(
						out,
						fmt.Sprintf("%v-%v", base, list[i]),
					)
				}
			} else {
				out = append(out, strconv.Itoa(list[i]))
			}
			break
		}

		if seq == true {
			if list[i]+1 != list[i+1] {
				seq = false
				if base + 1 == list[i] {
					out = append(out, strconv.Itoa(base))
					out = append(out, strconv.Itoa(list[i]))
				} else {
					out = append(
						out,
						fmt.Sprintf("%v-%v", base, list[i]),
					)
				}
			}
		} else if seq == false {
			if list[i]+1 == list[i+1] {
				base = list[i]
				seq = true
			} else {
				out = append(out, strconv.Itoa(list[i]))
			}
		}
	}

	return strings.Join(out, ",")
}

func main() {
	//n := []int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20}
	n := []int{-1, 0, 1, 2, 3, 5, 6, 8}
	fmt.Println(solve(n))
}
