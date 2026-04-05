# CODECRAFTERS — COMMAND & CONTROL CONSOLE

**Operation: Command & Control**

A unified Go-based CLI system integrating three independent tools — calculator, base converter, and string transformer — into a single, continuous console.

---

## Overview

The **SENTINEL Command & Control Console** eliminates the need to switch between tools by providing one prompt that controls all modules.

It supports:

* Modular commands (`calc`, `base`, `str`)
* Command piping (`|`)
* Shared history and state (`last`)
* Robust input validation

---

## Features

### 🔹 Core Modules

| Module             | Prefix | Function                             |
| ------------------ | ------ | ------------------------------------ |
| Calculator         | `calc` | Arithmetic operations                |
| Base Converter     | `base` | Convert between decimal, binary, hex |
| String Transformer | `str`  | Text transformations                 |

---

### 🔹 Global Commands

| Command   | Description                       |
| --------- | --------------------------------- |
| `help`    | Show all commands                 |
| `history` | Show last 10 inputs               |
| `clear`   | Reset session (with confirmation) |
| `exit`    | Quit program                      |

---

## How to Run

go run main.go
```

---

## Console Interface

```
════════════════════════════════════════════════
  SENTINEL — COMMAND & CONTROL CONSOLE
     All systems nominal. Type 'help' to begin.
════════════════════════════════════════════════
C&C>
```

---

## MODULE A — CALCULATOR (`calc`)

### Commands

```bash id="calc1"
calc add <a> <b>
calc sub <a> <b>
calc mul <a> <b>
calc div <a> <b>
calc mod <a> <b>
calc pow <a> <b>
calc last
calc history
```

### Example

```bash id="calc2"
C&C> calc add 42 58
✦ Result: 100

C&C> calc div 10 0
✗ Error: cannot divide by zero.
```

---

## MODULE B — BASE CONVERTER (`base`)

### Commands

```bash id="base1"
base dec <number>
base hex <number>
base bin <number>
```

### Example

```bash id="base2"
C&C> base dec 255
✦ Binary : 11111111
✦ Hex    : FF

C&C> base hex 1F
✦ Decimal: 31
```

---

## MODULE C — STRING TRANSFORMER (`str`)

### Commands

```bash id="str1"
str upper <text>
str lower <text>
str cap <text>
str title <text>
str snake <text>
str reverse <text>
```

### Example

```bash id="str2"
C&C> str snake Operation Gopher Protocol!
✦ operation_gopher_protocol
```

---

## Integration Features

### 1. PIPE (`|`)

Pass output from one command into another.

```bash id="pipe1"
C&C> calc add 200 55 | base dec
✦ Result : 255
✦ Binary : 11111111
✦ Hex    : FF

C&C> base hex FF | calc add last 1
✦ Decimal: 255
✦ Result : 256
```

---

### 2. LAST (Shared State)

`last` always refers to the most recent numeric result.

```bash id="last1"
C&C> base hex 1F
✦ Decimal: 31

C&C> calc mul last 2
✦ Result: 62
```

---

### 3. HISTORY (Global)

Stores last 10 commands across all modules.

```bash id="hist1"
C&C> history
1. calc add 42 58 → 100
2. base dec 255 → 11111111 / FF
3. str upper sentinel → SENTINEL
```

---

### 4. CLEAR (Session Reset)

C&C> clear
Type CONFIRM to clear the session:
```

* Clears:

  * history
  * last value

---

## Edge Case Handling

### Calculator

* Division/mod by zero
* Non-numeric inputs
* Missing arguments
* Negative numbers supported
* No previous result for `last`

---

### Base Converter

* Invalid hex (`1G`)
* Invalid binary (`10201`)
* Missing input
* Uppercase hex output
* Negative number handling defined

---

### String Transformer

* No text provided
* Extra spaces trimmed
* Case-insensitive commands
* Symbols preserved appropriately

---

### Global

* Empty input → ignored
* Unknown command → clear error
* Program never crashes

---

## Architecture

* Central command parser
* Module dispatch pattern:

  * `calcHandler`
  * `baseHandler`
  * `strHandler`
* Shared state:

  * `lastResult`
  * `history []string`
* Pipe processor:

  * Split by `|`
  * Execute sequentially
  * Pass output forward

---

## Core Concepts Used

* `bufio.Scanner` for input
* `strings` package for parsing
* `strconv` for conversions
* Slices for history tracking
* Error handling and validation

---

## Example Full Session

```bash id="full1"
C&C> calc add 10 5
✦ Result: 15

C&C> base dec last
✦ Binary : 1111
✦ Hex    : F

C&C> str upper result is last
✦ RESULT IS 15

C&C> history
1. calc add 10 5 → 15
2. base dec last → 1111 / F
3. str upper result is last → RESULT IS 15
```

---

## Completion Checklist

* ✔ All modules functional
* ✔ Unified console working
* ✔ Pipe system implemented
* ✔ Shared `last` working
* ✔ Global history working
* ✔ Edge cases handled
* ✔ No crashes

---

## Notes

* Built using **Go standard library only**
* Designed for **robust CLI interaction**
* Emphasis on **integration over isolation**

---

## 👤 Author

Otete Benjamin Agogo

---
