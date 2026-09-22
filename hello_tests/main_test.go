package main

import (
	"testing"
)

func TestAddNumbers(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"1+1", 1, 1, 2},
		{"0+0", 0, 0, 0},
		{"-1+1", -1, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddNumbers(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("AddNumbers(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
