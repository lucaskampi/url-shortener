package generator

import (
	"testing"
)

func TestGenerateShortCode(t *testing.T) {
	code := GenerateShortCode(6)
	if len(code) != 6 {
		t.Errorf("expected length 6, got %d", len(code))
	}
}

func TestGenerateShortCodeLength(t *testing.T) {
	for length := 1; length <= 10; length++ {
		code := GenerateShortCode(length)
		if len(code) != length {
			t.Errorf("expected length %d, got %d", length, len(code))
		}
	}
}

func TestGenerateShortCodeUniqueness(t *testing.T) {
	codes := make(map[string]bool)
	for i := 0; i < 100; i++ {
		code := GenerateShortCode(6)
		if codes[code] {
			t.Errorf("duplicate code generated: %s", code)
		}
		codes[code] = true
	}
}
