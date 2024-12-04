package mathf

import (
	"github.com/godot-ext/mathf/impl"
)

type Vector4i struct {
	X, Y, Z, W Int
}

func (v Vector4i) toImplf() impl.Vector4 {
	return impl.Vector4{impl.Float(v.X), impl.Float(v.Y), impl.Float(v.Z), impl.Float(v.W)}
}

func (v Vector4i) fromImplf(iv impl.Vector4) Vector4i {
	return NewVector4i(int(iv[0]), int(iv[1]), int(iv[2]), int(iv[3]))
}

func (v Vector4i) toImpl() impl.Vector4i {
	return impl.Vector4i{int32(v.X), int32(v.Y), int32(v.Z), int32(v.W)}
}

func (v Vector4i) fromImpl(iv impl.Vector4i) Vector4i {
	return NewVector4i(int(iv[0]), int(iv[1]), int(iv[2]), int(iv[3]))
}

func NewVector4i(x, y, z, w int) Vector4i {
	return Vector4i{Int(x), Int(y), Int(z), Int(w)}
}

func (v Vector4i) Add(other Vector4i) Vector4i {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vector4i) Sub(other Vector4i) Vector4i {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vector4i) Mul(other Vector4i) Vector4i {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vector4i) Div(other Vector4i) Vector4i {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vector4i) Addi(i int) Vector4i {
	return v.fromImpl(v.toImpl().Addi(int64(i)))
}

func (v Vector4i) Subi(i int) Vector4i {
	return v.fromImpl(v.toImpl().Subi(int64(i)))
}

func (v Vector4i) Muli(i int) Vector4i {
	return v.fromImpl(v.toImpl().Muli(int64(i)))
}

func (v Vector4i) Divi(i int) Vector4i {
	return v.fromImpl(v.toImpl().Divi(int64(i)))
}

func (v Vector4i) Dot(other Vector4i) float64 {
	return v.toImplf().Dot(other.toImplf())
}

func (v Vector4i) Length() float64 {
	return v.toImplf().Length()
}

func (v Vector4i) LengthSquared() float64 {
	return v.toImplf().LengthSquared()
}

func (v Vector4i) DistanceTo(other Vector4i) float64 {
	return v.toImplf().DistanceTo(other.toImplf())
}

func (v Vector4i) DistanceSquaredTo(other Vector4i) float64 {
	return v.toImplf().DistanceSquaredTo(other.toImplf())
}

func (v Vector4i) Abs() Vector4i {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vector4i) Sign() Vector4i {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vector4i) Clamp(min, max Vector4i) Vector4i {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vector4i) Neg() Vector4i {
	return v.fromImpl(v.toImpl().Neg())
}
