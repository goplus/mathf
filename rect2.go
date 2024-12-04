package mathf

import (
	"fmt"
	"unsafe"

	"github.com/godot-ext/mathf/impl"
)

type Side = impl.Side

type Rect2 struct {
	Position Vec2
	Size     Vec2
}

func (v Rect2) toImpl() impl.Rect2 {
	return *(*impl.Rect2)(unsafe.Pointer(&v))
}

func (v Rect2) fromImpl(iv impl.Rect2) Rect2 {
	return *(*Rect2)(unsafe.Pointer(&iv))
}

func NewRect2(x, y, w, h float64) Rect2 {
	return Rect2{
		Position: NewVec2(x, y),
		Size:     NewVec2(w, h),
	}
}

func (r *Rect2) End() Vec2 {
	impl := r.toImpl()
	end := impl.End()
	return NewVec2(float64(end[0]), float64(end[1]))
}

func (r *Rect2) SetEnd(end Vec2) {
	impl := r.toImpl()
	impl.SetEnd(end.toImpl())
	*r = r.fromImpl(impl)
}

func (r Rect2) Rect2i() Rect2i {
	return NewRect2i(int(r.Position.X), int(r.Position.Y), int(r.Size.X), int(r.Size.Y))
}

func (r Rect2) Abs() Rect2 {
	return r.fromImpl(r.toImpl().Abs())
}

func (r Rect2) Encloses(b Rect2) bool {
	return r.toImpl().Encloses(b.toImpl())
}

func (r Rect2) Expand(to Vec2) Rect2 {
	return r.fromImpl(r.toImpl().Expand(to.toImpl()))
}

func (r Rect2) Area() float64 {
	return r.toImpl().Area()
}

func (r Rect2) Center() Vec2 {
	impl := r.toImpl()
	center := impl.Center()
	return NewVec2(float64(center[0]), float64(center[1]))
}

func (r Rect2) Grow(amount float64) Rect2 {
	return r.fromImpl(r.toImpl().Grow(amount))
}

func (r Rect2) GrowIndividual(left, top, right, bottom float64) Rect2 {
	return r.fromImpl(r.toImpl().GrowIndividual(left, top, right, bottom))
}

func (r Rect2) GrowSide(side Side, amount float64) Rect2 {
	return r.fromImpl(r.toImpl().GrowSide(side, amount))
}

func (r Rect2) HasArea() bool {
	return r.toImpl().HasArea()
}

func (r Rect2) HasPoint(point Vec2) bool {
	return r.toImpl().HasPoint(point.toImpl())
}

func (r Rect2) Intersection(b Rect2) Rect2 {
	return r.fromImpl(r.toImpl().Intersection(b.toImpl()))
}

func (r Rect2) Intersects(b Rect2, includeBorders bool) bool {
	return r.toImpl().Intersects(b.toImpl(), includeBorders)
}

func (r Rect2) Merge(b Rect2) Rect2 {
	return r.fromImpl(r.toImpl().Merge(b.toImpl()))
}

func (r Rect2) String() string {
	return fmt.Sprintf("[pos=(%g, %g), size=(%g, %g)]", r.Position.X, r.Position.Y, r.Size.X, r.Size.Y)
}

func minf(a, b Float) Float {
	if a < b {
		return a
	}
	return b
}

func maxf(a, b Float) Float {
	if a > b {
		return a
	}
	return b
}
