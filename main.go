package main

import (
    "fmt"
	//os is for Args (arguments we take from the user)
    "os"
	//strconv is for Atoi (converts ASCII to integer)
    "strconv"
    "piscine"
)

func main() {
	//check if the user has given the required arguments and not less
    if len(os.Args) < 4 {
        fmt.Println("Usage: go run . <quad> <width> <height>")
        return
    }

    quad := os.Args[1]

	//Atoi gives two values, one is the int and the other is an error if the converted input is not a valid integer
    x, errX := strconv.Atoi(os.Args[2])
    y, errY := strconv.Atoi(os.Args[3])

    //Check if numbers are valid. If they are, errX and errY are "nil" (empty)
    if errX != nil || errY != nil {
        fmt.Println("Width and height must be valid integers.")
        return
    }

    //Check if negative
    if x <= 0 || y <= 0 {
        fmt.Println("Width and height must be positive.")
        return
    }

    //Choose the function
    if quad == "A" || quad == "a" || quad == "quadA" || quad == "quada" || quad == "Quada" || quad == "QuadA" {
        piscine.QuadA(x, y)
    } else if quad == "B" || quad == "b" || quad == "quadB" || quad == "quadb" || quad == "Quadb" || quad == "QuadB" {
        piscine.QuadB(x, y)
    } else if quad == "C" || quad == "c" || quad == "quadC" || quad == "quadc" || quad == "Quadc" || quad == "QuadC" {
        piscine.QuadC(x, y)
    } else if quad == "D" || quad == "d" || quad == "quadD" || quad == "quadd" || quad == "Quadd" || quad == "QuadD" {
        piscine.QuadD(x, y)
    } else if quad == "E" || quad == "e" || quad == "quadE" || quad == "quade" || quad == "Quade" || quad == "QuadE" {
        piscine.QuadE(x, y)
    } else {
        fmt.Println("Unknown quad type. Use A, B, C, D, or E.")
    }
}

/*
import "quad/piscine"

func main() {
	piscine.QuadA(6, 6)
	piscine.QuadA(1, 0)
	piscine.QuadA(2, 3)
	piscine.QuadA(1, 1)

	piscine.QuadB(5, 5)
	piscine.QuadB(1, 2)
	piscine.QuadB(3, 9)
	piscine.QuadB(1, 1)

	piscine.QuadC(1, 3)
	piscine.QuadC(1, 1)
	piscine.QuadC(3, 3)
	piscine.QuadC(6, 3)
	piscine.QuadC(2, 3)

	piscine.QuadD(1, 1)
	piscine.QuadD(1, 3)
	piscine.QuadD(1, 0)
	piscine.QuadD(-1, 3)
	piscine.QuadD(1, 3)

	piscine.QuadE(8, 8)
	piscine.QuadE(-1, 8)
	piscine.QuadE(0, 8)
	piscine.QuadE(1, 8)
	piscine.QuadE(1, 1)
	piscine.QuadE(1, 2)
	piscine.QuadE(3, 1)
}
*/