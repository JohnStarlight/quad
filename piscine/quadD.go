package piscine

import "fmt"

func QuadD(x, y int) {
	// check if input is null
	if x <= 0 || y <= 0 {
		// if either x or y is less than or equal to 0, we don't print anything
		return
	}
	// with "j" counting the height and "i" counting the width, we go through the rectangle row by row
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			// check if it's the first or last line, and print the appropriate characters
			if j == 0 {
				// print the corners and the line for the first line
				printCornerOrLine(i, x, "A", "B", "C")
			} else if j == y-1 {
				// print the corners and the line for the last line
				printCornerOrLine(i, x, "A", "B", "C")
			} else {
				// print the middle lines
				printCornerOrLine(i, x, "B", " ", "B")
			}
		}
		// print a new line after each row
		fmt.Println()
	}
	fmt.Println()
}
