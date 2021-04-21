package morse

import (
	"fmt"
	"strings"
)

var dict = map[string]string{
	".-":   "A",
	"-...": "B",
	"-.-.": "C",
	"-..":  "D",
	".":    "E",
	"..-.": "F",
	"--.":  "G",
	"....": "H",
	"..":   "I",
	".---": "J",
	"-.-":  "K",
	".-..": "L",
	"--":   "M",
	"-.":   "N",
	"---":  "O",
	".--.": "P",
	"--.-": "Q",
	".-.":  "R",
	"...":  "S",
	"-":    "T",
	"..-":  "U",
	"...-": "V",
	".--":  "W",
	"-..-": "X",
	"-.--": "Y",
	"--..": "Z",
}

func DecodeMorse(code string) (decoded string) {
	words := strings.Split(code, "   ")
	var decodedWordsArr []string
	for i := 0; i < len(words); i++ {
		word := strings.TrimSpace(words[i])
		symbols := strings.Split(word, " ")
		var decodedSymbolsArr []string
		for j := 0; j < len(symbols); j++ {
			decodedSymbolsArr = append(decodedSymbolsArr, dict[symbols[j]])
		}
		decodedWordsArr = append(decodedWordsArr, strings.Join(decodedSymbolsArr, ""))
	}
	return strings.Join(decodedWordsArr, " ")
}

func main() {
	fmt.Print(DecodeMorse(".... . -.--   .--- ..- -.. ."))
}
