package yart

import "testing"
import "github.com/stretchr/testify/assert"

func TestTuples(t *testing.T) {
	a := tuple{4.3, -4.2, 3.1, 1.0}

	assert.Equal(t, a.x, 4.3)
}
