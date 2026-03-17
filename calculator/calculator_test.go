package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-1, -1, 0},
	}
	for _, tt := range tests {
		got := Subtract(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{3, 4, 12},
		{0, 5, 0},
		{5, 0, 0},
		{0, 0, 0},
		{-2, 3, -6},
		{-2, -3, 6},
		{1, 7, 7},
	}
	for _, tt := range tests {
		got := Multiply(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		base, exp, want int
	}{
		// exp=0 always returns 1
		{2, 0, 1},
		{0, 0, 1},
		{-5, 0, 1},
		// negative exponent returns 0 (integer truncation)
		{2, -1, 0},
		{2, -3, 0},
		{10, -2, 0},
		// exp=1 returns base
		{5, 1, 5},
		{-3, 1, -3},
		{0, 1, 0},
		// normal positive cases
		{2, 3, 8},
		{2, 10, 1024},
		{3, 2, 9},
		{3, 4, 81},
		{10, 3, 1000},
		// base=1
		{1, 5, 1},
		{1, 100, 1},
		// base=0
		{0, 3, 0},
		{0, 5, 0},
		// negative base
		{-2, 2, 4},
		{-2, 3, -8},
		{-3, 3, -27},
	}
	for _, tt := range tests {
		got := Power(tt.base, tt.exp)
		if got != tt.want {
			t.Errorf("Power(%d, %d) = %d, want %d", tt.base, tt.exp, got, tt.want)
		}
	}
}
