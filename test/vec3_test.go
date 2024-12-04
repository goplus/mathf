package test

import (
	"testing"

	"github.com/godot-ext/mathf"
)

func TestVec3(t *testing.T) {
	v1 := mathf.NewVec3(1, 2, 3)
	v2 := mathf.NewVec3(4, 5, 6)

	// Test basic arithmetic
	sum := v1.Add(v2)
	expected := mathf.NewVec3(5, 7, 9)
	if !vec3AlmostEqual(sum, expected) {
		t.Errorf("Add failed: got %v, expected %v", sum, expected)
	}

	diff := v2.Sub(v1)
	expected = mathf.NewVec3(3, 3, 3)
	if !vec3AlmostEqual(diff, expected) {
		t.Errorf("Sub failed: got %v, expected %v", diff, expected)
	}

	prod := v1.Mul(v2)
	expected = mathf.NewVec3(4, 10, 18)
	if !vec3AlmostEqual(prod, expected) {
		t.Errorf("Mul failed: got %v, expected %v", prod, expected)
	}

	div := v2.Div(v1)
	expected = mathf.NewVec3(4, 2.5, 2)
	if !vec3AlmostEqual(div, expected) {
		t.Errorf("Div failed: got %v, expected %v", div, expected)
	}

	// Test scalar operations
	added := v1.Addf(2)
	expected = mathf.NewVec3(3, 4, 5)
	if !vec3AlmostEqual(added, expected) {
		t.Errorf("Addf failed: got %v, expected %v", added, expected)
	}

	subbed := v1.Subf(1)
	expected = mathf.NewVec3(0, 1, 2)
	if !vec3AlmostEqual(subbed, expected) {
		t.Errorf("Subf failed: got %v, expected %v", subbed, expected)
	}

	multed := v1.Mulf(2)
	expected = mathf.NewVec3(2, 4, 6)
	if !vec3AlmostEqual(multed, expected) {
		t.Errorf("Mulf failed: got %v, expected %v", multed, expected)
	}

	divided := v2.Divf(2)
	expected = mathf.NewVec3(2, 2.5, 3)
	if !vec3AlmostEqual(divided, expected) {
		t.Errorf("Divf failed: got %v, expected %v", divided, expected)
	}

	// Test cross product
	cross := v1.Cross(v2)
	expected = mathf.NewVec3(-3, 6, -3)
	if !vec3AlmostEqual(cross, expected) {
		t.Errorf("Cross failed: got %v, expected %v", cross, expected)
	}

	// Test dot product
	dot := v1.Dot(v2)
	expectedDot := float64(32) // 1*4 + 2*5 + 3*6
	if !almostEqual(dot, expectedDot) {
		t.Errorf("Dot failed: got %v, expected %v", dot, expectedDot)
	}

	// Test length calculations
	lengthSq := v1.LengthSquared()
	expectedLenSq := float64(14) // 1*1 + 2*2 + 3*3
	if !almostEqual(lengthSq, expectedLenSq) {
		t.Errorf("LengthSquared failed: got %v, expected %v", lengthSq, expectedLenSq)
	}

	// Test distance calculations
	distSq := v1.DistanceSquaredTo(v2)
	expectedDistSq := float64(27) // (4-1)^2 + (5-2)^2 + (6-3)^2
	if !almostEqual(distSq, expectedDistSq) {
		t.Errorf("DistanceSquaredTo failed: got %v, expected %v", distSq, expectedDistSq)
	}

	// Test absolute value
	neg := mathf.NewVec3(-1, -2, -3)
	absV := neg.Abs()
	expected = mathf.NewVec3(1, 2, 3)
	if !vec3AlmostEqual(absV, expected) {
		t.Errorf("Abs failed: got %v, expected %v", absV, expected)
	}

	// Test clamp
	min := mathf.NewVec3(0, 0, 0)
	max := mathf.NewVec3(1, 1, 1)
	toClamp := mathf.NewVec3(-1, 2, 0.5)
	clamped := toClamp.Clamp(min, max)
	expected = mathf.NewVec3(0, 1, 0.5)
	if !vec3AlmostEqual(clamped, expected) {
		t.Errorf("Clamp failed: got %v, expected %v", clamped, expected)
	}

	// Test negation
	negated := v1.Neg()
	expected = mathf.NewVec3(-1, -2, -3)
	if !vec3AlmostEqual(negated, expected) {
		t.Errorf("Neg failed: got %v, expected %v", negated, expected)
	}
}
