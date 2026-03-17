package calculator

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Power returns base raised to the power of exp.
// If exp is 0, it returns 1. If exp is negative, it returns 0.
func Power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp < 0 {
		return 0
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
