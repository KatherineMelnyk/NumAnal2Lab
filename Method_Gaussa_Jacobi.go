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

func mul(A [][]float64, B [][]float64) [][]float64 {
	var tmp [][]float64
	n := len(A)
	for i := 0; i < n; i++ {
		tmp = append(tmp, []float64{})
		for j := 0; j < n; j++ {
			tmp[i] = append(tmp[i], 0)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				//noinspection ALL
				tmp[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	k := [][]float64{{}, {}, {}}
	printMatrix(tmp, k)
	return tmp
}

func norm(A [][]float64) float64 {
	var sum float64
	norm := 0.
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			sum += math.Abs(A[i][j])
		}
		if sum > norm {
			norm = sum
		}
		sum = 0
	}
	return norm
}

func vecNorm(v []float64) float64 {
	var res float64
	for i := 0; i < len(v); i++ {
		res += math.Pow(v[i], 2)
	}
	return math.Sqrt(res)
}

func vecMinus(v1 []float64, v2 []float64) []float64 {
	var res []float64
	for j := 0; j < len(v1); j++ {
		res = append(res, 0)
	}
	for i := 0; i < len(v1); i++ {
		res[i] = v1[i] - v2[i]
	}
	return res
}

func inverMat(A [][]float64) [][]float64 {
	var E [][]float64
	n := len(A)
	for i := 0; i < n; i++ {
		E = append(E, []float64{})
		for j := 0; j < n; j++ {
			if i == j {
				E[i] = append(E[i], 1)
			} else {
				E[i] = append(E[i], 0)
			}
		}
	}

	GaussianElimination(A, E)
	return E
}

func diagMat(A [][]float64) [][]float64 {
	var diagonal [][]float64
	n := len(A)
	for i := 0; i < n; i++ {
		diagonal = append(diagonal, []float64{})
		for j := 0; j < n; j++ {
			if i == j {
				diagonal[i] = append(diagonal[i], A[i][i])
			} else {
				diagonal[i] = append(diagonal[i], 0)
			}
		}
	}
	return diagonal
}

func Jacobi(A [][]float64, b []float64) {
	n := len(A)
	x := []float64{-0.8, 0.9, 1.8}
	//diagonal := diagMat(A)
	var tempX []float64
	for j := 0; j < n; j++ {
		tempX = append(tempX, 0)
	}

	l := (1 - norm(A)) * EPS / norm(A)

	//for i := 0; i < n; i++ {
	//	x = append(x, 1)
	//}

	for vecNorm(vecMinus(tempX, x)) > l {
		x = tempX
		for i := 0; i < n; i++ {
			tempX[i] = b[i]
			for g := 0; g < n; g++ {
				if i != g {
					tempX[i] -= A[i][g] * x[g]
				}
				tempX[i] /= A[i][i]
			}
		}
		/*norma = math.Abs(x[0] - tmpX[0])
		for h := 0; h < n; h++ {
			if (math.Abs(x[h] - tmpX[h])) > norma {
				norma = math.Abs(x[h] - tmpX[h])
				x[h] = tmpX[h]
			}
		}*/
	}
}

func main() {
	arr := [][]float64{{2, -4, -1}, {-2, 3, -2}, {4, -11, -13}}
	//c := [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	arr2 := [][]float64{{-15.25, -10.25, 2.75}, {-8.5, -5.5, 1.5}, {2.5, 1.5, -0.5}}
	//printMatrix(arr, c)
	//GaussianElimination(arr, c)
	mul(arr, arr2)
	inverMat(arr)
	//b := []float64{1, 1, 10}
	//Jacobi(arr, b)
	//fmt.Printf("%2.3f", DET)
	fmt.Printf("%2f\n", norm(arr))
	fmt.Printf("%2f\n", norm(arr2))
	fmt.Printf("cond(A) %f \n", norm(arr)*norm(arr2))
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
