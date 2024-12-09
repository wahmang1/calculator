package gxcalculator

import "fmt"

func Add(a, b int) int {
	return a + b
}

// Fungsi untuk mengurangi dua angka
func Subtract(a, b int) int {
	return a - b
}

// Fungsi untuk mengalikan dua angka
func Multiply(a, b int) int {
	return a * b
}

// Fungsi untuk membagi dua angka
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
