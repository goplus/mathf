package test

import (
	"godot-ext/mathf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRect2i(t *testing.T) {
	rect := mathf.NewRect2i(1, 2, 3, 4)
	expectedPos := mathf.NewVector2i(1, 2)
	expectedSize := mathf.NewVector2i(3, 4)

	// Test basic properties
	assert.True(t, vec2iEqual(rect.Position, expectedPos), "Position should be %v, got %v", expectedPos, rect.Position)
	assert.True(t, vec2iEqual(rect.Size, expectedSize), "Size should be %v, got %v", expectedSize, rect.Size)

	// Test Area
	expectedArea := int64(12)
	assert.Equal(t, expectedArea, rect.Area(), "Area should be %v, got %v", expectedArea, rect.Area())

	// Test HasArea
	assert.True(t, rect.HasArea(), "HasArea should be true")

	// Test Center
	center := rect.Center()
	expectedCenter := mathf.NewVector2i(2, 4)
	assert.True(t, vec2iEqual(center, expectedCenter), "Center should be %v, got %v", expectedCenter, center)

	// Test End
	end := rect.End()
	expectedEnd := mathf.NewVector2i(4, 6)
	assert.True(t, vec2iEqual(end, expectedEnd), "End should be %v, got %v", expectedEnd, end)

	// Test HasPoint
	pointInside := mathf.NewVector2i(2, 3)
	assert.True(t, rect.HasPoint(pointInside), "Point %v should be inside rect %v", pointInside, rect)

	pointOutside := mathf.NewVector2i(0, 0)
	assert.False(t, rect.HasPoint(pointOutside), "Point %v should be outside rect %v", pointOutside, rect)

	// Test Intersects
	otherRect := mathf.NewRect2i(2, 3, 2, 2)
	assert.True(t, rect.Intersects(otherRect, true), "Rect %v should intersect with %v", rect, otherRect)

	nonIntersectingRect := mathf.NewRect2i(10, 10, 1, 1)
	assert.False(t, rect.Intersects(nonIntersectingRect, true), "Rect %v should not intersect with %v", rect, nonIntersectingRect)

	// Test Intersection
	intersection := rect.Intersection(otherRect)
	expectedIntersection := mathf.NewRect2i(2, 3, 2, 2)
	assert.True(t, rect2iEqual(intersection, expectedIntersection),
		"Intersection should be %v, got %v", expectedIntersection, intersection)

	// Test Merge
	merged := rect.Merge(otherRect)
	expectedMerged := mathf.NewRect2i(1, 2, 3, 4)
	assert.True(t, rect2iEqual(merged, expectedMerged),
		"Merge should be %v, got %v", expectedMerged, merged)

	// Test Grow
	grown := rect.Grow(1)
	expectedGrown := mathf.NewRect2i(0, 1, 5, 6)
	assert.True(t, rect2iEqual(grown, expectedGrown),
		"Grow should be %v, got %v", expectedGrown, grown)

	// Test GrowIndividual
	grownIndiv := rect.GrowIndividual(1, 2, 3, 4)
	expectedGrownIndiv := mathf.NewRect2i(0, 0, 7, 10)
	assert.True(t, rect2iEqual(grownIndiv, expectedGrownIndiv),
		"GrowIndividual should be %v, got %v", expectedGrownIndiv, grownIndiv)

	// Test Expand
	expanded := rect.Expand(mathf.NewVector2i(5, 6))
	expectedExpanded := mathf.NewRect2i(1, 2, 4, 4)
	assert.True(t, rect2iEqual(expanded, expectedExpanded),
		"Expand should be %v, got %v", expectedExpanded, expanded)

	// Test Abs
	negRect := mathf.NewRect2i(-1, -2, -3, -4)
	absRect := negRect.Abs()
	expectedAbs := mathf.NewRect2i(1, 2, 3, 4)
	assert.True(t, rect2iEqual(absRect, expectedAbs),
		"Abs should be %v, got %v", expectedAbs, absRect)
}
