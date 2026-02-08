# Quad ASCII Generator (Go)

This project prints simple ASCII-style rectangles directly in the terminal. Each rectangle is built from different border characters depending on which function you call. You choose a width and a height, and the program draws the shape line by line.

The code is beginner-friendly and helps you understand how loops, conditions, and helper functions work together in Go.

---

## What the Program Does

The project includes five functions:

- `QuadA`
- `QuadB`
- `QuadC`
- `QuadD`
- `QuadE`

Each one prints a rectangle with its own character style.

All of them use the same helper function: printCornerOrLine(i, x, left, mid, right)


This helper decides whether the current position in a row should be a corner, an edge, or a middle character. This keeps the code clean and avoids repeating the same logic in every Quad function.

---

## How It Works

Each Quad function:

- Takes two integers:
  - `x` → width (characters per line)
  - `y` → height (number of lines)
- Checks that both numbers are positive
- Loops through each row and each column
- Chooses the correct character based on the position
- Prints the rectangle directly to the terminal

The functions do not return anything; they simply print the result.

---

## Example

Calling: QuadA(5, 3)


prints:
```
o---o  
|   |  
o---o
```

Calling a different Quad function produces the same shape but with different characters.

---

## How to Run the Program

If you already have Go installed, you can run the program by creating a small `main.go` file and calling any of the Quad functions.

If you do not have Go installed, you can download it from the official website:

https://go.dev/dl

The page includes installation instructions for Windows, macOS, and Linux.

After installing Go:

1. Place the project files in a folder, for example: piscine/


2. Create a `main.go` file in the same folder or one level above and inside it paste:

```
package main

import "piscine"

func main() {
piscine.QuadA(5, 3)
piscine.QuadC(7, 4)
} 
```

3. Run the program from the terminal: 
```
go run .
```

The rectangles will appear immediately in the terminal.

---

## Input Notes

- If `x <= 0` or `y <= 0`, nothing is printed.
- Very large values can produce huge amounts of output and may slow down or freeze the terminal.
- The functions print directly to standard output and do not return strings.

---

## Project Structure

piscine/  
│  
├── quadA.go  
├── quadB.go  
├── quadC.go  
├── quadD.go  
├── quadE.go  
└── printCoLi.go


Each Quad file contains one rectangle style, and the helper printCoLi file contains the shared printing logic.
