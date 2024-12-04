package test

import (
	"testing"

	"github.com/godot-ext/mathf"

	"github.com/stretchr/testify/assert"
)

func TestRect2(t *testing.T) {
	// Test construction and basic properties
	rect := mathf.NewRect2(1, 2, 3, 4)
	expectedPos := mathf.NewVec2(1, 2)
	expectedSize := mathf.NewVec2(3, 4)

	if !vec2AlmostEqual(rect.Position, expectedPos) {
		t.Errorf("Position not set correctly: got %v, expected %v", rect.Position, expectedPos)
	}
	if !vec2AlmostEqual(rect.Size, expectedSize) {
		t.Errorf("Size not set correctly: got %v, expected %v", rect.Size, expectedSize)
	}

	// Test Area
	expectedArea := float64(12) // 3 * 4
	area := rect.Area()
	if !almostEqual(area, expectedArea) {
		t.Errorf("Area calculation failed: got %v, expected %v", area, expectedArea)
	}

	// Test HasPoint
	pointInside := mathf.NewVec2(2, 3)
	assert.True(t, rect.HasPoint(pointInside), "Point %v should be inside rect %v", pointInside, rect)

	pointOutside := mathf.NewVec2(0, 0)
	assert.False(t, rect.HasPoint(pointOutside), "Point %v should be outside rect %v", pointOutside, rect)

	// Test Intersects
	otherRect := mathf.NewRect2(2, 3, 2, 2)
	assert.True(t, rect.Intersects(otherRect, true), "Rect %v should intersect with %v", rect, otherRect)

	nonIntersectingRect := mathf.NewRect2(10, 10, 1, 1)
	assert.False(t, rect.Intersects(nonIntersectingRect, true), "Rect %v should not intersect with %v", rect, nonIntersectingRect)

	// Test Intersection
	intersection := rect.Intersection(otherRect)
	expectedIntersection := mathf.NewRect2(2, 3, 2, 2)
	if !rect2AlmostEqual(intersection, expectedIntersection) {
		t.Errorf("Intersection calculation failed: got %v, expected %v", intersection, expectedIntersection)
	}

	// Test Merge
	mergedRect := rect.Merge(otherRect)
	expectedMerged := mathf.NewRect2(1, 2, 3, 4)
	if !rect2AlmostEqual(mergedRect, expectedMerged) {
		t.Errorf("Merge failed: got %v, expected %v", mergedRect, expectedMerged)
	}

	// Test Grow
	grownRect := rect.Grow(1)
	expectedGrown := mathf.NewRect2(0, 1, 5, 6)
	if !rect2AlmostEqual(grownRect, expectedGrown) {
		t.Errorf("Grow failed: got %v, expected %v", grownRect, expectedGrown)
	}

	// Test Expand
	expandedRect := rect.Expand(pointOutside)
	expectedExpanded := mathf.NewRect2(0, 0, 4, 6)
	if !rect2AlmostEqual(expandedRect, expectedExpanded) {
		t.Errorf("Expand failed: got %v, expected %v", expandedRect, expectedExpanded)
	}
}

func TestRect2Creation(t *testing.T) {
	rect := mathf.NewRect2(10, 20, 30, 40)
	expectedPos := mathf.NewVec2(10, 20)
	expectedSize := mathf.NewVec2(30, 40)

	assert.True(t, vec2AlmostEqual(rect.Position, expectedPos), "Position should be %v, got %v", expectedPos, rect.Position)
	assert.True(t, vec2AlmostEqual(rect.Size, expectedSize), "Size should be %v, got %v", expectedSize, rect.Size)
}

func TestRect2Intersection(t *testing.T) {
	rect1 := mathf.NewRect2(0, 0, 10, 10)
	rect2 := mathf.NewRect2(5, 5, 10, 10)
	rect3 := mathf.NewRect2(20, 20, 10, 10)

	// Test intersects
	assert.True(t, rect1.Intersects(rect2, true), "rect1 should intersect with rect2")
	assert.False(t, rect1.Intersects(rect3, true), "rect1 should not intersect with rect3")

	// Test intersection result
	intersection := rect1.Intersection(rect2)
	expectedIntersection := mathf.NewRect2(5, 5, 5, 5)
	assert.True(t, rect2AlmostEqual(intersection, expectedIntersection),
		"Intersection should be %v, got %v", expectedIntersection, intersection)
}
