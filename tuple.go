package yart

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
	return EqualFloat64(t.x, t2.x) && EqualFloat64(t.y, t2.y) && EqualFloat64(t.z, t2.z) && EqualFloat64(t.w, t2.w)
}
