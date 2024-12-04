package mathf

import (
	"godot-ext/mathf/impl"
	"unsafe"
)

type Vector3i struct {
	X, Y, Z Int
}

func (v *Vector3i) toImplf() impl.Vector3 {
	vec := NewVector3(float64(v.X), float64(v.Y), float64(v.Z))
	return *(*impl.Vector3)(unsafe.Pointer(&vec))
}

func (v *Vector3i) fromImplf(iv impl.Vector3) Vector3i {
	vec := *(*Vector3)(unsafe.Pointer(&iv))
	return Vector3i{Int(vec.X), Int(vec.Y), Int(vec.Z)}
}
func (v Vector3i) toImpl() impl.Vector3i {
	return *(*impl.Vector3i)(unsafe.Pointer(&v)) //
}

func (v Vector3i) fromImpl(iv impl.Vector3i) Vector3i {
	return *(*Vector3i)(unsafe.Pointer(&iv))
}

func NewVector3i(x, y, z int) Vector3i {
	return Vector3i{Int(x), Int(y), Int(z)}
}

func (v Vector3i) Add(other Vector3i) Vector3i {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vector3i) Sub(other Vector3i) Vector3i {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vector3i) Mul(other Vector3i) Vector3i {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vector3i) Div(other Vector3i) Vector3i {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vector3i) Addi(i int) Vector3i {
	return v.fromImpl(v.toImpl().Addi(int64(i)))
}

func (v Vector3i) Subi(i int) Vector3i {
	return v.fromImpl(v.toImpl().Subi(int64(i)))
}

func (v Vector3i) Muli(i int) Vector3i {
	return v.fromImpl(v.toImpl().Muli(int64(i)))
}

func (v Vector3i) Divi(i int) Vector3i {
	return v.fromImpl(v.toImpl().Divi(int64(i)))
}

func (v Vector3i) Dot(other Vector3i) float64 {
	return v.toImplf().Dot(other.toImplf())
}

func (v Vector3i) Length() float64 {
	return v.toImplf().Length()
}

func (v Vector3i) LengthSquared() float64 {
	return v.toImplf().LengthSquared()
}

func (v Vector3i) DistanceTo(other Vector3i) float64 {
	return v.toImplf().DistanceTo(other.toImplf())
}

func (v Vector3i) DistanceSquaredTo(other Vector3i) float64 {
	return v.toImplf().DistanceSquaredTo(other.toImplf())
}

func (v Vector3i) Abs() Vector3i {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vector3i) Sign() Vector3i {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vector3i) Clamp(min, max Vector3i) Vector3i {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vector3i) Neg() Vector3i {
	return v.fromImpl(v.toImpl().Neg())
}

func (v Vector3i) Cross(other Vector3i) Vector3i {
	return Vector3i{
		X: Int(v.Y*other.Z - v.Z*other.Y),
		Y: Int(v.Z*other.X - v.X*other.Z),
		Z: Int(v.X*other.Y - v.Y*other.X),
	}
}
