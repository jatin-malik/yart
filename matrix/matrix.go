package matrix

import "github.com/jatin-malik/yart/util"

type Matrix struct {
	rows int
	cols int
	grid [][]float64
}

func New(rows, cols int) *Matrix {
	grid := make([][]float64, rows)
	for i := range grid {
		grid[i] = make([]float64, cols)
	}
	return &Matrix{rows: rows, cols: cols, grid: grid}
}

func (m *Matrix) Equals(m2 *Matrix) bool {
	if m.rows != m2.rows || m.cols != m2.cols {
		return false
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if !util.EqualFloat64(m.grid[i][j], m2.grid[i][j]) {
				return false
			}
		}
	}
	return true
}

func NewFromGrid(grid [][]float64) *Matrix {
	matrix := New(len(grid), len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			matrix.grid[i][j] = grid[i][j]
		}
	}
	return matrix
}
