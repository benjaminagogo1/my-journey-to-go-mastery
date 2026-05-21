package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}
	if len(banner) != 95 {
		return fmt.Errorf("banner has %d entries, expected 95", len(banner))
	}
	for i, v := range banner {
		if i < 32 && i > 126 {
			return fmt.Errorf("missing character '%c' (Ascii: %c)", i, i)
		}
		if len(v) != 8 {
			return fmt.Errorf("character %c has %d lines, expected 8", i, len(v))
		}
	}
	return nil
}
