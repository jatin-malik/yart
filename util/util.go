package util

import "math"

const EPSILON = 0.00001

// TODO: Do we really need to have this utility? Does Golang does this reliably for us?

func EqualFloat64(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}
