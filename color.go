package mathf

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"golang.org/x/image/colornames"
)

var (
	errInvalidColorFormat     = errors.New("invalid color format")
	errUnsupportedColorFormat = errors.New("unsupported color format")
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

// New a color, s can be int, float64 or string
func NewColorAny(s interface{}) (Color, error) {
	return parseColor(s)
}

func NewColorHSV(h, s, v float64) Color {
	c := Color{}
	c.FromHSV(h, s, v)
	return c
}

func NewColorRGBi(r, g, b uint8) Color {
	return NewColor(float64(r)/255, float64(g)/255, float64(b)/255, 1)
}

func NewColorRGBAi(r, g, b, a uint8) Color {
	return NewColor(float64(r)/255, float64(g)/255, float64(b)/255, float64(a)/255)
}

func NewColorRGB(r, g, b float64) Color {
	return NewColor(r, g, b, 1)
}

func NewColorRGBA(r, g, b, a float64) Color {
	return NewColor(r, g, b, a)
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
		R: Clamp01f(c.R),
		G: Clamp01f(c.G),
		B: Clamp01f(c.B),
		A: Clamp01f(c.A),
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
	t = Clamp01f(t)
	return Color{
		R: Lerpf(c.R, to.R, t),
		G: Lerpf(c.G, to.G, t),
		B: Lerpf(c.B, to.B, t),
		A: Lerpf(c.A, to.A, t),
	}
}

// HSV2RGB converts hue (0-360), saturation (0-1), and brightness (0-1) to RGB.
func (c *Color) FromHSV(h, s, v float64) {
	var r, g, b float64
	h = math.Mod(h, 360)
	if h < 0 {
		h += 360
	}
	s = math.Max(0, math.Min(s, 1))
	v = math.Max(0, math.Min(v, 1))

	i := math.Floor(h / 60)
	f := (h / 60) - i
	p := v * (1 - s)
	q := v * (1 - (s * f))
	t := v * (1 - (s * (1 - f)))
	switch int(i) {
	case 0:
		r = v
		g = t
		b = p
	case 1:
		r = q
		g = v
		b = p
	case 2:
		r = p
		g = v
		b = t
	case 3:
		r = p
		g = q
		b = v
	case 4:
		r = t
		g = p
		b = v
	case 5:
		r = v
		g = p
		b = q
	}
	c.R = r
	c.G = g
	c.B = b
}

// RGB2HSV converts RGB to an array containing the hue, saturation, and brightness.
func (c *Color) ToHSV() (h, s, v float64) {
	var f, i float64
	r := float64(c.R)
	g := float64(c.G)
	b := float64(c.B)
	x := math.Min(math.Min(r, g), b)
	v = math.Max(math.Max(r, g), b)
	if x == v {
		return // gray; hue arbitrarily reported as zero
	}
	if r == x {
		f = g - b
		i = 3
	} else if g == x {
		f = b - r
		i = 5
	} else {
		f = r - g
		i = 1
	}
	h = math.Mod((i-(f/(v-x)))*60, 360)
	s = (v - x) / v
	return
}

// ScaleBrightness changes color brightness.
func (c *Color) ScaleBrightness(scale float64) {
	h, s, v := c.ToHSV()
	val := Clamp01f(scale * v)
	c.FromHSV(h, s, val)
}

func LerpColor(from Color, to Color, t float64) Color {
	return from.Lerp(to, t)
}

// Random returns a random color.
func RandomColor() Color {
	h := 360 * rand.Float64()
	s := 0.7 + (0.3 * rand.Float64())
	v := 0.6 + (0.4 * rand.Float64())
	return NewColorHSV(h, s, v)
}

// Parse, s can be int, float64 or string
func parseColor(s interface{}) (Color, error) {
	if s == nil {
		return Color{}, errors.New("color is nil")
	}
	if c, ok := s.(int); ok {
		return NewColorRGBAi(uint8(c>>16), uint8((c>>8)&0xff), uint8(c&0xff), 0xff), nil
	}
	if f, ok := s.(float64); ok {
		c := int(f)
		return NewColorRGBAi(uint8(c>>16), uint8((c>>8)&0xff), uint8(c&0xff), 0xff), nil
	}
	ss, ok := s.(string)
	if !ok {
		return Color{}, errUnsupportedColorFormat
	}
	return parseHexColor(ss)
}

func parseHexColor(s string) (Color, error) {
	c, ok := colornames.Map[s]
	if ok {
		clr := NewColorRGBAi(c.R, c.G, c.B, c.A)
		return clr, nil
	}

	if s == "" || s[0] != '#' {
		return Color{}, errInvalidColorFormat
	}
	var err error
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		default:
			err = errInvalidColorFormat
			return 0
		}
	}
	var r, g, b, a uint8 = 0, 0, 0, 0xff
	switch len(s) {
	case 7:
		r = hexToByte(s[1])<<4 + hexToByte(s[2])
		g = hexToByte(s[3])<<4 + hexToByte(s[4])
		b = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		r = hexToByte(s[1]) * 17
		g = hexToByte(s[2]) * 17
		b = hexToByte(s[3]) * 17
	default:
		err = errInvalidColorFormat
	}
	return NewColorRGBAi(r, g, b, a), err
}
