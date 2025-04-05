package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix_Equals(t *testing.T) {
	grid1 := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	matrix1 := NewFromGrid(grid1)

	grid2 := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	matrix2 := NewFromGrid(grid2)

	assert.True(t, matrix1.Equals(matrix2))
}

func TestMatrix_Equals2(t *testing.T) {
	grid1 := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	matrix1 := NewFromGrid(grid1)

	grid2 := [][]float64{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	}

	matrix2 := NewFromGrid(grid2)

	assert.False(t, matrix1.Equals(matrix2))
}
