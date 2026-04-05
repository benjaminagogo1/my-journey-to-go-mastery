# CODECRAFTERS — STRING TRANSFORMER

**Operation: Gopher Protocol**

A command-line string processing tool built in Go to transform corrupted intelligence reports into clean, structured text.

---

## Overview

The String Transformer accepts a command and a text input, applies a transformation, and returns a formatted result.

It runs continuously in the terminal until the user exits.

---

## Features

Supports six core transformations:

| Command          | Description                               |
| ---------------- | ----------------------------------------- |
| `upper <text>`   | Converts all letters to uppercase         |
| `lower <text>`   | Converts all letters to lowercase         |
| `cap <text>`     | Capitalizes the first letter of each word |
| `title <text>`   | Title case with small-word rules          |
| `snake <text>`   | Converts text to snake_case               |
| `reverse <text>` | Reverses each word individually           |
| `exit`           | Terminates the program                    |

---

## Project Structure

 ### File structure
string-transformer/
│── main.go
│── README.md
```

---

## How to Run




go run main.go
```

---

## Example Usage

SENTINEL STRING TRANSFORMER — ONLINE

> upper sentinel is watching
→ SENTINEL IS WATCHING

> cap director adaeze okonkwo
→ Director Adaeze Okonkwo

> title the rise of the eastern threat
→ The Rise of the Eastern Threat

> snake Operation Gopher Protocol!
→ operation_gopher_protocol

> reverse Go saves the world
→ oG sevas eht dlrow

> exit
Shutting down String Transformer. Goodbye.
```

---

## Command Examples

### 🔹 upper

> upper sentinel is online
→ SENTINEL IS ONLINE
```

---

### 🔹 lower

> lower ALERT LEVEL FIVE DETECTED
→ alert level five detected
```

---

### 🔹 cap

> cap director adaeze okonkwo
→ Director Adaeze Okonkwo
```

---

### 🔹 title

> title the fall of the western power grid
→ The Fall of the Western Power Grid
```

---

### 🔹 snake

> snake Alert! Level 5 detected.
→ alert_level_5_detected
```

**Rule Definition:**

* Convert to lowercase
* Replace spaces with `_`
* Keep only:

  * letters (`a-z`)
  * digits (`0-9`)
  * underscore (`_`)
* Remove all other characters

---

### 🔹 reverse

> reverse Lagos Nigeria
→ sogaL airegiN
```

---

## Edge Case Handling

The program is designed to **never crash**, regardless of input.

### Unknown command

> fly sentinel is down
✗ Unknown command: "fly"
Valid commands: upper, lower, cap, title, snake, reverse, exit
```

---

### Missing text

> upper
✗ No text provided. Usage: upper <text>
```

---

### Extra spaces

>   upper   sentinel online
→ SENTINEL ONLINE
```

---

### Numbers and symbols

> upper alert! level 5 breach at 03:47
→ ALERT! LEVEL 5 BREACH AT 03:47
```


### Mixed-case commands

> UPPER sentinel is watching
→ SENTINEL IS WATCHING
```

---

### Single word

> reverse Okonkwo
→ ownokO
```

---

## Design Decisions

* Each transformation is implemented as its own function
* Input is split into:

  * command
  * remaining text (preserving spaces where needed)
* Validation occurs **before transformation**
* Case-insensitive command handling
* Clean separation of logic to avoid monolithic `main()`

---

## Core Concepts Used

* Strings manipulation (`strings` package)
* Slices and iteration
* Conditionals and validation
* CLI input handling
* Function decomposition

---

## 
* `history` command (last 5 transformations)
* `count <text>` (characters, words, spaces, letters)
* `palindrome <text>` checker


## Notes

* Built using **Go standard library only**
* No external dependencies
* Focused on correctness, clarity, and robustness

---

## 👤 Author

Otete Benjamin Agogo

---
