package piscine

import "fmt"

func printLBifLast(i, x int) {
    if i == x-1 {
        fmt.Println()
    }
}

func printIfMiddleLines(i, x int) {
    if i == 0 || i == x-1 {
        fmt.Printf("|")
        // check if last letter if true \n
        printLBifLast(i, x)
    } else {
        fmt.Printf(" ")
    }
}

func printIfFirstAndLastLines(j, i, y, x int) {
    if j == 0 || j == x-1 {
        if i == 0 || i == x-1 {
            fmt.Printf("o")
            // check if last letter if true \n
            printLBifLast(i, x)
        } else {
            fmt.Printf("-")
        }
    } else {
        printIfMiddleLines(i, x)
    }
}

func QuadA(x, y int) {
    // check if input is null
    if x <= 0 || y <= 0 {
        return
    }
    for j := 0; j < y; j++ {
        for i := 0; i < x; i++ {
            printIfFirstAndLastLines(j, i, y, x)
        }
    }
}
