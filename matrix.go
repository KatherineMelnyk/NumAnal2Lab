package main

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

func EMatrix(n int) [][]float64 {
	var E [][]float64
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

	return E
}

func InvertMat(A [][]float64) [][]float64 {
	n := len(A)
	E := EMatrix(n)

	GaussianElimination(A, E)
	return E
}

func mul(A, B [][]float64) [][]float64 {
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
	//k := [][]float64{{}, {}, {}}
	//printMatrix(tmp, k)
	return tmp
}

func copyM(A [][]float64) [][]float64 {
	B := make([][]float64, len(A))
	for i, r := range A {
		B[i] = make([]float64, len(r))
		copy(B[i], r)
	}
	return B
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
