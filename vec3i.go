package mathf

import (
	"unsafe"

	"github.com/godot-ext/mathf/impl"
)

type Vec3i struct {
	X, Y, Z Int
}

func (v *Vec3i) toImplf() impl.Vector3 {
	vec := NewVec3(float64(v.X), float64(v.Y), float64(v.Z))
	return *(*impl.Vector3)(unsafe.Pointer(&vec))
}

func (v *Vec3i) fromImplf(iv impl.Vector3) Vec3i {
	vec := *(*Vec3)(unsafe.Pointer(&iv))
	return Vec3i{Int(vec.X), Int(vec.Y), Int(vec.Z)}
}
func (v Vec3i) toImpl() impl.Vector3i {
	return *(*impl.Vector3i)(unsafe.Pointer(&v)) //
}

func (v Vec3i) fromImpl(iv impl.Vector3i) Vec3i {
	return *(*Vec3i)(unsafe.Pointer(&iv))
}

func NewVec3i(x, y, z int) Vec3i {
	return Vec3i{Int(x), Int(y), Int(z)}
}

func (v Vec3i) Add(other Vec3i) Vec3i {
	return v.fromImpl(v.toImpl().Add(other.toImpl()))
}

func (v Vec3i) Sub(other Vec3i) Vec3i {
	return v.fromImpl(v.toImpl().Sub(other.toImpl()))
}

func (v Vec3i) Mul(other Vec3i) Vec3i {
	return v.fromImpl(v.toImpl().Mul(other.toImpl()))
}

func (v Vec3i) Div(other Vec3i) Vec3i {
	return v.fromImpl(v.toImpl().Div(other.toImpl()))
}

func (v Vec3i) Addi(i int) Vec3i {
	return v.fromImpl(v.toImpl().Addi(int64(i)))
}

func (v Vec3i) Subi(i int) Vec3i {
	return v.fromImpl(v.toImpl().Subi(int64(i)))
}

func (v Vec3i) Muli(i int) Vec3i {
	return v.fromImpl(v.toImpl().Muli(int64(i)))
}

func (v Vec3i) Divi(i int) Vec3i {
	return v.fromImpl(v.toImpl().Divi(int64(i)))
}

func (v Vec3i) Dot(other Vec3i) float64 {
	return v.toImplf().Dot(other.toImplf())
}

func (v Vec3i) Length() float64 {
	return v.toImplf().Length()
}

func (v Vec3i) LengthSquared() float64 {
	return v.toImplf().LengthSquared()
}

func (v Vec3i) DistanceTo(other Vec3i) float64 {
	return v.toImplf().DistanceTo(other.toImplf())
}

func (v Vec3i) DistanceSquaredTo(other Vec3i) float64 {
	return v.toImplf().DistanceSquaredTo(other.toImplf())
}

func (v Vec3i) Abs() Vec3i {
	return v.fromImpl(v.toImpl().Abs())
}

func (v Vec3i) Sign() Vec3i {
	return v.fromImpl(v.toImpl().Sign())
}

func (v Vec3i) Clamp(min, max Vec3i) Vec3i {
	return v.fromImpl(v.toImpl().Clamp(min.toImpl(), max.toImpl()))
}

func (v Vec3i) Neg() Vec3i {
	return v.fromImpl(v.toImpl().Neg())
}

func (v Vec3i) Cross(other Vec3i) Vec3i {
	return Vec3i{
		X: Int(v.Y*other.Z - v.Z*other.Y),
		Y: Int(v.Z*other.X - v.X*other.Z),
		Z: Int(v.X*other.Y - v.Y*other.X),
	}
}
