package piscine

import "fmt"

// asdasd
func NegativeInputCheck(x, y int) bool {
	if x <= 0 || y <= 0 {
		return true
	} else {
		return false
	}
}

func IterateQuad(y, x int, TL, TR, BL, BR, lineFill, centerFill string) {
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			HandleCorners(y, x, j, i, TL, TR, BL, BR, lineFill, centerFill)
			HandleCenter(y, x, j, i, centerFill)
			HandleRestOfLines()
		}
	}
}

func HandleCorners(y, x, j, i int, TL, TR, BL, BR, lineFill, centerFill string) {
	if j == 0 && i == 0 {
		fmt.Printf(TL)
	} else if j == 0 && i == x-1 {
		fmt.Println(TR)
	} else if j == y-1 && i == 0 {
		fmt.Printf(BL)
	} else if j == y-1 && i == x-1 {
		fmt.Println(BR)
	}
}

func HandleCenter(y, x, j, i int, fill string) {
	if j > y && i > x && j < y && i < x {
		fmt.Printf(fill)
	}
}

func HandleRestOfLines(y, x, j, i int, fill string) {
	if x != 0 && (x != i-1) {
		fmt.Printf(fill)
	}
}
