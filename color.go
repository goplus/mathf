package mathf

import (
	"fmt"
)

type Color struct {
	R, G, B, A float32
}

func NewColor(r, g, b, a float64) Color {
	return Color{
		R: float32(r),
		G: float32(g),
		B: float32(b),
		A: float32(a),
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

func (c Color) Mulf(f float32) Color {
	return Color{
		R: c.R * f,
		G: c.G * f,
		B: c.B * f,
		A: c.A * f,
	}
}

func (c Color) Clamp() Color {
	return Color{
		R: clampf32(c.R, 0, 1),
		G: clampf32(c.G, 0, 1),
		B: clampf32(c.B, 0, 1),
		A: clampf32(c.A, 0, 1),
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

func (c Color) Lerp(to Color, t float32) Color {
	return Color{
		R: lerpf32(c.R, to.R, t),
		G: lerpf32(c.G, to.G, t),
		B: lerpf32(c.B, to.B, t),
		A: lerpf32(c.A, to.A, t),
	}
}

func clampf32(v, min, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func lerpf32(from, to, t float32) float32 {
	return from + (to-from)*t
}
