package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	c1 := New(0.9, 0.6, 0.75)
	c2 := New(0.7, 0.1, 0.25)

	assert.True(t, c1.Add(c2).Equals(New(1.6, 0.7, 1.0)))
}

func TestSubtraction(t *testing.T) {
	c1 := New(0.9, 0.6, 0.75)
	c2 := New(0.7, 0.1, 0.25)

	assert.True(t, c1.Sub(c2).Equals(New(0.2, 0.5, 0.5)))
}

func TestMultiplicationByScalar(t *testing.T) {
	c := New(0.2, 0.3, 0.4)
	assert.True(t, c.Multiply(2).Equals(New(0.4, 0.6, 0.8)))
}

func TestColorMultiplication(t *testing.T) {
	c1 := New(1, 0.2, 0.4)
	c2 := New(0.9, 1, 0.1)
	assert.True(t, c1.Blend(c2).Equals(New(0.9, 0.2, 0.04)))
}
