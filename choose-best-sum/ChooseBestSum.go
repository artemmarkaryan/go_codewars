package choose_best_sum

import (
	"sort"
)

func ChooseBestSum(maxDistance, toVisitAmount int, cities []int) int {
	if toVisitAmount > len(cities) {
		return -1
	}

	sort.Ints(cities)
	flag := true
	sum := 0
	for i := 0; i < toVisitAmount; i++ {
		sum += cities[i]
		if sum > maxDistance {
			flag = false
			break
		}
	}
	if !flag {return -1}

	highBoundary := 0
	for i := 0; i < len(cities) - toVisitAmount - 1; i ++ {
		sum = sum - cities[i] + cities[i + toVisitAmount]
		if sum > maxDistance {
			highBoundary = i + toVisitAmount
			break
		}
	}

	if highBoundary == 0 {highBoundary = len(cities)-1}


	return 0
}

func main() {
	ChooseBestSum(100, 2, []int{1, 2, 3, 50, 94, })
}