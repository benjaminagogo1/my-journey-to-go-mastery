package main

import (
	"fmt"
	"strings"
)

func splitFlag(flag string) (string, error) {
	if !strings.HasPrefix(flag, "--align=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=right something standard")
		return flag, fmt.Errorf("")
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

--align
--align=right

--align=

--align=
