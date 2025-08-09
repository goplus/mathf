package vector3

import "github.com/goplus/mathf/impl"

func New(x, y, z float64) impl.Vector3                         { return impl.NewVector3(x, y, z) }
func Abs(v impl.Vector3) impl.Vector3                          { return v.Abs() }
func Ceil(v impl.Vector3) impl.Vector3                         { return v.Ceil() }
func Clamp(v impl.Vector3, min, max impl.Vector3) impl.Vector3 { return v.Clamp(min, max) }
func Cross(v, with impl.Vector3) impl.Vector3                  { return v.Cross(with) }
func Distance(v, to impl.Vector3) float64                      { return v.DistanceTo(to) }
func DistanceSquared(v, to impl.Vector3) float64               { return v.DistanceSquaredTo(to) }
func Dot(v, with impl.Vector3) float64                         { return v.Dot(with) }
func Floor(v impl.Vector3) impl.Vector3                        { return v.Floor() }
func IsFinite(v impl.Vector3) bool                             { return v.IsFinite() }
func IsNormalized(v impl.Vector3) bool                         { return v.IsNormalized() }
func IsApproximatelyZero(v impl.Vector3) bool                  { return v.IsApproximatelyZero() }
func Length(v impl.Vector3) float64                            { return v.Length() }
func LengthSquared(v impl.Vector3) float64                     { return v.LengthSquared() }
func Lerp(from, to impl.Vector3, weight float64) impl.Vector3  { return from.Lerp(to, weight) }
func Normalize(v impl.Vector3) impl.Vector3                    { return v.Normalized() }
func Round(v impl.Vector3) impl.Vector3                        { return v.Round() }

func Add(v, with impl.Vector3) impl.Vector3 { return v.Add(with) }
func Sub(v, with impl.Vector3) impl.Vector3 { return v.Sub(with) }
func Mul(v, with impl.Vector3) impl.Vector3 { return v.Mul(with) }
func Div(v, with impl.Vector3) impl.Vector3 { return v.Div(with) }

func Addf(v impl.Vector3, with float64) impl.Vector3 { return v.Addf(with) }
func Subf(v impl.Vector3, with float64) impl.Vector3 { return v.Subf(with) }
func Mulf(v impl.Vector3, with float64) impl.Vector3 { return v.Mulf(with) }
func Divf(v impl.Vector3, with float64) impl.Vector3 { return v.Divf(with) }

func Neg(v impl.Vector3) impl.Vector3 { return v.Neg() }

func Transform(v impl.Vector3, by impl.Transform3D) impl.Vector3 { return v.Transform(by) }
