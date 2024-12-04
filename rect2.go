package mathf

import (
	"godot-ext/mathf/impl"
)

type Side = impl.Side

type Rect2 struct {
	Position Vector2
	Size     Vector2
}

func (v Rect2) toImpl() impl.Rect2 {
	return impl.Rect2{
		Position: v.Position.toImpl(),
		Size:     v.Size.toImpl(),
	}
}

func (v Rect2) fromImpl(iv impl.Rect2) Rect2 {
	return Rect2{
		Position: NewVector2(float64(iv.Position[0]), float64(iv.Position[1])),
		Size:     NewVector2(float64(iv.Size[0]), float64(iv.Size[1])),
	}
}

func NewRect2(x, y, w, h float64) Rect2 {
	return Rect2{
		Position: NewVector2(x, y),
		Size:     NewVector2(w, h),
	}
}

func (r *Rect2) End() Vector2 {
	impl := r.toImpl()
	end := impl.End()
	return NewVector2(float64(end[0]), float64(end[1]))
}

func (r *Rect2) SetEnd(end Vector2) {
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

func (r Rect2) Expand(to Vector2) Rect2 {
	return r.fromImpl(r.toImpl().Expand(to.toImpl()))
}

func (r Rect2) Area() float64 {
	return r.toImpl().Area()
}

func (r Rect2) Center() Vector2 {
	impl := r.toImpl()
	center := impl.Center()
	return NewVector2(float64(center[0]), float64(center[1]))
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

func (r Rect2) HasPoint(point Vector2) bool {
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
