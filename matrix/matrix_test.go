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

func TestMatrix_Multiply(t *testing.T) {
	grid1 := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	matrix1 := NewFromGrid(grid1)

	grid2 := [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}

	matrix2 := NewFromGrid(grid2)

	grid3 := [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}

	matrix3 := NewFromGrid(grid3)

	assert.True(t, matrix1.MultiplyMatrix(matrix2).Equals(matrix3))
}

func TestMatrix_MultiplyTuple(t *testing.T) {
	grid := [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}

	matrix := NewFromGrid(grid)

	tuple := []float64{1, 2, 3, 1}

	expected := []float64{18, 24, 33, 1}

	assert.Equal(t, expected, matrix.MultiplyTuple(tuple))
}

func TestMatrix_MultiplyIdentity(t *testing.T) {
	grid := [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}

	matrix := NewFromGrid(grid)

	identity := NewIdentity(4)

	assert.True(t, matrix.MultiplyMatrix(identity).Equals(matrix))
}

func TestMatrix_Transpose(t *testing.T) {
	grid := [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}

	matrix := NewFromGrid(grid)

	grid2 := [][]float64{
		{1, 2, 8, 0},
		{2, 4, 6, 0},
		{3, 4, 4, 0},
		{4, 2, 1, 1},
	}

	expected := NewFromGrid(grid2)

	assert.True(t, matrix.Transpose().Equals(expected))
}

func TestMatrix_TransposeIdentity(t *testing.T) {
	identity := NewIdentity(4)
	assert.True(t, identity.Transpose().Equals(identity))
}

func TestDeterminant(t *testing.T) {
	det, _ := NewFromGrid([][]float64{
		{1, 5},
		{-3, 2},
	}).Determinant()

	assert.Equal(t, 17.0, det)

	det, _ = NewFromGrid([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}).Determinant()

	assert.Equal(t, -196.0, det)

	det, _ = NewFromGrid([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}).Determinant()

	assert.Equal(t, -4071.0, det)

}

func TestSubmatrix(t *testing.T) {
	matrix := NewFromGrid([][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})

	expected := NewFromGrid([][]float64{
		{-3, 2},
		{0, 6},
	})

	assert.True(t, matrix.Submatrix(0, 2).Equals(expected))

	matrix = NewFromGrid([][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})

	expected = NewFromGrid([][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})

	assert.True(t, matrix.Submatrix(2, 1).Equals(expected))
}

func TestMinor(t *testing.T) {
	A := NewFromGrid([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	B := A.Submatrix(1, 0)

	det, _ := B.Determinant()
	minor, _ := A.Minor(1, 0)
	assert.Equal(t, 25.0, det)
	assert.Equal(t, 25.0, minor)
}

func TestCofactors(t *testing.T) {
	A := NewFromGrid([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	minor, _ := A.Minor(0, 0)
	cof, _ := A.Cofactor(0, 0)
	assert.Equal(t, -12.0, minor)
	assert.Equal(t, -12.0, cof)
	minor, _ = A.Minor(1, 0)
	cof, _ = A.Cofactor(1, 0)
	assert.Equal(t, 25.0, minor)
	assert.Equal(t, -25.0, cof)
}

func TestInvertibility(t *testing.T) {
	matrix := NewFromGrid([][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})

	assert.True(t, matrix.IsInvertible())

	matrix = NewFromGrid([][]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})

	assert.False(t, matrix.IsInvertible())
}

func TestInversion(t *testing.T) {
	matrix := NewFromGrid([][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})

	expected := NewFromGrid([][]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	})

	inv, _ := matrix.Inverse()
	assert.True(t, inv.Equals(expected))

	matrix = NewFromGrid([][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})

	expected = NewFromGrid([][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	})

	inv, _ = matrix.Inverse()
	assert.True(t, inv.Equals(expected))

	matrix = NewFromGrid([][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	})

	expected = NewFromGrid([][]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	})

	inv, _ = matrix.Inverse()
	assert.True(t, inv.Equals(expected))

	A := NewFromGrid([][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})

	B := NewFromGrid([][]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})

	C := A.MultiplyMatrix(B)
	invB, _ := B.Inverse()
	assert.True(t, C.MultiplyMatrix(invB).Equals(A))
}
