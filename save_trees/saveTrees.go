package save_trees

import "fmt"

func decompose(n int64) (out []int64) {
	left := n*n
	for i := n - 1; i > 0; i -- {
		if i*i <= left {
			fmt.Printf(
				"%v - %v = %v -> i=%v \n", left, i*i, left - i*i, i,
			)
			left -= i*i
			out = append(out, i)
		}
	}
	return
}

func main() {
	fmt.Println(decompose(5))
	fmt.Println(decompose(7))
}