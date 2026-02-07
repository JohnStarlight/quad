package piscine

/*
import "github.com/01-edu/z01"

func QuadD(x, y int) {
    z01.PrintRune('a')
}
*/
// /*
import "fmt"

func QuadB(x, y int) {
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
            printCornerOrLine(i, x, "/", "*", "\\")
            } else if j == y-1 {
            // print the corners and the line for the last line
            printCornerOrLine(i, x, "\\", "*", "/")
            } else {
            // print the middle lines
            printCornerOrLine(i, x, "*", " ", "*")
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