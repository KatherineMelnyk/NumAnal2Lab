package main

import (
	"errors"
	"fmt"
	"math"
)

const EPS = 0.0000001

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
	A = copyM(A)
	n := len(A)
	for k := 0; k < n; k++ {
		imax := argMax(A, k)
		if A[imax][k] == 0 {
			return errors.New("matrix is singular")
		}
		A[k], A[imax] = A[imax], A[k]
		b[k], b[imax] = b[imax], b[k]

		//printMatrix(A, b)

		a := A[k][k]
		for i := k; i < n; i++ {
			A[k][i] /= a
		}
		for i := 0; i < n; i++ {
			b[k][i] /= a
		}
		DET *= a

		//printMatrix(A, b)

		for i := k + 1; i < n; i++ {
			c := A[i][k]
			A[i][k] = 0
			for j := k + 1; j < n; j++ {
				A[i][j] -= A[k][j] * c
			}
			for s := 0; s < n; s++ {
				b[i][s] -= b[k][s] * c
			}
		}
		//printMatrix(A, b)
	}

	for k := n - 1; k >= 0; k-- {
		for i := k - 1; i >= 0; i-- {
			c := A[i][k]
			for j := n - 1; j >= 0; j-- {
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

func Jacobi(A [][]float64, b []float64) {
	n := len(A)
	var tmpX, x []float64
	for j := 0; j < n; j++ {
		tmpX = append(tmpX, 0)
		x = append(x, b[j]/A[j][j])
	}

	norm := 0.01
	for norm > EPS {
		for i := 0; i < n; i++ {
			tmpX[i] = b[i]
			for g := 0; g < n; g++ {
				if i != g {
					tmpX[i] -= A[i][g] * x[g]
				}
			}
			tmpX[i] /= A[i][i]
		}
		res := vecMinus(tmpX, x)
		norm = vecNorm(res)
		copy(x, tmpX)
		printMatrixVector(A, x)
	}
}

func main() {
	//arr := [][]float64{{1, 2, 3}, {-1, 0, -3}, {-1, -2, 0}}
	//c := [][]float64{{3, 0, 0, 7}, {2, 0, 0, 6}, {1, 0, 0, 5}, {1, 2, 3, 4}}
	//arr2 := [][]float64{{-15.25, -10.25, 2.75}, {-8.5, -5.5, 1.5}, {2.5, 1.5, -0.5}}
	//printMatrix(MyMatrix(4), c)
	//printMatrix(MyVector(4), c)
	//GaussianElimination(MyMatrix(100), MyVector(100))
	Jacobi(MyMatrix(100), myVector(100))
	//b := []float64{1, 1, 10}
	//arr := [][]float64{{10, -1, 1}, {1, 10, -1}, {-1, 1, 10}}
	//GaussianElimination(arr, c)
	//mul(arr, arr2)
	//InvertMat(arr)
	//b := []float64{11, 10, 10}
	//arr := [][]float64{{115, -20, -75}, {15, -50, -5}, {6, 2, 20}}
	//b := []float64{3, 2, 1}
	//Jacobi(arr, b)
	//arr := [][]float64{{1, 2}, {3, 4}}
	//fmt.Print(norm(arr) * norm(InvertMat(arr)))
	//fmt.Printf("%2.3f", DET)
	//fmt.Printf("%2f\n", norm(arr))
	//fmt.Printf("%2f\n", norm(arr2))
	//fmt.Printf("cond(A) %f \n", norm(arr)*norm(arr2))
	//a := [][]float64{{2, 3}, {4, 5}}
	//b := [][]float64{{2, 3}, {4, 5}}
	//minus(a, b)
}

func MyVector(n int) [][]float64 {
	var vector [][]float64
	for i := 0; i < n; i++ {
		vector = append(vector, []float64{})
		for j := 0; j < n; j++ {
			if j == 0 {
				vector[i] = append(vector[i], float64(i+1))
			} else {
				vector[i] = append(vector[i], 0.)
			}
		}
	}
	return vector
}

func myVector(n int) []float64 {
	var vector []float64
	for i := 0; i < n; i++ {
		vector = append(vector, float64(i+1))
	}
	return vector
}

func MyMatrix(n int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < n; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < n; j++ {
			if i == j && i != 0 && j != 0 {
				matrix[i] = append(matrix[i], 0.)
			} else if i == 0 && j == 0 && i == j {
				matrix[i] = append(matrix[i], 1.)
			} else {
				matrix[i] = append(matrix[i], float64(j+1))
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				break
			} else {
				matrix[i][j] *= -1
			}
		}
	}
	return matrix
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

func printMatrixVector(A [][]float64, b []float64) {
	for i := range A {
		for j := range A[i] {
			fmt.Printf("%9.3f", A[i][j])
		}
		fmt.Printf(" | ")
		fmt.Printf("  %7.3f", b[i])
		fmt.Print("\n")
	}
	fmt.Printf("\n")
}
