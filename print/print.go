package print

import "fmt"

// Clears the current line.
func Clear() {
	fmt.Print("\r\033[K") // carriage return + clear to end of line
}

// Add to the end of the current line.
// Does not guard against newlines and control characters in input.
func Add(args ...string) {
	for _, s := range args {
		fmt.Print(s)
	}
}

// Updates the current line by clearing it then printing.
// Does not guard against newlines and control characters in input.
func Update(args ...string) {
	Clear()
	for _, s := range args {
		fmt.Print(s)
	}
}

// Print then newline. On empty input, just goes to next line.
// Does not guard against newlines and control characters in input.
func Line(args ...string) {
	for _, s := range args {
		fmt.Print(s)
	}
	fmt.Println("")
}
