package mathf

import (
	"github.com/godot-ext/mathf/impl"
)

type Vec4i struct {
	X, Y, Z, W Int
}

func (v Vec4i) toImplf() impl.Vector4 {
	return impl.Vector4{impl.Float(v.X), impl.Float(v.Y), impl.Float(v.Z), impl.Float(v.W)}
}

func (v Vec4i) fromImplf(iv impl.Vector4) Vec4i {
	return NewVec4i(int(iv[0]), int(iv[1]), int(iv[2]), int(iv[3]))
}

func (v Vec4i) toImpl() impl.Vector4i {
	return impl.Vector4i{int32(v.X), int32(v.Y), int32(v.Z), int32(v.W)}
}

func (v Vec4i) fromImpl(iv impl.Vector4i) Vec4i {
	return NewVec4i(int(iv[0]), int(iv[1]), int(iv[2]), int(iv[3]))
}

func NewVec4i(x, y, z, w int) Vec4i {
	return Vec4i{Int(x), Int(y), Int(z), Int(w)}
}

func (v Vec4i) Add(other Vec4i) Vec4i {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vec4i) Sub(other Vec4i) Vec4i {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vec4i) Mul(other Vec4i) Vec4i {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vec4i) Div(other Vec4i) Vec4i {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vec4i) Addi(i int) Vec4i {
	return v.fromImpl(v.toImpl().Addi(int64(i)))
}

func (v Vec4i) Subi(i int) Vec4i {
	return v.fromImpl(v.toImpl().Subi(int64(i)))
}

func (v Vec4i) Muli(i int) Vec4i {
	return v.fromImpl(v.toImpl().Muli(int64(i)))
}

func (v Vec4i) Divi(i int) Vec4i {
	return v.fromImpl(v.toImpl().Divi(int64(i)))
}

func (v Vec4i) Dot(other Vec4i) float64 {
	return v.toImplf().Dot(other.toImplf())
}

func (v Vec4i) Length() float64 {
	return v.toImplf().Length()
}

func (v Vec4i) LengthSquared() float64 {
	return v.toImplf().LengthSquared()
}

func (v Vec4i) DistanceTo(other Vec4i) float64 {
	return v.toImplf().DistanceTo(other.toImplf())
}

func (v Vec4i) DistanceSquaredTo(other Vec4i) float64 {
	return v.toImplf().DistanceSquaredTo(other.toImplf())
}

func (v Vec4i) Abs() Vec4i {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vec4i) Sign() Vec4i {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vec4i) Clamp(min, max Vec4i) Vec4i {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vec4i) Neg() Vec4i {
	return v.fromImpl(v.toImpl().Neg())
}
