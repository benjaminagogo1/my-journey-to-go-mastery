package main

import "fmt"

func validateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}

	if len(banner) != 95 {
		return fmt.Errorf("banner has %d, expected 95", len(banner))
	}

	for i, v := range banner {
		if i < 32 || i > 126 {
			return fmt.Errorf("missing character")
		}
		if len(v) != 8 {
			return fmt.Errorf("expected 8 lines, got %d", len(v))
		}
	}
	return nil
}
