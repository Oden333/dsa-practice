package structs

import (
	"testing"
)

func TestBuildList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"empty", []int{}, ""},
		{"single", []int{5}, "5"},
		{"multiple", []int{1, 2, 3}, "1 2 3"},
		{"with zeros", []int{0, 1, 0}, "0 1 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildList(tt.input)
			if gotStr := got.String(); gotStr != tt.expected {
				t.Errorf("BuildList(%v) = %q, want %q", tt.input, gotStr, tt.expected)
			}
		})
	}
}
