package main

import "math"

func maxElem(A [][]float64, k int) int {
	n := len(A)
	iMax := k //index of row with max pivot
	maxPivot := A[k][k]
	for i := k + 1; i < n; i++ {
		a := math.Abs(A[i][k])
		if a > maxPivot {
			maxPivot = a
			iMax = i
		}
	}
	return iMax
}

func GaussianElimination(A [][]float64, b []float64) bool {
	n := len(A)
	for k := 0; k < n; k++ {
		imax := maxElem(A, k)
		if A[imax][k] == 0 {
			return true
		}
		A[k], A[imax] = A[imax], A[k]
		b[k], b[imax] = b[imax], b[k]
		// matrix is singular
		for i := k + 1; i < n; i++ {
			e := A[i][k]
			t := A[k][k]
			c := e / t
			A[i][k] = 0
			for j := k + 1; j < n; j++ {
				A[i][j] = A[i][j] - A[k][j]*c
				b[i] = b[i] - b[k]*c
			}
		}
	}
	return false
}

func main() {
	arr := [][]float64{{2, -4, -1}, {-2, 3, -2}, {4, -11, -13}}
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			println(int(arr[i][j]))
		}
	}
	b := []float64{1, 1, 10}
	GaussianElimination(arr, b)
}
