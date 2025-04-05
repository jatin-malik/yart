package geometry

type Point struct {
	tuple Tuple
}

func NewPoint(x, y, z float64) Point {
	return Point{Tuple{x, y, z}}
}

func (p Point) GetX() float64 {
	return p.tuple.X
}

func (p Point) GetY() float64 {
	return p.tuple.Y
}

func (p Point) GetZ() float64 {
	return p.tuple.Z
}

func (p1 Point) Equals(p2 Point) bool {
	return p1.tuple.Equals(p2.tuple)
}

func (p Point) Add(v Vector) Point {
	t := p.tuple.Add(v.tuple)
	return Point{t}
}

func (p Point) Sub(v Vector) Point {
	t := p.tuple.Sub(v.tuple)
	return Point{t}
}

func (p Point) Negate() Point {
	return Point{p.tuple.Negate()}
}

func (p Point) String() string {
	return p.tuple.String()
}
