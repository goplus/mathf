package mathf

import (
	"fmt"
	"unsafe"

	"github.com/goplus/mathf/impl"
)

type Vec3 struct {
	X, Y, Z Float
}

func (v Vec3) toImpl() impl.Vector3 {
	return *(*impl.Vector3)(unsafe.Pointer(&v))
}

func (v Vec3) fromImpl(iv impl.Vector3) Vec3 {
	return *(*Vec3)(unsafe.Pointer(&iv))
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{Float(x), Float(y), Float(z)}
}

func NewVec3FromImpl(iv impl.Vector3) Vec3 {
	return NewVec3(iv.X(), iv.Y(), iv.Z())
}

func (v *Vec3) Add(other Vec3) Vec3 {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v *Vec3) Sub(other Vec3) Vec3 {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v *Vec3) Mul(other Vec3) Vec3 {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v *Vec3) Div(other Vec3) Vec3 {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v *Vec3) Addf(f float64) Vec3 {
	return v.fromImpl(v.toImpl().Addf(f))
}

func (v *Vec3) Subf(f float64) Vec3 {
	return v.fromImpl(v.toImpl().Subf(f))
}

func (v *Vec3) Mulf(f float64) Vec3 {
	return v.fromImpl(v.toImpl().Mulf(f))
}

func (v *Vec3) Divf(f float64) Vec3 {
	return v.fromImpl(v.toImpl().Divf(f))
}

func (v *Vec3) Cross(other Vec3) Vec3 {
	return v.fromImpl(v.toImpl().Cross(other.toImpl()))
}

func (v *Vec3) Dot(other Vec3) float64 {
	return v.toImpl().Dot(other.toImpl())
}

func (v *Vec3) Length() float64 {
	return v.toImpl().Length()
}

func (v *Vec3) LengthSquared() float64 {
	return v.toImpl().LengthSquared()
}

func (v *Vec3) DistanceTo(other Vec3) float64 {
	return v.toImpl().DistanceTo(other.toImpl())
}

func (v *Vec3) DistanceSquaredTo(other Vec3) float64 {
	return v.toImpl().DistanceSquaredTo(other.toImpl())
}

func (v *Vec3) Normalize() Vec3 {
	return v.fromImpl(v.toImpl().Normalized())
}

func (v *Vec3) IsNormalized() bool {
	return v.toImpl().IsNormalized()
}

func (v *Vec3) IsFinite() bool {
	return v.toImpl().IsFinite()
}

func (v *Vec3) IsApproximatelyZero() bool {
	return v.toImpl().IsApproximatelyZero()
}

func (v *Vec3) Abs() Vec3 {
	return v.fromImpl(v.toImpl().Abs())
}

func (v *Vec3) Ceil() Vec3 {
	return v.fromImpl(v.toImpl().Ceil())
}

func (v *Vec3) Floor() Vec3 {
	return v.fromImpl(v.toImpl().Floor())
}

func (v *Vec3) Round() Vec3 {
	return v.fromImpl(v.toImpl().Round())
}

func (v *Vec3) Clamp(min, max Vec3) Vec3 {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v *Vec3) Lerp(to Vec3, weight float64) Vec3 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vec3) Lerpf(to Vec3, weight float64) Vec3 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v *Vec3) Neg() Vec3 {
	return v.fromImpl(v.toImpl().Neg())
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%g, %g, %g)", v.X, v.Y, v.Z)
}
