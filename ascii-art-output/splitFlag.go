package main

import (
	"fmt"
	"strings"
)

func splitFlag(flag string) (string, error) {
	if !strings.HasPrefix(flag, "--output=") {
		return flag, fmt.Errorf("err")
	}

	result := strings.Split(flag, "=")

	if len(result) != 2 {
		return flag, fmt.Errorf("invalid")
	}
	if result[1] == "" {
		return flag, fmt.Errorf("incorrect file name")
	}
	if !strings.HasSuffix(result[1], ".txt") {
		return flag, fmt.Errorf("add `.txt` at the end of the file name")
	}
	validFileName := result[1]

	return validFileName, nil
}



