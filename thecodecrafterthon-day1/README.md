# CODECRAFTERS — CLI CALCULATOR (PHASE 1)

A simple command-line calculator built in Go as part of the CodeCrafters Hackathon Challenge.

This project demonstrates input parsing, validation, and continuous CLI interaction using only Go’s standard library.

---

## Features

* Supports basic arithmetic operations:

  * `add <a> <b>` → Addition
  * `sub <a> <b>` → Subtraction
  * `mul <a> <b>` → Multiplication
  * `div <a> <b>` → Division

* Runs continuously until user exits

* Built-in help command

* Robust input validation (no crashes)

* Handles negative numbers correctly

---

## roject Structure

```
thecodecrafterthon-day1/
│── main.go
│── README.md
```

---

## How to Run

Make sure you have Go installed.

```bash
go run main.go
```

---

## Usage

Type a command in the terminal:

```
- Select an operation or command

- Enter the first number
- Enter the second number
And the result print out.  
---

## Available Commands

| Command   | Description                        |
| --------- | ---------------------------------- |
| `add a b` | Adds two numbers                   |
| `sub a b` | Subtracts second number from first |
| `mul a b` | Multiplies two numbers             |
| `div a b` | Divides first number by second     |
| `help`    | Displays usage instructions        |
| `quit`    | Exits the calculator               |

---

## Input Validation

The program safely handles all invalid inputs:

### Division by zero

```
> div 10 0
✗ Error: cannot divide by zero
```

### Non-numeric input

```
> add cat dog
✗ Error: arguments must be numbers
```

### Wrong number of arguments

```
> add 5
✗ Error: expected 2 arguments
```

### Unknown command

```
> power 2 3
✗ Unknown command. Type 'help' for usage.
```

---

## Continuous Execution

The calculator keeps running after each operation:

```
> add 1 2
✦ Result: 3

> sub 5 3
✦ Result: 2
```

Until the user exits:

```
> quit
 Goodbye!
```


## Design Considerations

* Input is parsed using string splitting
* Validation is done **before execution**
* Each operation is handled in a structured way (not a long if-else chain)
* Errors are handled gracefully to prevent crashes

---
## Note

* Built using **Go standard library only**
* No external packages used
* Designed for clarity, stability, and correctness

---

##  Author

Otete Benjamin Agogo [botete]

---
