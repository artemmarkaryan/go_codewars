package main

import "strings"

func FirstNonRepeating(str string) string {
	if len(str) == 0 {
		return ""
	}

	var runes = make(map[rune][]int, 100) // "t": [1 - index, 2 - occurencies]

	for idx, r := range str {
		if _, ok := runes[rune(strings.ToLower(string(r))[0])]; ok {
			runes[rune(strings.ToLower(string(r))[0])][1] += 1
		} else if _, ok := runes[rune(strings.ToUpper(string(r))[0])]; ok {
			runes[rune(strings.ToUpper(string(r))[0])][1] += 1
		} else {
			runes[r] = []int{idx, 1}
		}
	}

	var defaultIdx = len(str) + 5
	var minIdx = defaultIdx
	for _, v := range runes {
		if v[1] == 1 && v[0] < minIdx {
			minIdx = v[0]
		}
	}

	if minIdx == defaultIdx {
		return ""
	}

	return string(str[minIdx])
}
