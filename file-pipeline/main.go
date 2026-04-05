// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Structs
// Solo: Otete Benjamin Agogo [botete])
// ───────────────────────────────────────────
// Input line types:
//   - Normal lines
//   - ALL CAPS lines
//   - TODO: lines
//   - CLASSIFIED: lines
//   - Blank/dash lines
//
// Transformation rules (in order):
//   1. Trim whitespace
//   2. Replace TODO: with ✦ ACTION:
//   3. Replace CLASSIFIED: with [REDACTED]:
//   4. Remove blank/dash lines
//   5. Add line numbering (001.)
//
// Output format:
//   Header: SENTINEL FIELD REPORT — PROCESSED
//   Line numbering: 001.
//   Summary block: No
//
// Terminal summary fields:
//   Lines read, Lines written, Lines removed, Rules applied
// ═══════════════════════════════════════════

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . <input.txt> <output.txt>")
        return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    if inputFile == outputFile {
        fmt.Println("✗ Input and output cannot be the same file.")
        return
    }


    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("✗ File not found:", inputFile)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var processedLines []string

    linesRead := 0
    linesRemoved := 0

    for scanner.Scan() {
        line := scanner.Text()
        linesRead++


        line = trim(line)
        line = replaceTODO(line)
        line = replaceClassified(line)

        if isRemovable(line) {
            linesRemoved++
            continue
        }

        processedLines = append(processedLines, line)

    }

    for i := range processedLines {
        processedLines[i] = addLineNumber(processedLines[i], i+1)
    }

    out, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("✗ Cannot write to output file")
        return
    }
    defer out.Close()

    writer := bufio.NewWriter(out)

    writer.WriteString("SENTINEL FIELD REPORT — PROCESSED\n")

    for _, line := range processedLines {
        writer.WriteString(line + "\n")
    }

    writer.Flush()

    fmt.Println("✦ Lines read    :", linesRead)
    fmt.Println("✦ Lines written :", len(processedLines))
    fmt.Println("✦ Lines removed :", linesRemoved)
    fmt.Println("✦ Rules applied : trim, TODO replace, classified replace, remove blank, numbering")
}

func trim(line string) string {
    return strings.TrimSpace(line)
}

func replaceTODO(line string) string {
    return strings.ReplaceAll(strings.ToLower(line), "todo:", "✦ action:")
}

func isRemovable(line string) bool {
    if line == "" {
        return true
    }

    for _, ch := range line {
        if ch != '-' {
            return false
        }
    }
    return true
}

func replaceClassified(line string) string {
    return strings.ReplaceAll(line, "CLASSIFIED:", "[REDACTED]:")
}

func addLineNumber(line string, index int) string {
    return fmt.Sprintf("%03d. %s", index, line)
}