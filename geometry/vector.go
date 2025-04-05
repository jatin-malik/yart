package geometry

import "math"

type Vector struct {
	tuple Tuple
}

func NewVector(x, y, z float64) Vector {
	return Vector{Tuple{x, y, z}}
}

func (v Vector) GetX() float64 {
	return v.tuple.X
}

func (v Vector) GetY() float64 {
	return v.tuple.Y
}

func (v Vector) GetZ() float64 {
	return v.tuple.Z
}

func (v1 Vector) Equals(v2 Vector) bool {
	return v1.tuple.Equals(v2.tuple)
}

func (v1 Vector) Add(v2 Vector) Vector {
	t := v1.tuple.Add(v2.tuple)
	return Vector{t}
}

func (v1 Vector) Sub(v2 Vector) Vector {
	t := v1.tuple.Sub(v2.tuple)
	return Vector{t}
}

func (v Vector) Negate() Vector {
	return Vector{v.tuple.Negate()}
}

func (v Vector) Multiply(n float64) Vector {
	t := v.tuple.Multiply(n)
	return Vector{t}
}

func (v Vector) Divide(n float64) (Vector, error) {
	t, err := v.tuple.Divide(n)
	return Vector{t}, err
}

func (v Vector) Magnitude() float64 {
	x, y, z := v.GetX(), v.GetY(), v.GetZ()
	return math.Sqrt(
		math.Pow(x, 2.0) + math.Pow(y, 2.0) + math.Pow(z, 2.0))
}

func (v Vector) String() string {
	return v.tuple.String()
}

func (v Vector) Normalize() (Vector, error) {
	mag := v.Magnitude()
	nt, err := v.Divide(mag)
	if err != nil {
		return nt, err
	}
	return nt, nil
}

func (v Vector) DotProduct(v2 Vector) float64 {
	x1, y1, z1 := v.GetX(), v.GetY(), v.GetZ()
	x2, y2, z2 := v2.GetX(), v2.GetY(), v2.GetZ()
	return x1*x2 + y1*y2 + z1*z2
}

func (v Vector) CrossProduct(v2 Vector) Vector {
	x1, y1, z1 := v.GetX(), v.GetY(), v.GetZ()
	x2, y2, z2 := v2.GetX(), v2.GetY(), v2.GetZ()
	return NewVector(y1*z2-z1*y2,
		z1*x2-x1*z2,
		x1*y2-y1*x2)
}
