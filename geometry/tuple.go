package geometry

import (
	"fmt"
	"github.com/jatin-malik/yart/util"
)

type Tuple struct {
	X, Y, Z float64
}

func (t Tuple) Equals(t2 Tuple) bool {
	return util.EqualFloat64(t.X, t2.X) && util.EqualFloat64(t.Y, t2.Y) && util.EqualFloat64(t.Z, t2.Z)
}

func (t Tuple) Add(t2 Tuple) Tuple {
	nt := Tuple{}
	nt.X = t.X + t2.X
	nt.Y = t.Y + t2.Y
	nt.Z = t.Z + t2.Z
	return nt
}

func (t Tuple) Sub(t2 Tuple) Tuple {
	nt := Tuple{}
	nt.X = t.X - t2.X
	nt.Y = t.Y - t2.Y
	nt.Z = t.Z - t2.Z
	return nt
}

func (t Tuple) Negate() Tuple {
	nt := Tuple{}
	nt.X = -t.X
	nt.Y = -t.Y
	nt.Z = -t.Z
	return nt
}

func (t Tuple) Multiply(n float64) Tuple {
	nt := Tuple{}
	nt.X = t.X * n
	nt.Y = t.Y * n
	nt.Z = t.Z * n
	return nt
}

func (t Tuple) Divide(n float64) (Tuple, error) {
	if n == 0.0 {
		return Tuple{}, fmt.Errorf("cannot divide by zero")
	}
	nt := Tuple{}
	nt.X = t.X / n
	nt.Y = t.Y / n
	nt.Z = t.Z / n
	return nt, nil
}

func (t Tuple) String() string {
	return fmt.Sprintf("(%f, %f, %f)", t.X, t.Y, t.Z)
}
