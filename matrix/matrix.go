package matrix

import (
	"errors"
	"fmt"
	"github.com/jatin-malik/yart/util"
	"strings"
)

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

func NewIdentity(size int) *Matrix {
	grid := make([][]float64, size)
	for i := range grid {
		grid[i] = make([]float64, size)
	}

	for i := 0; i < size; i++ {
		grid[i][i] = 1
	}

	return &Matrix{rows: size, cols: size, grid: grid}
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

func (m *Matrix) MultiplyMatrix(m2 *Matrix) *Matrix {
	// Ensure multiplication compatibility
	if m.cols != m2.rows {
		return nil
	}

	// Init resultant matrix
	res := New(m.rows, m2.cols)

	for r := 0; r < res.rows; r++ {
		for c := 0; c < res.cols; c++ {
			s := 0.0
			for i := 0; i < m.cols; i++ {
				s += m.grid[r][i] * m2.grid[i][c]
			}
			res.grid[r][c] = s
		}
	}

	return res
}

func (m *Matrix) MultiplyTuple(t []float64) []float64 {
	m2 := New(len(t), 1)

	for r := 0; r < m2.rows; r++ {
		m2.grid[r][0] = t[r]
	}

	res := m.MultiplyMatrix(m2)

	if res == nil {
		return nil
	}

	var resTuple []float64
	for r := 0; r < res.rows; r++ {
		resTuple = append(resTuple, res.grid[r][0])
	}
	return resTuple
}

func (m *Matrix) Transpose() *Matrix {
	r := m.rows
	c := m.cols

	res := New(c, r)

	for r := 0; r < res.rows; r++ {
		for c := 0; c < res.cols; c++ {
			res.grid[r][c] = m.grid[c][r]
		}
	}

	return res
}

func (m *Matrix) Determinant() (float64, error) {

	// Check for square matrix
	if m.rows != m.cols {
		return -1.0, errors.New("matrix should be square")
	}

	if m.rows == 2 {
		return m.grid[0][0]*m.grid[1][1] - m.grid[1][0]*m.grid[0][1], nil
	}

	s := 0.0
	for i := 0; i < m.cols; i++ {
		x := m.grid[0][i]
		cof, _ := m.Cofactor(0, i)
		s += x * cof
	}

	return s, nil
}

func (m *Matrix) Submatrix(row, col int) *Matrix {
	res := New(m.rows-1, m.cols-1)

	i := 0
	for r := 0; r < m.rows; r++ {
		if r == row {
			continue
		}
		j := 0
		for c := 0; c < m.cols; c++ {
			if c == col {
				continue
			}
			res.grid[i][j] = m.grid[r][c]
			j++
		}
		i++
	}

	return res
}

func (m *Matrix) Minor(r, c int) (float64, error) {
	sm := m.Submatrix(r, c)
	return sm.Determinant()
}

func (m *Matrix) Cofactor(r, c int) (float64, error) {
	minor, err := m.Minor(r, c)
	if err != nil {
		return 0.0, err
	}
	if (r+c)%2 == 1 {
		return -minor, nil
	}
	return minor, nil
}

func (m *Matrix) IsInvertible() bool {
	det, err := m.Determinant()
	if err != nil {
		return false
	}
	return det != 0
}

// Inverse gives the inverse of the matrix.
func (m *Matrix) Inverse() (*Matrix, error) {
	if !m.IsInvertible() {
		return nil, errors.New("matrix is not invertible")
	}

	det, _ := m.Determinant()

	cofMatrix := New(m.rows, m.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			cof, _ := m.Cofactor(r, c)
			cofMatrix.grid[r][c] = cof
		}
	}

	res := cofMatrix.Transpose()
	res.DivideEach(det)
	return res, nil
}

// DivideEach divides each element of the matrix by n in-place.
func (m *Matrix) DivideEach(n float64) {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.grid[r][c] = m.grid[r][c] / n
		}
	}
}

func (m *Matrix) String() string {
	sb := strings.Builder{}
	for r := 0; r < m.rows; r++ {
		sb.WriteString(fmt.Sprintf("%v\n", m.grid[r]))
	}
	return sb.String()
}
