package main

import "fmt"

type SimplexTable struct {
	Matrix    [][]float64
	RowLabels []string
	ColLabels []string
}

func CopyMatrix(mat [][]float64) [][]float64 {
	copyMat := make([][]float64, len(mat))
	for i := range mat {
		copyMat[i] = make([]float64, len(mat[i]))
		copy(copyMat[i], mat[i])
	}
	return copyMat
}

func SimplexSolve(table *SimplexTable) {
	step := 0
	l := len(table.Matrix)

	for {
		fmt.Printf("\n=== Шаг %d ===\n", step)
		PrintSimplexTable(*table)

		fRow := table.Matrix[l-2]
		pivotCol := jordanCol(fRow)
		if pivotCol == -1 {
			fmt.Println("\nРешение:")
			break
		}

		pivotRow := jordanRow(table.Matrix, pivotCol)
		if pivotRow == -1 {
			fmt.Println("\nЗадача неразрешима")
			break
		}

		fmt.Printf("Pivot: row %d (%s), col %d (%s)\n",
			pivotRow, table.RowLabels[pivotRow],
			pivotCol, table.ColLabels[pivotCol])

		table.Matrix = Jordan(table.Matrix, pivotRow, pivotCol)
		table.RowLabels[pivotRow], table.ColLabels[pivotCol] = table.ColLabels[pivotCol], table.RowLabels[pivotRow]
		table.Matrix[l-1] = computeG(table.Matrix[:l-2])

		step++
	}
}
