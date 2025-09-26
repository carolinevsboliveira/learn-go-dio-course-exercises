package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"soma de números positivos", 2, 3, 5},
		{"soma de números negativos", -2, -3, -5},
		{"soma de positivo e negativo", 5, -3, 2},
		{"soma com zero", 0, 5, 5},
		{"soma de decimais", 1.5, 2.5, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v, esperado %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"subtração de números positivos", 5, 3, 2},
		{"subtração de números negativos", -2, -3, 1},
		{"subtração de positivo e negativo", 5, -3, 8},
		{"subtração com zero", 5, 0, 5},
		{"subtração de decimais", 4.5, 2.5, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v, esperado %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"multiplicação de números positivos", 2, 3, 6},
		{"multiplicação de números negativos", -2, -3, 6},
		{"multiplicação de positivo e negativo", 2, -3, -6},
		{"multiplicação com zero", 5, 0, 0},
		{"multiplicação de decimais", 1.5, 2.5, 3.75},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v, esperado %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
	}{
		{"divisão de números positivos", 6, 2, 3, false},
		{"divisão de números negativos", -6, -2, 3, false},
		{"divisão de positivo e negativo", 6, -2, -3, false},
		{"divisão de decimais", 7.5, 2.5, 3, false},
		{"divisão por zero", 5, 0, 0, true},
		{"divisão resultando em decimal", 5, 2, 2.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%v, %v) deveria retornar erro, mas não retornou", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%v, %v) retornou erro inesperado: %v", tt.a, tt.b, err)
				}
				if !math.IsNaN(result) && !math.IsNaN(tt.expected) {
					if math.Abs(result-tt.expected) > 1e-9 {
						t.Errorf("Divide(%v, %v) = %v, esperado %v", tt.a, tt.b, result, tt.expected)
					}
				}
			}
		})
	}
}
