package main

import (
	"fmt"
	//os is for Args (arguments we take from the user)
	"os"
	//strconv is for Atoi (converts ASCII to integer)
	"quad/piscine"
	"strconv"
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

	//Choose the function with a switch statement based on the user's input for the quad type.
	//Switch simplifies multiple if-else statements when checking the same variable against different values.
	//When the user inputs a quad type, it will match one of the cases and execute the corresponding function from the piscine package
	//or if the input doesn't match any case, it will execute the default case and print an error message.
	switch quad {
	case "A", "a", "quadA", "quada", "Quada", "QuadA":
		piscine.QuadA(x, y)
	case "B", "b", "quadB", "quadb", "Quadb", "QuadB":
		piscine.QuadB(x, y)
	case "C", "c", "quadC", "quadc", "Quadc", "QuadC":
		piscine.QuadC(x, y)
	case "D", "d", "quadD", "quadd", "Quadd", "QuadD":
		piscine.QuadD(x, y)
	case "E", "e", "quadE", "quade", "Quade", "QuadE":
		piscine.QuadE(x, y)
	default:
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
