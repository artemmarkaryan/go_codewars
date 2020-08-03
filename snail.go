package main

func Snail(field [][]int) (result []int) {
	switch len(field) {
	case 0: return []int{}
	case 1: return field[0]
	}

	x, y := 0, 0
	xMin, yMin := 0, -1
	xMax, yMax := len(field[0])-1, len(field)-1
	xRising, yRising := 0, 1  // rising shows directions, where to "move"

	for i := 0; i < len(field[0]) * len(field); i++ {
		result = append(result, field[x][y])
		switch [2]int{x,y} {
		case [2]int{xMin, yMin}: xRising=0;  yRising=1
		case [2]int{xMin, yMax}: xRising=1;  yRising=0
		case [2]int{xMax, yMin}: xRising=-1; yRising=0
		case [2]int{xMax, yMax}: xRising=0;  yRising=-1
		case [2]int{xMax-1, yMax}: xMin++
		case [2]int{xMax, yMin+1}: yMax--
		case [2]int{xMin+1, yMin}: xMax--
		case [2]int{xMin, yMax-1}: yMin++
		}

		switch xRising {case 1: x++; case -1: x--}
		switch yRising {case 1: y++; case -1: y--}
	}

	return result
}


func main() {
}
