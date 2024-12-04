package vector2

import "github.com/godot-ext/mathf/impl"

func New(x, y float64) impl.Vector2                            { return impl.NewVector2(x, y) }
func Abs(v impl.Vector2) impl.Vector2                          { return v.Abs() }
func Angle(v impl.Vector2) impl.Radians                        { return v.Angle() }
func Aspect(v impl.Vector2) float64                            { return v.Aspect() }
func Ceil(v impl.Vector2) impl.Vector2                         { return v.Ceil() }
func Clamp(v impl.Vector2, min, max impl.Vector2) impl.Vector2 { return v.Clamp(min, max) }
func Cross(v, with impl.Vector2) float64                       { return v.Cross(with) }
func Distance(v, to impl.Vector2) float64                      { return v.DistanceTo(to) }
func DistanceSquared(v, to impl.Vector2) float64               { return v.DistanceSquaredTo(to) }
func Dot(v, with impl.Vector2) float64                         { return v.Dot(with) }
func Floor(v impl.Vector2) impl.Vector2                        { return v.Floor() }
func IsFinite(v impl.Vector2) bool                             { return v.IsFinite() }
func IsNormalized(v impl.Vector2) bool                         { return v.IsNormalized() }
func IsApproximatelyZero(v impl.Vector2) bool                  { return v.IsApproximatelyZero() }
func Length(v impl.Vector2) float64                            { return v.Length() }
func LengthSquared(v impl.Vector2) float64                     { return v.LengthSquared() }
func Lerp(from, to impl.Vector2, weight float64) impl.Vector2  { return from.Lerp(to, weight) }
func Normalize(v impl.Vector2) impl.Vector2                    { return v.Normalized() }
func Round(v impl.Vector2) impl.Vector2                        { return v.Round() }
func Sign(v impl.Vector2) impl.Vector2                         { return v.Sign() }

func Add(v, with impl.Vector2) impl.Vector2 { return v.Add(with) }
func Sub(v, with impl.Vector2) impl.Vector2 { return v.Sub(with) }
func Mul(v, with impl.Vector2) impl.Vector2 { return v.Mul(with) }
func Div(v, with impl.Vector2) impl.Vector2 { return v.Div(with) }

func Addf(v impl.Vector2, with float64) impl.Vector2 { return v.Addf(with) }
func Subf(v impl.Vector2, with float64) impl.Vector2 { return v.Subf(with) }
func Mulf(v impl.Vector2, with float64) impl.Vector2 { return v.Mulf(with) }
func Divf(v impl.Vector2, with float64) impl.Vector2 { return v.Divf(with) }

func Neg(v impl.Vector2) impl.Vector2 { return v.Neg() }

func Transform(v impl.Vector2, by impl.Transform2D) impl.Vector2 { return v.Transform(by) }
