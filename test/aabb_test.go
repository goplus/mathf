package test

import (
	"godot-ext/mathf"
	"testing"
)

func TestAABB(t *testing.T) {
	// Test construction and basic properties
	pos := mathf.NewVector3(1, 2, 3)
	size := mathf.NewVector3(4, 5, 6)
	aabb := mathf.AABB{Position: pos, Size: size}

	if !vec3AlmostEqual(aabb.Position, pos) {
		t.Errorf("Position not set correctly: got %v, expected %v", aabb.Position, pos)
	}
	if !vec3AlmostEqual(aabb.Size, size) {
		t.Errorf("Size not set correctly: got %v, expected %v", aabb.Size, size)
	}
}
