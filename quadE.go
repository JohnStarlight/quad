package piscine

import "github.com/01-edu/z01"

func printLBifLastE(i, x int) {
	if i == x-1 {
		z01.PrintRune('\n')
	}
}

func printIfMiddleLinesE(i, x int) {
	if i == 0 || i == x-1 {
		z01.PrintRune('B')
		// check if last letter if true \n
		printLBifLastE(i, x)
	} else {
		z01.PrintRune(' ')
	}
}

func printIfFirstAndLastLinesE(j, i, y, x int) {
	  if (j == 0 && i == 0)||(j == y-1 && i == x-1) {
			z01.PrintRune('A')
			// check if last letter if true \n
			printLBifLasE(i, x)
	} else if (j == 0 && i == x-1)||(j == y-1 && i == 0) {
			z01.PrintRune('C')
	} else if (j==0 && i>0 && i<x-1) || (j==y-1 && i>0 && i<x-1) {
			z01.PrintRune('B')
	} else {
		printIfMiddleLinesE(i, x)
	}
}

func QuadE(x, y int) {
	// check if input is null
	// if x <= 0 || y <= 0 {
	// 	return 0
	// }
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			printIfFirstAndLastLinesE(j, i, y, x)
		}
	}
}