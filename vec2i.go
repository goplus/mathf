package mathf

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/godot-ext/mathf/impl"
)

type Vec2i struct {
	X, Y Int
}

func (v *Vec2i) toImplf() impl.Vector2 {
	return *(*impl.Vector2)(unsafe.Pointer(v))
}

func (v *Vec2i) fromImplf(iv impl.Vector2) Vec2i {
	return *(*Vec2i)(unsafe.Pointer(&iv))
}

func (v Vec2i) toImpl() impl.Vector2i {
	return *(*impl.Vector2i)(unsafe.Pointer(&v))
}

func (v Vec2i) fromImpl(iv impl.Vector2i) Vec2i {
	return *(*Vec2i)(unsafe.Pointer(&iv))
}

func NewVec2i(x, y int) Vec2i {
	return Vec2i{Int(x), Int(y)}
}

func (v Vec2i) Add(other Vec2i) Vec2i {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vec2i) Sub(other Vec2i) Vec2i {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vec2i) Mul(other Vec2i) Vec2i {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vec2i) Div(other Vec2i) Vec2i {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vec2i) Addi(i int) Vec2i {
	return v.fromImpl(v.toImpl().Addi(int64(i)))
}

func (v Vec2i) Subi(i int) Vec2i {
	return v.fromImpl(v.toImpl().Subi(int64(i)))
}

func (v Vec2i) Muli(i int) Vec2i {
	return v.fromImpl(v.toImpl().Muli(int64(i)))
}

func (v Vec2i) Divi(i int) Vec2i {
	return v.fromImpl(v.toImpl().Divi(int64(i)))
}

func (v Vec2i) Dot(other Vec2i) float64 {
	return float64(v.X*other.X + v.Y*other.Y)
}

func (v Vec2i) Length() float64 {
	return float64(math.Sqrt(float64(v.LengthSquared())))
}

func (v Vec2i) LengthSquared() float64 {
	return float64(v.X*v.X + v.Y*v.Y)
}

func (v Vec2i) DistanceTo(other Vec2i) float64 {
	return float64(math.Sqrt(float64(v.DistanceSquaredTo(other))))
}

func (v Vec2i) DistanceSquaredTo(other Vec2i) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float64(dx*dx + dy*dy)
}

func (v Vec2i) Abs() Vec2i {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vec2i) Sign() Vec2i {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vec2i) Clamp(min, max Vec2i) Vec2i {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vec2i) Neg() Vec2i {
	return v.fromImpl(v.toImpl().Neg())
}

func (v Vec2i) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}
