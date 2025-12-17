package main

import (
	"fmt"
)

func PrintSimplexTable(table SimplexTable) {
	m := len(table.Matrix)
	n := len(table.Matrix[0])

	fmt.Printf("%-6s", "")
	for _, label := range table.ColLabels {
		fmt.Printf("%8s", label)
	}
	fmt.Println()

	for i := range m {
		fmt.Printf("%-6s", table.RowLabels[i])
		for j := range n {
			fmt.Printf("%8.2f", table.Matrix[i][j])
		}
		fmt.Println()
	}
}

func PrintSolution(table SimplexTable) {
	limCount := len(table.Matrix) - 1
	varCount := len(table.Matrix[0]) - 1
	res := make([]float64, varCount)

	for i := range len(table.Matrix) {
		label := table.RowLabels[i]
		var idx int
		_, err := fmt.Sscanf(label, "x%d", &idx)
		if err != nil {
			continue
		}

		if idx > varCount {
			idx = idx + 1 - limCount
		}

		res[idx-1] = table.Matrix[i][0]
	}

	f := table.Matrix[len(table.Matrix)-2][0]

	for i, v := range res {
		fmt.Printf("x%d = %.3f\n", i+1, v)
	}
	fmt.Printf("F = %.3f\n", f)
}

func GenerateLabels(matrix [][]float64) (rowLabels, colLabels []string) {
	m := len(matrix) - 1
	n := len(matrix[0]) - 1

	colLabels = make([]string, n+1)
	colLabels[0] = "1"
	for j := 1; j <= n; j++ {
		colLabels[j] = fmt.Sprintf("x%d", j)
	}

	rowLabels = make([]string, m+2)
	for i := range m {
		rowLabels[i] = fmt.Sprintf("x%d", n+1+i)
	}
	rowLabels[m] = "f"
	rowLabels[m+1] = "g"

	return rowLabels, colLabels
}
