package mathf

import (
	"godot-ext/mathf/impl"
	"unsafe"
)

type Vector2 struct {
	X, Y Float
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{Float(x), Float(y)}
}

func (v *Vector2) toImpl() impl.Vector2 {
	return *(*impl.Vector2)(unsafe.Pointer(v))
}

func (v *Vector2) fromImpl(iv impl.Vector2) Vector2 {
	return *(*Vector2)(unsafe.Pointer(&iv))
}

func (v Vector2) Add(other Vector2) Vector2 {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vector2) Mul(other Vector2) Vector2 {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vector2) Div(other Vector2) Vector2 {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vector2) Addf(f float64) Vector2 {
	return v.fromImpl(v.toImpl().Addf(f))
}

func (v Vector2) Subf(f float64) Vector2 {
	return v.fromImpl(v.toImpl().Subf(f))
}

func (v Vector2) Mulf(f float64) Vector2 {
	return v.fromImpl(v.toImpl().Mulf(f))
}

func (v Vector2) Divf(f float64) Vector2 {
	return v.fromImpl(v.toImpl().Divf(f))
}

func (v Vector2) Cross(other Vector2) float64 {
	return v.toImpl().Cross(other.toImpl())
}

func (v Vector2) Dot(other Vector2) float64 {
	return v.toImpl().Dot(other.toImpl())
}

func (v Vector2) Length() float64 {
	return v.toImpl().Length()
}

func (v Vector2) LengthSquared() float64 {
	return v.toImpl().LengthSquared()
}

func (v Vector2) DistanceTo(other Vector2) float64 {
	return v.toImpl().DistanceTo(other.toImpl())
}

func (v Vector2) DistanceSquaredTo(other Vector2) float64 {
	return v.toImpl().DistanceSquaredTo(other.toImpl())
}

func (v Vector2) Normalize() Vector2 {
	return v.fromImpl(v.toImpl().Normalized())
}

func (v Vector2) IsNormalized() bool {
	return v.toImpl().IsNormalized()
}

func (v Vector2) IsFinite() bool {
	return v.toImpl().IsFinite()
}

func (v Vector2) IsApproximatelyZero() bool {
	return v.toImpl().IsApproximatelyZero()
}

func (v Vector2) Abs() Vector2 {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vector2) Ceil() Vector2 {
	return v.fromImpl(v.toImpl().Ceil())
}

func (v Vector2) Floor() Vector2 {
	return v.fromImpl(v.toImpl().Floor())
}

func (v Vector2) Round() Vector2 {
	return v.fromImpl(v.toImpl().Round())
}

func (v Vector2) Sign() Vector2 {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vector2) Clamp(min, max Vector2) Vector2 {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vector2) Lerp(to Vector2, weight float64) Vector2 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v Vector2) Lerpf(to Vector2, weight float64) Vector2 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v Vector2) Neg() Vector2 {
	return v.fromImpl(v.toImpl().Neg())
}
