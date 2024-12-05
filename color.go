package mathf

import (
	"fmt"
)

type Color struct {
	R, G, B, A float64
}

func NewColor(r, g, b, a float64) Color {
	return Color{
		R: float64(r),
		G: float64(g),
		B: float64(b),
		A: float64(a),
	}
}

func (val Color) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", val.R, val.G, val.B, val.A)
}

func (c Color) Add(other Color) Color {
	return Color{
		R: c.R + other.R,
		G: c.G + other.G,
		B: c.B + other.B,
		A: c.A + other.A,
	}
}

func (c Color) Sub(other Color) Color {
	return Color{
		R: c.R - other.R,
		G: c.G - other.G,
		B: c.B - other.B,
		A: c.A - other.A,
	}
}

func (c Color) Mul(other Color) Color {
	return Color{
		R: c.R * other.R,
		G: c.G * other.G,
		B: c.B * other.B,
		A: c.A * other.A,
	}
}

func (c Color) Mulf(f float64) Color {
	return Color{
		R: c.R * f,
		G: c.G * f,
		B: c.B * f,
		A: c.A * f,
	}
}

func (c Color) Clamp() Color {
	return Color{
		R: clampf64(c.R, 0, 1),
		G: clampf64(c.G, 0, 1),
		B: clampf64(c.B, 0, 1),
		A: clampf64(c.A, 0, 1),
	}
}

func (c Color) Invert() Color {
	return Color{
		R: 1 - c.R,
		G: 1 - c.G,
		B: 1 - c.B,
		A: c.A,
	}
}

func (c Color) Lerp(to Color, t float64) Color {
	return Color{
		R: lerpf64(c.R, to.R, t),
		G: lerpf64(c.G, to.G, t),
		B: lerpf64(c.B, to.B, t),
		A: lerpf64(c.A, to.A, t),
	}
}

func clampf64(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func lerpf64(from, to, t float64) float64 {
	return from + (to-from)*t
}
