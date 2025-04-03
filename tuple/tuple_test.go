package tuple

import (
	"github.com/jatin-malik/yart/util"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTuples(t *testing.T) {
	// 1. A tuple with w=1.0 is a point
	a := tuple{4.3, -4.2, 3.1, 1.0}

	assert.Equal(t, a.x, 4.3)
	assert.Equal(t, a.y, -4.2)
	assert.Equal(t, a.z, 3.1)
	assert.Equal(t, a.w, 1.0)

	assert.True(t, a.IsPoint())
	assert.False(t, a.IsVector())

	// 2. A tuple with w=0 is a vector
	b := tuple{4.3, -4.2, 3.1, 0.0}

	assert.Equal(t, b.x, 4.3)
	assert.Equal(t, b.y, -4.2)
	assert.Equal(t, b.z, 3.1)
	assert.Equal(t, b.w, 0.0)

	assert.False(t, b.IsPoint())
	assert.True(t, b.IsVector())
}

func TestTupleCreation(t *testing.T) {
	// 1. Point() creates tuples with w=1
	t1 := Point(4, -4, 3)

	assert.True(t, t1.Equal(tuple{4, -4, 3, 1}))

	// 2. Vector() creates tuples with w=0
	t2 := Vector(4, -4, 3)

	assert.True(t, t2.Equal(tuple{4, -4, 3, 0}))
}

func TestTupleAddition(t *testing.T) {
	a1 := tuple{3, -2, 5, 1}
	a2 := tuple{-2, 3, 1, 0}

	assert.True(t, a1.Add(a2).Equal(tuple{1, 1, 6, 1}))
}

func TestTupleSubtraction(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	assert.True(t, p1.Sub(p2).Equal(Vector(-2, -4, -6)))

	v1 := Vector(5, 6, 7)

	assert.True(t, p1.Sub(v1).Equal(Point(-2, -4, -6)))

	v2 := Vector(3, 2, 1)
	assert.True(t, v2.Sub(v1).Equal(Vector(-2, -4, -6)))
}

func TestTupleNegation(t *testing.T) {
	a := tuple{1, -2, 3, -4}
	a.Negate()
	assert.True(t, a.Negate().Equal(tuple{-1, 2, -3, 4}))
}

func TestTupleMultiplication(t *testing.T) {
	a1 := tuple{1, -2, 3, -4}

	assert.True(t, a1.Multiply(3.5).Equal(tuple{3.5, -7, 10.5, -14}))
}

func TestTupleDivision(t *testing.T) {
	a1 := tuple{1, -2, 3, -4}
	dt, err := a1.Divide(2.0)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, dt.Equal(tuple{0.5, -1, 1.5, -2}))
}

func TestTupleMagnitude(t *testing.T) {
	tests := []struct {
		v        tuple
		expected float64
	}{
		{Vector(1, 0, 0), 1.0},
		{Vector(0, 1, 0), 1.0},
		{Vector(0, 0, 1), 1.0},
		{Vector(1, 2, 3), math.Sqrt(14)},
		{Vector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, tt := range tests {
		t.Run(tt.v.String(), func(t *testing.T) {
			assert.True(t, util.EqualFloat64(tt.v.Magnitude(), tt.expected))
		})
	}

}

func TestTupleNormalization(t *testing.T) {
	tests := []struct {
		v  tuple
		nv tuple
	}{
		{Vector(4, 0, 0), Vector(1, 0, 0)},
		{Vector(1, 2, 3), Vector(0.26726, 0.53452, 0.80178)},
	}

	for _, tt := range tests {
		t.Run(tt.v.String(), func(t *testing.T) {
			nv, err := tt.v.Normalize()
			if err != nil {
				t.Error(err)
			}
			assert.True(t, nv.Equal(tt.nv))
			assert.True(t, util.EqualFloat64(tt.nv.Magnitude(), 1.0))
		})
	}
}

func TestDotProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	assert.True(t, util.EqualFloat64(v1.DotProduct(v2), 20))
}

func TestCrossProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	assert.True(t, v1.CrossProduct(v2).Equal(Vector(-1, 2, -1)))
	assert.True(t, v2.CrossProduct(v1).Equal(Vector(1, -2, 1)))
}
