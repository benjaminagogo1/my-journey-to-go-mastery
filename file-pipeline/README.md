#  OPERATION: GOPHER PROTOCOL
### Sentinel String Transformer  
**A CodeCrafters Weekend Mission — 2026-03-28**

---

---

The **Sentinel String Transformer** is a command-line interface (CLI) tool written in Go that processes and transforms text input based on predefined commands.



The system takes user input, applies a transformation, and outputs clean, readable text — continuously — until termination.

---

#

- Interactive CLI (runs continuously until `exit`)
- Six core string transformation commands
- Robust input handling (whitespace, casing, symbols)
- Graceful error handling (no crashes)
- Modular function design (each transformation isolated)

---

##  HOW TO RUN


Each command follows this structure:

        <command> <text>
## upper

Convert all letters to uppercase.

Example:

> upper sentinel is online
SENTINEL IS ONLINE

## lower

Convert all letters to lowercase.

Example:

> lower ALERT LEVEL FIVE
alert level five cap

## Capitalize the first letter of each word.

Example:

> cap director adaeze okonkwo
Director Adaeze Okonkwo
## title

Title case formatting with rules for small words.

Small words:
a, an, the, and, but, or, for, nor, on, at, to, by, in, of, up, as, is, it

Example:

> title the fall of the western power grid
The Fall of the Western Power Grid
## snake

Convert text to snake_case.

Lowercase
Spaces → _
Removes invalid characters (non-alphanumeric except _)

Example:

>snake Operation Gopher Protocol!
operation_gopher_protocol
## reverse

Reverse each word individually.

Example:

> reverse Go saves the world
oG sevas eht dlrow
## PROGRAM FLOW
## SENTINEL STRING TRANSFORMER — ONLINE

> exit
Goodbye...
## EDGE CASE HANDLING

The program is designed to handle all of the following safely:

- Unknown command
> fly sentinel
✗ Unknown command: "fly"
- bMissing input
> upper
✗ No text provided. Usage: upper <text>
- Extra whitespace
>   upper   sentinel   online
SENTINEL ONLINE
- Mixed case commands
> UPPER sentinel
SENTINEL
- Symbols and numbers
> upper alert! level 5 breach
ALERT! LEVEL 5 BREACH
- Empty input

Pressing Enter does nothing (no crash, no error).

## DESIGN DECISIONS
1. Function-Based Architecture

Each transformation is implemented as its own function:

- uppercaseWord
- lowercaseWord
- capitalizeFirstLetter
- titleWord
- snakeCase
- reverseWord
2. Input Parsing
Uses string splitting to separate command and text
Handles multiple spaces and trims input safely
3. Title Case Logic
Uses a predefined set of “small words”
First word is always capitalized
4. Snake Case Rules
Converts to lowercase
Replaces spaces with _
Removes special characters
