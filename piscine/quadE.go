package piscine
/*
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
}    for j := 0; j < y; j++ {
        for i := 0; i < x; i++ {
            printIfFirstAndLastLinesE(j, i, y, x)
        }
        z01.PrintRune('\n')
    }
}
*/

import "fmt"

func QuadE(x, y int) {
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
            printCornerOrLine(i, x, "C", "B", "A")
            } else {
            // print the middle lines
            printCornerOrLine(i, x, "B", " ", "B")
            }
        } 
            // print a new line after each row
            fmt.Println()
    }
}

/*
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
}
*/
