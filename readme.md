# Quad ASCII Generator (Go)

This project prints simple ASCII-style rectangles directly in the terminal.  
You choose which Quad function to use (A–E) and provide the width and height as command-line arguments.  
The program validates the input and then prints the corresponding rectangle.

It is beginner-friendly and demonstrates how to work with loops, conditions, helper functions, and command-line arguments in Go.

---

## What the Program Does

The project includes five functions:

- `QuadA`
- `QuadB`
- `QuadC`
- `QuadD`
- `QuadE`

Each one prints a rectangle using its own set of border characters.

All Quads use the same helper function:

```
printCornerOrLine(i, x, left, mid, right)
```

This helper decides whether the current position in a row should be a corner, an edge, or a middle character.

---

## How It Works

Each Quad function:

- Accepts two integers:
  - `x` → width (characters per line)
  - `y` → height (number of lines)
- Validates that both values are positive
- Uses nested loops to build the rectangle row by row
- Prints the correct character depending on the position
- Prints directly to the terminal (does not return a string)

---

## Running the Program

The program accepts user input through command-line arguments.

### Syntax

```
go run . <quad> <width> <height>
```

### Example

```
go run . A 5 3
```

Output:

```
o---o
|   |
o---o
```


### Accepted Quad Names

The program accepts multiple variations for convenience:

```
A, a, quadA, quada, QuadA, Quada
B, b, quadB, quadb, QuadB, Quadb
C, c, quadC, quadc, QuadC, Quadc
D, d, quadD, quadd, QuadD, Quadd
E, e, quadE, quade, QuadE, Quade
```

---

## Input Validation

The program checks:

- Width and height must be valid integers  
- Width and height must be positive  
- Quad name must match one of the accepted variations  

If any rule is violated, the program prints an error message and stops.

---

## Installation

If you do not have Go installed, download it from:

https://go.dev/dl

Verify installation:

```
go version
```

---

## Project Structure

```
piscine/
│
├── quadA.go
├── quadB.go
├── quadC.go
├── quadD.go
├── quadE.go
└── printCornerOrLine.go
```

Each Quad file contains one rectangle style.  
The helper file contains the shared printing logic.