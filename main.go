package main

func main() {

	matrix := [][]float64{
		{32, 1, 7, 1, 0, 0},
		{42, 2, 5, 0, 1, 0},
		{62, 3, 4, 0, 0, 1},
		{34, 2, 1, 0, 0, 0},
		{0, -3, -8, 0, 0, 0},
	}

	// matrix := [][]float64{
	// 	{4, 2, -1, 1, 0, 0},
	// 	{2, 1, -2, 0, 1, 0},
	// 	{5, 1, 1, 0, 0, 1},
	// 	{0, 1, -1, 0, 0, 0},
	// }

	rowLabels, colLabels := GenerateLabels(matrix)

	table := SimplexTable{
		Matrix:    matrix,
		RowLabels: rowLabels,
		ColLabels: colLabels,
	}

	gRow := computeG(table.Matrix[:4])
	table.Matrix = append(table.Matrix, gRow)
	table.RowLabels = append(table.RowLabels, "g")

	SimplexSolve(&table)
	PrintSolution(table)
}
