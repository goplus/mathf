package mathf

import (
	"godot-ext/mathf/impl"
	"unsafe"
)

type Vector2i struct {
	X, Y Int
}

func (v *Vector2i) toImplf() impl.Vector2 {
	vec := NewVector2(float64(v.X), float64(v.Y))
	return *(*impl.Vector2)(unsafe.Pointer(&vec))
}

func (v *Vector2i) fromImplf(iv impl.Vector2) Vector2i {
	vec := *(*Vector2)(unsafe.Pointer(&iv))
	return Vector2i{Int(vec.X), Int(vec.Y)}
}
func (v Vector2i) toImpl() impl.Vector2i {
	return *(*impl.Vector2i)(unsafe.Pointer(&v))
}

func (v Vector2i) fromImpl(iv impl.Vector2i) Vector2i {
	return *(*Vector2i)(unsafe.Pointer(&iv))
}

func NewVector2i(x, y int) Vector2i {
	return Vector2i{Int(x), Int(y)}
}

func (v Vector2i) Add(other Vector2i) Vector2i {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vector2i) Sub(other Vector2i) Vector2i {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vector2i) Mul(other Vector2i) Vector2i {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vector2i) Div(other Vector2i) Vector2i {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vector2i) Addi(i int) Vector2i {
	return v.fromImpl(v.toImpl().Addi(int64(i)))
}

func (v Vector2i) Subi(i int) Vector2i {
	return v.fromImpl(v.toImpl().Subi(int64(i)))
}

func (v Vector2i) Muli(i int) Vector2i {
	return v.fromImpl(v.toImpl().Muli(int64(i)))
}

func (v Vector2i) Divi(i int) Vector2i {
	return v.fromImpl(v.toImpl().Divi(int64(i)))
}

func (v Vector2i) Dot(other Vector2i) float64 {
	return v.toImplf().Dot(other.toImplf())
}

func (v Vector2i) Length() float64 {
	return v.toImpl().Length()
}

func (v Vector2i) LengthSquared() int {
	return int(v.toImpl().LengthSquared())
}

func (v Vector2i) DistanceTo(other Vector2i) float64 {
	return v.toImplf().DistanceTo(other.toImplf())
}

func (v Vector2i) DistanceSquaredTo(other Vector2i) int {
	return int(v.toImplf().DistanceSquaredTo(other.toImplf()))
}

func (v Vector2i) Abs() Vector2i {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vector2i) Sign() Vector2i {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vector2i) Clamp(min, max Vector2i) Vector2i {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vector2i) Neg() Vector2i {
	return v.fromImpl(v.toImpl().Neg())
}
