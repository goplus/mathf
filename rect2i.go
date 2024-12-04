package mathf

import (
	"fmt"
	"unsafe"

	"github.com/godot-ext/mathf/impl"
)

type Rect2i struct {
	Position Vec2i
	Size     Vec2i
}

func (v Rect2i) toImpl() impl.Rect2i {
	return *(*impl.Rect2i)(unsafe.Pointer(&v))
}

func (v Rect2i) fromImpl(iv impl.Rect2i) Rect2i {
	return *(*Rect2i)(unsafe.Pointer(&iv))
}

func (v *Rect2i) toImplf() impl.Rect2 {
	rect := NewRect2(float64(v.Position.X), float64(v.Position.Y), float64(v.Size.X), float64(v.Size.Y))
	return *(*impl.Rect2)(unsafe.Pointer(&rect))
}

func (v *Rect2i) fromImplf(iv impl.Rect2) Rect2i {
	rect := *(*Rect2)(unsafe.Pointer(&iv))
	return Rect2i{
		Position: Vec2i{Int(rect.Position.X), Int(rect.Position.Y)},
		Size:     Vec2i{Int(rect.Size.X), Int(rect.Size.Y)},
	}
}

func NewRect2i(x, y, w, h int) Rect2i {
	return Rect2i{
		Position: NewVec2i(x, y),
		Size:     NewVec2i(w, h),
	}
}

func (r *Rect2i) End() Vec2i {
	impl := r.toImpl()
	end := impl.End()
	return NewVec2i(int(end[0]), int(end[1]))
}

func (r *Rect2i) SetEnd(end Vec2i) {
	impl := r.toImpl()
	impl.SetEnd(end.toImpl())
	*r = r.fromImpl(impl)
}

func (r *Rect2i) Rect2() Rect2 {
	return NewRect2(float64(r.Position.X), float64(r.Position.Y), float64(r.Size.X), float64(r.Size.Y))
}

func (r *Rect2i) Abs() Rect2i {
	return r.fromImpl(r.toImpl().Abs())
}

func (r *Rect2i) Encloses(b Rect2i) bool {
	return r.toImpl().Encloses(b.toImpl())
}

func (r *Rect2i) Expand(to Vec2i) Rect2i {
	return r.fromImpl(r.toImpl().Expand(to.toImpl()))
}

func (r *Rect2i) Area() int64 {
	return r.toImpl().Area()
}

func (r *Rect2i) Center() Vec2i {
	impl := r.toImpl()
	center := impl.Center()
	return NewVec2i(int(center[0]), int(center[1]))
}

func (r *Rect2i) Grow(amount int64) Rect2i {
	return r.fromImpl(r.toImpl().Grow(amount))
}

func (r *Rect2i) GrowIndividual(left, top, right, bottom int64) Rect2i {
	return r.fromImpl(r.toImpl().GrowIndividual(left, top, right, bottom))
}

func (r *Rect2i) GrowSide(side Side, amount int64) Rect2i {
	return r.fromImpl(r.toImpl().GrowSide(side, amount))
}

func (r *Rect2i) HasArea() bool {
	return r.toImpl().HasArea()
}

func (r *Rect2i) HasPoint(point Vec2i) bool {
	return r.toImpl().HasPoint(point.toImpl())
}

func (r *Rect2i) Intersection(b Rect2i) Rect2i {
	return r.fromImpl(r.toImpl().Intersection(b.toImpl()))
}

func (r *Rect2i) Intersects(b Rect2i, includeBorders bool) bool {
	return r.toImpl().Intersects(b.toImpl(), includeBorders)
}

func (r *Rect2i) Merge(b Rect2i) Rect2i {
	return r.fromImpl(r.toImpl().Merge(b.toImpl()))
}

func (r Rect2i) String() string {
	return fmt.Sprintf("[pos=(%d, %d), size=(%d, %d)]", r.Position.X, r.Position.Y, r.Size.X, r.Size.Y)
}
