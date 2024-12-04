package mathf

import (
	"fmt"
)

type AABB struct {
	Position Vec3
	Size     Vec3
}

func (a AABB) String() string {
	return fmt.Sprintf("[pos=(%g, %g, %g), size=(%g, %g, %g)]",
		a.Position.X, a.Position.Y, a.Position.Z,
		a.Size.X, a.Size.Y, a.Size.Z)
}
