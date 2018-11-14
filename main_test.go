package main

import (
	"math"
	"testing"
)

func TestGaussianElimination(t *testing.T) {
	A := [][]float64{
		{1, 2, 3},
		{6, 2, 3},
		{1, 8, 9},
	}
	A2 := copyM(A)
	E := EMatrix(3)
	E2 := EMatrix(3)

	GaussianElimination(A, E)

	if !eqMatrix(mul(A2, E), E2) {
		t.Error("3x3 failed")
	}
}

func TestEMatrix(t *testing.T) {
	t.Run("2x3", func(t *testing.T) {
		E := EMatrix(2)

		if !eqMatrix(E, [][]float64{{1, 0}, {0, 1}}) {
			t.Error("2x2 failed")
		}
	})

	t.Run("3x3", func(t *testing.T) {
		E := EMatrix(3)

		if !eqMatrix(E, [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}) {
			t.Error("3x3 failed")
		}
	})
}

func eqMatrix(A, B [][]float64) bool {
	if len(A) != len(B) {
		return false
	}
	for i, r := range A {
		if len(r) != len(B[i]) {
			return false
		}
		for j, v := range r {
			if math.Abs(v-B[i][j]) > 0.00001 {
				return false
			}
		}
	}
	return true
}
