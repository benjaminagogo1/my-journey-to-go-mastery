package main

import (
	"fmt"
	"strings"
)

func splitFlag(flag string) (string, error) {
	if !strings.HasPrefix(flag, "--align=") {
		return "", fmt.Errorf("")
	}
	result := strings.Split(flag, "=")

	if len(result) != 2 {
		return flag, fmt.Errorf("invalid color format")
	}

	if result[1] == "" {
		return flag, fmt.Errorf("incorrect type format")
	}
	validTypevalue := result[1]

	return validTypevalue, nil
}

