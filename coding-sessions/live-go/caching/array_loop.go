package caching

import (
	"fmt"
	"time"
)

func LoopByRows(matrix [][]int) {
	start := time.Now()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_ = matrix[i][j]
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken to loop by rows: %s\n", elapsed)
}

func LoopByColumns(matrix [][]int) {
	start := time.Now()
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			_ = matrix[i][j]
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken to loop by columns: %s\n", elapsed)
}

func Create_Matrix(rows int, columns int) *[][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns)
	}
	return &matrix
}
