package piscine

import "github.com/01-edu/z01"

func printIfMiddleLinesE(i, x int) {
	if i == 0 || i == x-1 {
		z01.PrintRune('B')
	} else {
		z01.PrintRune(' ')
	}
}

func printIfFirstAndLastLinesE(j, i, y, x int) {
	  if (j==0 && i==0)||(j==y-1 && i==x-1) {
			z01.PrintRune('A')
	} else if (j==0 && i==x-1)||(j==y-1 && i==0) {
			z01.PrintRune('C')
	} else if (j==0 && i>0 && i<x-1) || (j==y-1 && i>0 && i<x-1) {
			z01.PrintRune('B')
	} else {
		printIfMiddleLinesE(i, x)
	}
}

func QuadE(x, y int) {
	// check if input is null
	if x <= 0 || y <= 0 {
		return
}	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			printIfFirstAndLastLinesE(j, i, y, x)
		}
		z01.PrintRune('\n')
	}
}