package geometry

import (
	"github.com/jatin-malik/yart/util"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAddition(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v1 := NewVector(5, 6, 7)

	// Point + Vector gives Point
	assert.True(t, p.Add(v1).Equals(NewPoint(8, 8, 8)))

	// Vector + Vector gives Vector
	v2 := NewVector(3, 2, 1)
	assert.True(t, v1.Add(v2).Equals(NewVector(8, 8, 8)))
}

func TestSubtraction(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v1 := NewVector(5, 6, 7)

	// Point - Vector gives Point
	assert.True(t, p.Sub(v1).Equals(NewPoint(-2, -4, -6)))

	// Vector - Vector gives Vector
	v2 := NewVector(3, 2, 1)
	assert.True(t, v2.Sub(v1).Equals(NewVector(-2, -4, -6)))
}

func TestTupleNegation(t *testing.T) {
	a := Tuple{1, -2, 3}
	a.Negate()
	assert.True(t, a.Negate().Equals(Tuple{-1, 2, -3}))
}

func TestTupleMultiplication(t *testing.T) {
	a1 := Tuple{1, -2, 3}

	assert.True(t, a1.Multiply(3.5).Equals(Tuple{3.5, -7, 10.5}))
}

func TestTupleDivision(t *testing.T) {
	a1 := Tuple{1, -2, 3}
	dt, err := a1.Divide(2.0)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, dt.Equals(Tuple{0.5, -1, 1.5}))
}

func TestVectorMagnitude(t *testing.T) {
	tests := []struct {
		v        Vector
		expected float64
	}{
		{NewVector(1, 0, 0), 1.0},
		{NewVector(0, 1, 0), 1.0},
		{NewVector(0, 0, 1), 1.0},
		{NewVector(1, 2, 3), math.Sqrt(14)},
		{NewVector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, tt := range tests {
		t.Run(tt.v.String(), func(t *testing.T) {
			assert.True(t, util.EqualFloat64(tt.v.Magnitude(), tt.expected))
		})
	}

}

func TestVectorNormalization(t *testing.T) {
	tests := []struct {
		v  Vector
		nv Vector
	}{
		{NewVector(4, 0, 0), NewVector(1, 0, 0)},
		{NewVector(1, 2, 3), NewVector(0.26726, 0.53452, 0.80178)},
	}

	for _, tt := range tests {
		t.Run(tt.v.String(), func(t *testing.T) {
			nv, err := tt.v.Normalize()
			if err != nil {
				t.Error(err)
			}
			assert.True(t, nv.Equals(tt.nv))
			assert.True(t, util.EqualFloat64(tt.nv.Magnitude(), 1.0))
		})
	}
}

func TestDotProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	assert.True(t, util.EqualFloat64(v1.DotProduct(v2), 20))
}

func TestCrossProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	assert.True(t, v1.CrossProduct(v2).Equals(NewVector(-1, 2, -1)))
	assert.True(t, v2.CrossProduct(v1).Equals(NewVector(1, -2, 1)))
}
