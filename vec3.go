package mathf

import (
	"github.com/godot-ext/mathf/impl"
)

type Vector3 struct {
	X, Y, Z Float
}

func (v Vector3) toImpl() impl.Vector3 {
	return impl.NewVector3(float64(v.X), float64(v.Y), float64(v.Z))
}

func (v Vector3) fromImpl(iv impl.Vector3) Vector3 {
	return NewVector3(iv.X(), iv.Y(), iv.Z())
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{Float(x), Float(y), Float(z)}
}

func NewVector3FromImpl(iv impl.Vector3) Vector3 {
	return NewVector3(iv.X(), iv.Y(), iv.Z())
}

func (v *Vector3) Add(other Vector3) Vector3 {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v *Vector3) Sub(other Vector3) Vector3 {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v *Vector3) Mul(other Vector3) Vector3 {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v *Vector3) Div(other Vector3) Vector3 {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v *Vector3) Addf(f float64) Vector3 {
	return v.fromImpl(v.toImpl().Addf(f))
}

func (v *Vector3) Subf(f float64) Vector3 {
	return v.fromImpl(v.toImpl().Subf(f))
}

func (v *Vector3) Mulf(f float64) Vector3 {
	return v.fromImpl(v.toImpl().Mulf(f))
}

func (v *Vector3) Divf(f float64) Vector3 {
	return v.fromImpl(v.toImpl().Divf(f))
}

func (v *Vector3) Cross(other Vector3) Vector3 {
	return v.fromImpl(v.toImpl().Cross(other.toImpl()))
}

func (v *Vector3) Dot(other Vector3) float64 {
	return v.toImpl().Dot(other.toImpl())
}

func (v *Vector3) Length() float64 {
	return v.toImpl().Length()
}

func (v *Vector3) LengthSquared() float64 {
	return v.toImpl().LengthSquared()
}

func (v *Vector3) DistanceTo(other Vector3) float64 {
	return v.toImpl().DistanceTo(other.toImpl())
}

func (v *Vector3) DistanceSquaredTo(other Vector3) float64 {
	return v.toImpl().DistanceSquaredTo(other.toImpl())
}

func (v *Vector3) Normalize() Vector3 {
	return v.fromImpl(v.toImpl().Normalized())
}

func (v *Vector3) IsNormalized() bool {
	return v.toImpl().IsNormalized()
}

func (v *Vector3) IsFinite() bool {
	return v.toImpl().IsFinite()
}

func (v *Vector3) IsApproximatelyZero() bool {
	return v.toImpl().IsApproximatelyZero()
}

func (v *Vector3) Abs() Vector3 {
	return v.fromImpl(v.toImpl().Abs())
}

func (v *Vector3) Ceil() Vector3 {
	return v.fromImpl(v.toImpl().Ceil())
}

func (v *Vector3) Floor() Vector3 {
	return v.fromImpl(v.toImpl().Floor())
}

func (v *Vector3) Round() Vector3 {
	return v.fromImpl(v.toImpl().Round())
}

func (v *Vector3) Clamp(min, max Vector3) Vector3 {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v *Vector3) Lerp(to Vector3, weight float64) Vector3 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vector3) Lerpf(to Vector3, weight float64) Vector3 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vector3) Neg() Vector3 {
	return v.fromImpl(v.toImpl().Neg())
}
