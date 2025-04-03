package tuple

import (
	"fmt"
	"github.com/jatin-malik/yart/util"
	"math"
)

// TODO: Why do we need to distinguish between point and a vector?
type tuple struct {
	x, y, z, w float64
}

func (t tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t tuple) IsVector() bool {
	return t.w == 0.0
}

func Point(x, y, z float64) tuple {
	return tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) tuple {
	return tuple{x, y, z, 0.0}
}

func (t tuple) Equal(t2 tuple) bool {
	return util.EqualFloat64(t.x, t2.x) && util.EqualFloat64(t.y, t2.y) && util.EqualFloat64(t.z, t2.z) && util.EqualFloat64(t.w, t2.w)
}

func (t tuple) Add(t2 tuple) tuple {
	nt := tuple{}
	nt.x = t.x + t2.x
	nt.y = t.y + t2.y
	nt.z = t.z + t2.z
	nt.w = t.w + t2.w
	return nt
}

func (t tuple) Sub(t2 tuple) tuple {
	nt := tuple{}
	nt.x = t.x - t2.x
	nt.y = t.y - t2.y
	nt.z = t.z - t2.z
	nt.w = t.w - t2.w
	return nt
}

func (t tuple) Negate() tuple {
	nt := tuple{}
	nt.x = -t.x
	nt.y = -t.y
	nt.z = -t.z
	nt.w = -t.w
	return nt
}

func (t tuple) Multiply(n float64) tuple {
	nt := tuple{}
	nt.x = t.x * n
	nt.y = t.y * n
	nt.z = t.z * n
	nt.w = t.w * n
	return nt
}

func (t tuple) Divide(n float64) (tuple, error) {
	if n == 0.0 {
		return tuple{}, fmt.Errorf("cannot divide by zero")
	}
	nt := tuple{}
	nt.x = t.x / n
	nt.y = t.y / n
	nt.z = t.z / n
	nt.w = t.w / n
	return nt, nil
}

func (t tuple) Magnitude() float64 {
	return math.Sqrt(
		math.Pow(t.x, 2.0) + math.Pow(t.y, 2.0) + math.Pow(t.z, 2.0))
}

func (t tuple) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", t.x, t.y, t.z, t.w)
}

func (t tuple) Normalize() (tuple, error) {
	mag := t.Magnitude()
	nt, err := t.Divide(mag)
	if err != nil {
		return nt, err
	}
	return nt, nil
}

func (t tuple) DotProduct(t2 tuple) float64 {
	return t.x*t2.x + t.y*t2.y + t.z*t2.z
}

func (t tuple) CrossProduct(t2 tuple) tuple {
	return Vector(t.y*t2.z-t.z*t2.y,
		t.z*t2.x-t.x*t2.z,
		t.x*t2.y-t.y*t2.x)
}
