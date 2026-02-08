package piscine

import "fmt"

// function to print the correct character based on the position in the line
func printCornerOrLine(i, x int, left, mid, right string) {
	// check if it's the first or last character in the line, and print the appropriate character
	if i == 0 {
		// print the left character
		fmt.Print(left)
	} else if i == x-1 {
		// print the right character
		fmt.Print(right)
	} else {
		// print the middle character
		fmt.Print(mid)
	}
	// fmt.Println()
}
