package mathf

import (
	"unsafe"

	"github.com/godot-ext/mathf/impl"
)

type Vec4 struct {
	X, Y, Z, W Float
}

func NewVec4(x, y, z, w float64) Vec4 {
	return Vec4{Float(x), Float(y), Float(z), Float(w)}
}

func (v *Vec4) toImpl() impl.Vector4 {
	return *(*impl.Vector4)(unsafe.Pointer(v)) //
}

func (v *Vec4) fromImpl(iv impl.Vector4) Vec4 {
	return *(*Vec4)(unsafe.Pointer(&iv))
}

func (v *Vec4) Add(other Vec4) Vec4 {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v *Vec4) Sub(other Vec4) Vec4 {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v *Vec4) Mul(other Vec4) Vec4 {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v *Vec4) Div(other Vec4) Vec4 {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v *Vec4) Addf(f float64) Vec4 {
	return v.fromImpl(v.toImpl().Addf(f))
}

func (v *Vec4) Subf(f float64) Vec4 {
	return v.fromImpl(v.toImpl().Subf(f))
}

func (v *Vec4) Mulf(f float64) Vec4 {
	return v.fromImpl(v.toImpl().Mulf(f))
}

func (v *Vec4) Divf(f float64) Vec4 {
	return v.fromImpl(v.toImpl().Divf(f))
}

func (v *Vec4) Dot(other Vec4) float64 {
	return v.toImpl().Dot(other.toImpl())
}

func (v *Vec4) Length() float64 {
	return v.toImpl().Length()
}

func (v *Vec4) LengthSquared() float64 {
	return v.toImpl().LengthSquared()
}

func (v *Vec4) DistanceTo(other Vec4) float64 {
	return v.toImpl().DistanceTo(other.toImpl())
}

func (v *Vec4) DistanceSquaredTo(other Vec4) float64 {
	return v.toImpl().DistanceSquaredTo(other.toImpl())
}

func (v Vec4) Normalize() Vec4 {
	l := v.Length()
	if l == 0 {
		return v
	}
	return Vec4{
		X: Float(float64(v.X) / l),
		Y: Float(float64(v.Y) / l),
		Z: Float(float64(v.Z) / l),
		W: Float(float64(v.W) / l),
	}
}

func (v *Vec4) IsNormalized() bool {
	return v.toImpl().IsNormalized()
}

func (v *Vec4) IsFinite() bool {
	return v.toImpl().IsFinite()
}

func (v *Vec4) IsApproximatelyZero() bool {
	return v.toImpl().IsApproximatelyZero()
}

func (v *Vec4) Abs() Vec4 {
	return v.fromImpl(v.toImpl().Abs())
}

func (v *Vec4) Ceil() Vec4 {
	return v.fromImpl(v.toImpl().Ceil())
}

func (v *Vec4) Floor() Vec4 {
	return v.fromImpl(v.toImpl().Floor())
}

func (v *Vec4) Round() Vec4 {
	return v.fromImpl(v.toImpl().Round())
}

func (v *Vec4) Sign() Vec4 {
	return v.fromImpl(v.toImpl().Sign())
}

func (v *Vec4) Clamp(min, max Vec4) Vec4 {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v *Vec4) Lerp(to Vec4, weight float64) Vec4 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vec4) Lerpf(to Vec4, weight float64) Vec4 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vec4) Neg() Vec4 {
	return v.fromImpl(v.toImpl().Neg())
}
