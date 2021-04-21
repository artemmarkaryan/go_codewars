// https://www.codewars.com/kata/559a28007caad2ac4e000083/
// solved

package squares_in_rectangle

func FibonacciSum(n int) (sum int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		sum += x
		x, y = y, x+y
	}
	return
}

func Perimeter(n int) int {
	return 4 * FibonacciSum(n+1)
}