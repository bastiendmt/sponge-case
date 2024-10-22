package main

import (
	"testing"
)

func TestAlternateCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "hElLo"},
		{"world", "wOrLd"},
		{"HELLO", "hElLo"},
		{"WORLD", "wOrLd"},
		{"Hello World", "hElLo wOrLd"},
		{"12345", "12345"},
		{"!@#$%^", "!@#$%^"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := alternateCase(tc.input)
			if actual != tc.expected {
				t.Errorf("alternateCase(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}
