package mathf

import (
	"unsafe"

	"github.com/godot-ext/mathf/impl"
)

type Vec2i struct {
	X, Y Int
}

func (v *Vec2i) toImplf() impl.Vector2 {
	vec := NewVec2(float64(v.X), float64(v.Y))
	return *(*impl.Vector2)(unsafe.Pointer(&vec))
}

func (v *Vec2i) fromImplf(iv impl.Vector2) Vec2i {
	vec := *(*Vec2)(unsafe.Pointer(&iv))
	return Vec2i{Int(vec.X), Int(vec.Y)}
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
	return v.toImplf().Dot(other.toImplf())
}

func (v Vec2i) Length() float64 {
	return v.toImpl().Length()
}

func (v Vec2i) LengthSquared() int {
	return int(v.toImpl().LengthSquared())
}

func (v Vec2i) DistanceTo(other Vec2i) float64 {
	return v.toImplf().DistanceTo(other.toImplf())
}

func (v Vec2i) DistanceSquaredTo(other Vec2i) int {
	return int(v.toImplf().DistanceSquaredTo(other.toImplf()))
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
