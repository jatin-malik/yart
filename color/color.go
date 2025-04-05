package color

import (
	"github.com/jatin-malik/yart/geometry"
	"math"
)

type Color struct {
	tuple geometry.Tuple
}

func New(red, green, blue float64) Color {
	return Color{tuple: geometry.Tuple{X: red, Y: green, Z: blue}}
}

func (c1 Color) Equals(c2 Color) bool {
	return c1.tuple.Equals(c2.tuple)
}

func (c1 Color) Add(c2 Color) Color {
	return Color{c1.tuple.Add(c2.tuple)}
}

func (c1 Color) Sub(c2 Color) Color {
	return Color{c1.tuple.Sub(c2.tuple)}
}

func (c Color) Multiply(n float64) Color {
	return Color{c.tuple.Multiply(n)}
}

func (c1 Color) Blend(c2 Color) Color {
	return New(c1.GetRed()*c2.GetRed(), c1.GetGreen()*c2.GetGreen(), c1.GetBlue()*c2.GetBlue())
}

func (c Color) GetRed() float64 {
	return c.tuple.X
}

func (c Color) GetGreen() float64 {
	return c.tuple.Y
}

func (c Color) GetBlue() float64 {
	return c.tuple.Z
}

func (c Color) ToByte() Color {
	clamp := func(v float64) float64 {
		return math.Min(math.Max(v, 0), 1)
	}

	r := math.Ceil(clamp(c.GetRed()) * 255)
	g := math.Ceil(clamp(c.GetGreen()) * 255)
	b := math.Ceil(clamp(c.GetBlue()) * 255)

	return New(r, g, b)
}
