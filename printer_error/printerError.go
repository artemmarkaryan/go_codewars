package printer_error

import "fmt"

func PrinterError(s string) string {
	errors := 0
	for _, v := range s {
		if v > 'm' {errors += 1}
	}
	return fmt.Sprintf("error_printer(s) => \"%v/%v\"", errors, len(s))
}

func main() {
	fmt.Print(PrinterError("aaacncbb"))
}
