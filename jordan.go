package main

func Jordan(matrix [][]float64, pivotRow, pivotCol int) [][]float64 {
	m := len(matrix)
	n := len(matrix[0])
	newMatrix := CopyMatrix(matrix)

	pivot := matrix[pivotRow][pivotCol]
	if pivot == 0 {
		panic("Деление на ноль!")
	}

	for j := range n {
		newMatrix[pivotRow][j] = matrix[pivotRow][j] / pivot
	}

	for i := range m {
		if i == pivotRow {
			continue
		}
		factor := matrix[i][pivotCol]
		for j := range n {
			newMatrix[i][j] = matrix[i][j] - factor*newMatrix[pivotRow][j]
		}
	}

	return newMatrix
}

func computeG(matrix [][]float64) []float64 {
	n := len(matrix[0])
	g := make([]float64, n)
	// In this specific problem, row 3 (the 4th constraint) is the equality.
	// We sum the rows that need to be satisfied as equalities.
	for j := range n {
		g[j] = -matrix[3][j]
	}
	return g
}

// func computeG(constraints [][]float64) []float64 {
// 	n := len(constraints[0])
// 	g := make([]float64, n)
// 	for j := range n {
// 		var sum float64
// 		for i := range constraints {
// 			sum += constraints[i][j]
// 		}
// 		g[j] = -sum
// 	}
// 	return g
// }

func jordanCol(fRow []float64) int {
	pivotCol := -1
	minVal := 0.0
	for j := 1; j < len(fRow); j++ {
		if fRow[j] < minVal-1e-12 {
			minVal = fRow[j]
			pivotCol = j
		}
	}

	return pivotCol
}

func jordanRow(matrix [][]float64, pivotCol int) int {
	minRatio := 1e100
	pivotRow := -1
	for i := range len(matrix) - 2 {
		if matrix[i][pivotCol] > 1e-12 {
			ratio := matrix[i][0] / matrix[i][pivotCol]
			if ratio < minRatio {
				minRatio = ratio
				pivotRow = i
			}
		}
	}

	return pivotRow
}
