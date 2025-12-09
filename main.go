package main

func main() {
	matrix := [][]float64{
		{15, 3, 1, 1, 0, 0},
		{91, 13, 7, 0, 1, 0},
		{15, 5, 3, 0, 0, -1},
		{0, -2, -3, 0, 0, 0},
	}

	rowLabels, colLabels := GenerateLabels(matrix)

	table := SimplexTable{
		Matrix:    matrix,
		RowLabels: rowLabels,
		ColLabels: colLabels,
	}

	gRow := computeG(table.Matrix[:3])
	table.Matrix = append(table.Matrix, gRow)
	table.RowLabels = append(table.RowLabels, "g")

	SimplexSolve(&table)
	PrintSolution(table)
}
