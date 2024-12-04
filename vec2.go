package mathf

import (
	"fmt"
	"unsafe"

	"github.com/realdream-ai/mathf/impl"
)

type Vec2 struct {
	X, Y Float
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{Float(x), Float(y)}
}

func (v *Vec2) toImpl() impl.Vector2 {
	return *(*impl.Vector2)(unsafe.Pointer(v))
}

func (v *Vec2) fromImpl(iv impl.Vector2) Vec2 {
	return *(*Vec2)(unsafe.Pointer(&iv))
}

func (v Vec2) Add(other Vec2) Vec2 {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vec2) Mul(other Vec2) Vec2 {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vec2) Div(other Vec2) Vec2 {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vec2) Addf(f float64) Vec2 {
	return v.fromImpl(v.toImpl().Addf(f))
}

func (v Vec2) Subf(f float64) Vec2 {
	return v.fromImpl(v.toImpl().Subf(f))
}

func (v Vec2) Mulf(f float64) Vec2 {
	return v.fromImpl(v.toImpl().Mulf(f))
}

func (v Vec2) Divf(f float64) Vec2 {
	return v.fromImpl(v.toImpl().Divf(f))
}

func (v Vec2) Cross(other Vec2) float64 {
	return v.toImpl().Cross(other.toImpl())
}

func (v Vec2) Dot(other Vec2) float64 {
	return v.toImpl().Dot(other.toImpl())
}

func (v Vec2) Length() float64 {
	return v.toImpl().Length()
}

func (v Vec2) LengthSquared() float64 {
	return v.toImpl().LengthSquared()
}

func (v Vec2) DistanceTo(other Vec2) float64 {
	return v.toImpl().DistanceTo(other.toImpl())
}

func (v Vec2) DistanceSquaredTo(other Vec2) float64 {
	return v.toImpl().DistanceSquaredTo(other.toImpl())
}

func (v Vec2) Normalize() Vec2 {
	return v.fromImpl(v.toImpl().Normalized())
}

func (v Vec2) IsNormalized() bool {
	return v.toImpl().IsNormalized()
}

func (v Vec2) IsFinite() bool {
	return v.toImpl().IsFinite()
}

func (v Vec2) IsApproximatelyZero() bool {
	return v.toImpl().IsApproximatelyZero()
}

func (v Vec2) Abs() Vec2 {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vec2) Ceil() Vec2 {
	return v.fromImpl(v.toImpl().Ceil())
}

func (v Vec2) Floor() Vec2 {
	return v.fromImpl(v.toImpl().Floor())
}

func (v Vec2) Round() Vec2 {
	return v.fromImpl(v.toImpl().Round())
}

func (v Vec2) Sign() Vec2 {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vec2) Clamp(min, max Vec2) Vec2 {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vec2) Lerp(to Vec2, weight float64) Vec2 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v Vec2) Lerpf(to Vec2, weight float64) Vec2 {
	return v.fromImpl(v.toImpl().Lerp(to.toImpl(), weight))
}

func (v Vec2) Neg() Vec2 {
	return v.fromImpl(v.toImpl().Neg())
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%g, %g)", v.X, v.Y)
}
