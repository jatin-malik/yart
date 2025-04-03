package yart

import "testing"
import "github.com/stretchr/testify/assert"

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
