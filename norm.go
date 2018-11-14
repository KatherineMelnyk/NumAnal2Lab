package main

import "math"

func norm(A [][]float64) float64 {
	var sum float64
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			sum += math.Pow(A[i][j], 2)
		}
	}
	return math.Sqrt(sum)
}

func vecNorm(v []float64) float64 {
	var res float64
	for i := 0; i < len(v); i++ {
		res += math.Pow(v[i], 2)
	}
	return math.Sqrt(res)
}
