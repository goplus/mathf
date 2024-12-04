package mathf

import (
	"godot-ext/mathf/impl"
	"unsafe"
)

type Vector4 struct {
	X, Y, Z, W Float
}

func NewVector4(x, y, z, w float64) Vector4 {
	return Vector4{Float(x), Float(y), Float(z), Float(w)}
}

func (v *Vector4) toImpl() impl.Vector4 {
	return *(*impl.Vector4)(unsafe.Pointer(v)) //
}

func (v *Vector4) fromImpl(iv impl.Vector4) Vector4 {
	return *(*Vector4)(unsafe.Pointer(&iv))
}

func (v *Vector4) Add(other Vector4) Vector4 {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v *Vector4) Sub(other Vector4) Vector4 {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v *Vector4) Mul(other Vector4) Vector4 {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v *Vector4) Div(other Vector4) Vector4 {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v *Vector4) Addf(f float64) Vector4 {
	return v.fromImpl(v.toImpl().Addf(f))
}

func (v *Vector4) Subf(f float64) Vector4 {
	return v.fromImpl(v.toImpl().Subf(f))
}

func (v *Vector4) Mulf(f float64) Vector4 {
	return v.fromImpl(v.toImpl().Mulf(f))
}

func (v *Vector4) Divf(f float64) Vector4 {
	return v.fromImpl(v.toImpl().Divf(f))
}

func (v *Vector4) Dot(other Vector4) float64 {
	return v.toImpl().Dot(other.toImpl())
}

func (v *Vector4) Length() float64 {
	return v.toImpl().Length()
}

func (v *Vector4) LengthSquared() float64 {
	return v.toImpl().LengthSquared()
}

func (v *Vector4) DistanceTo(other Vector4) float64 {
	return v.toImpl().DistanceTo(other.toImpl())
}

func (v *Vector4) DistanceSquaredTo(other Vector4) float64 {
	return v.toImpl().DistanceSquaredTo(other.toImpl())
}

func (v Vector4) Normalize() Vector4 {
	l := v.Length()
	if l == 0 {
		return v
	}
	return Vector4{
		X: Float(float64(v.X) / l),
		Y: Float(float64(v.Y) / l),
		Z: Float(float64(v.Z) / l),
		W: Float(float64(v.W) / l),
	}
}

func (v *Vector4) IsNormalized() bool {
	return v.toImpl().IsNormalized()
}

func (v *Vector4) IsFinite() bool {
	return v.toImpl().IsFinite()
}

func (v *Vector4) IsApproximatelyZero() bool {
	return v.toImpl().IsApproximatelyZero()
}

func (v *Vector4) Abs() Vector4 {
	return v.fromImpl(v.toImpl().Abs())
}

func (v *Vector4) Ceil() Vector4 {
	return v.fromImpl(v.toImpl().Ceil())
}

func (v *Vector4) Floor() Vector4 {
	return v.fromImpl(v.toImpl().Floor())
}

func (v *Vector4) Round() Vector4 {
	return v.fromImpl(v.toImpl().Round())
}

func (v *Vector4) Sign() Vector4 {
	return v.fromImpl(v.toImpl().Sign())
}

func (v *Vector4) Clamp(min, max Vector4) Vector4 {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v *Vector4) Lerp(to Vector4, weight float64) Vector4 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vector4) Lerpf(to Vector4, weight float64) Vector4 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vector4) Neg() Vector4 {
	return v.fromImpl(v.toImpl().Neg())
}
