package test

import (
	"testing"

	"github.com/realdream-ai/mathf"
)

func TestVec3i(t *testing.T) {
	v1 := mathf.NewVec3i(1, 2, 3)
	v2 := mathf.NewVec3i(4, 5, 6)

	// Test basic arithmetic
	sum := v1.Add(v2)
	expected := mathf.NewVec3i(5, 7, 9)
	if !vec3iEqual(sum, expected) {
		t.Errorf("Add failed: got %v, expected %v", sum, expected)
	}

	diff := v2.Sub(v1)
	expected = mathf.NewVec3i(3, 3, 3)
	if !vec3iEqual(diff, expected) {
		t.Errorf("Sub failed: got %v, expected %v", diff, expected)
	}

	prod := v1.Mul(v2)
	expected = mathf.NewVec3i(4, 10, 18)
	if !vec3iEqual(prod, expected) {
		t.Errorf("Mul failed: got %v, expected %v", prod, expected)
	}

	div := v2.Div(v1)
	expected = mathf.NewVec3i(4, 2, 2)
	if !vec3iEqual(div, expected) {
		t.Errorf("Div failed: got %v, expected %v", div, expected)
	}

	// Test scalar operations
	added := v1.Addi(2)
	expected = mathf.NewVec3i(3, 4, 5)
	if !vec3iEqual(added, expected) {
		t.Errorf("Addi failed: got %v, expected %v", added, expected)
	}

	subbed := v1.Subi(1)
	expected = mathf.NewVec3i(0, 1, 2)
	if !vec3iEqual(subbed, expected) {
		t.Errorf("Subi failed: got %v, expected %v", subbed, expected)
	}

	multed := v1.Muli(2)
	expected = mathf.NewVec3i(2, 4, 6)
	if !vec3iEqual(multed, expected) {
		t.Errorf("Muli failed: got %v, expected %v", multed, expected)
	}

	divided := v2.Divi(2)
	expected = mathf.NewVec3i(2, 2, 3)
	if !vec3iEqual(divided, expected) {
		t.Errorf("Divi failed: got %v, expected %v", divided, expected)
	}

	// Test dot product
	dot := v1.Dot(v2)
	expectedDot := float64(32) // 1*4 + 2*5 + 3*6
	if !almostEqual(dot, expectedDot) {
		t.Errorf("Dot failed: got %v, expected %v", dot, expectedDot)
	}

	// Test length calculations
	lengthSq := v1.LengthSquared()
	expectedLenSq := 14.0 // 1*1 + 2*2 + 3*3
	if lengthSq != expectedLenSq {
		t.Errorf("LengthSquared failed: got %v, expected %v", lengthSq, expectedLenSq)
	}

	// Test distance calculations
	distSq := v1.DistanceSquaredTo(v2)
	expectedDistSq := 27.0 // (4-1)^2 + (5-2)^2 + (6-3)^2
	if distSq != expectedDistSq {
		t.Errorf("DistanceSquaredTo failed: got %v, expected %v", distSq, expectedDistSq)
	}

	// Test absolute value
	neg := mathf.NewVec3i(-1, -2, -3)
	absV := neg.Abs()
	expected = mathf.NewVec3i(1, 2, 3)
	if !vec3iEqual(absV, expected) {
		t.Errorf("Abs failed: got %v, expected %v", absV, expected)
	}

	// Test clamp
	min := mathf.NewVec3i(0, 0, 0)
	max := mathf.NewVec3i(1, 1, 1)
	toClamp := mathf.NewVec3i(-1, 2, 0)
	clamped := toClamp.Clamp(min, max)
	expected = mathf.NewVec3i(0, 1, 0)
	if !vec3iEqual(clamped, expected) {
		t.Errorf("Clamp failed: got %v, expected %v", clamped, expected)
	}

	// Test negation
	negated := v1.Neg()
	expected = mathf.NewVec3i(-1, -2, -3)
	if !vec3iEqual(negated, expected) {
		t.Errorf("Neg failed: got %v, expected %v", negated, expected)
	}
}
