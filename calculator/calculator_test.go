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
