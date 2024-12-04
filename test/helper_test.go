package test

import (
	"math"

	"github.com/godot-ext/mathf"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func vec2AlmostEqual(a, b mathf.Vector2) bool {
	return almostEqual(float64(a.X), float64(b.X)) && almostEqual(float64(a.Y), float64(b.Y))
}

func vec3AlmostEqual(a, b mathf.Vector3) bool {
	return almostEqual(float64(a.X), float64(b.X)) && almostEqual(float64(a.Y), float64(b.Y)) && almostEqual(float64(a.Z), float64(b.Z))
}

func vec4AlmostEqual(a, b mathf.Vector4) bool {
	return almostEqual(float64(a.X), float64(b.X)) && almostEqual(float64(a.Y), float64(b.Y)) && almostEqual(float64(a.Z), float64(b.Z)) && almostEqual(float64(a.W), float64(b.W))
}

func rect2AlmostEqual(a, b mathf.Rect2) bool {
	return vec2AlmostEqual(a.Position, b.Position) && vec2AlmostEqual(a.Size, b.Size)
}

func vec2iEqual(a, b mathf.Vector2i) bool {
	return a.X == b.X && a.Y == b.Y
}

func vec3iEqual(a, b mathf.Vector3i) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

func vec4iEqual(a, b mathf.Vector4i) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}

func rect2iEqual(a, b mathf.Rect2i) bool {
	return vec2iEqual(a.Position, b.Position) && vec2iEqual(a.Size, b.Size)
}
