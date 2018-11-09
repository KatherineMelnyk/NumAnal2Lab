package main

import (
	"errors"
	"fmt"
	"math"
)

const EPS = 0.000001

var DET = 1.
var CHAN = 0.

func argMax(A [][]float64, k int) int {
	n := len(A)
	iMax := k //index of row with max pivot
	maxPivot := math.Abs(A[k][k])
	for i := k + 1; i < n; i++ {
		a := math.Abs(A[i][k])
		if a > maxPivot {
			maxPivot = a
			iMax = i
			CHAN++
		}
	}
	return iMax
}

func GaussianElimination(A [][]float64, b [][]float64) error {
	n := len(A)
	for k := 0; k < n; k++ {
		imax := argMax(A, k)
		if A[imax][k] == 0 {
			return errors.New("matrix is singular")
		}
		A[k], A[imax] = A[imax], A[k]
		b[k], b[imax] = b[imax], b[k]

		printMatrix(A, b)

		a := A[k][k]
		for i := k; i < n; i++ {
			A[k][i] /= a
		}
		for i := 0; i < n; i++ {
			b[k][i] /= a
		}
		DET *= a

		printMatrix(A, b)

		for i := k + 1; i < n; i++ {
			c := A[i][k]
			A[i][k] = 0
			//b[i][k] -= b[k][k] * c
			for j := k + 1; j < n; j++ {
				A[i][j] -= A[k][j] * c
				//b[i][j] -= b[k][j] * c
			}
			for s := 0; s < n; s++ {
				b[i][s] -= b[k][s] * c
			}
		}
		printMatrix(A, b)
	}

	for k := n - 1; k >= 0; k-- {
		for i := k - 1; i >= 0; i-- {
			c := A[i][k]
			for j := n - 1; j >= 0; j-- { //k
				A[i][j] -= A[k][j] * c
			}
			for s := n - 1; s >= 0; s-- {
				b[i][s] -= b[k][s] * c
			}
		}
		printMatrix(A, b)
	}

	DET *= math.Pow(-1, CHAN)

	return nil
}

func main() {
	arr := [][]float64{{2, -4, -1}, {-2, 3, -2}, {4, -11, -13}}
	c := [][]float64{{1, 0, 0}, {1, 0, 0}, {10, 0, 0}}
	printMatrix(arr, c)
	GaussianElimination(arr, c)
	fmt.Printf("%2.3f", DET)
}

func printMatrix(A [][]float64, b [][]float64) {
	for i := range A {
		for j := range A[i] {
			fmt.Printf("%7.3f", A[i][j])
		}
		fmt.Printf(" | ")
		for t := range b[i] {
			fmt.Printf("  %7.3f", b[i][t])
		}
		fmt.Print("\n")
	}
	fmt.Printf("\n")
}
