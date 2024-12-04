package test

import (
	"godot-ext/mathf"
	"testing"
)

func TestVector4i(t *testing.T) {
	v1 := mathf.NewVector4i(1, 2, 3, 4)
	v2 := mathf.NewVector4i(5, 6, 7, 8)

	// Test basic arithmetic
	sum := v1.Add(v2)
	expected := mathf.NewVector4i(6, 8, 10, 12)
	if !vec4iEqual(sum, expected) {
		t.Errorf("Add failed: got %v, expected %v", sum, expected)
	}

	diff := v2.Sub(v1)
	expected = mathf.NewVector4i(4, 4, 4, 4)
	if !vec4iEqual(diff, expected) {
		t.Errorf("Sub failed: got %v, expected %v", diff, expected)
	}

	prod := v1.Mul(v2)
	expected = mathf.NewVector4i(5, 12, 21, 32)
	if !vec4iEqual(prod, expected) {
		t.Errorf("Mul failed: got %v, expected %v", prod, expected)
	}

	div := v2.Div(v1)
	expected = mathf.NewVector4i(5, 3, 2, 2)
	if !vec4iEqual(div, expected) {
		t.Errorf("Div failed: got %v, expected %v", div, expected)
	}

	// Test scalar operations
	added := v1.Addi(2)
	expected = mathf.NewVector4i(3, 4, 5, 6)
	if !vec4iEqual(added, expected) {
		t.Errorf("Addi failed: got %v, expected %v", added, expected)
	}

	subbed := v1.Subi(1)
	expected = mathf.NewVector4i(0, 1, 2, 3)
	if !vec4iEqual(subbed, expected) {
		t.Errorf("Subi failed: got %v, expected %v", subbed, expected)
	}

	multed := v1.Muli(2)
	expected = mathf.NewVector4i(2, 4, 6, 8)
	if !vec4iEqual(multed, expected) {
		t.Errorf("Muli failed: got %v, expected %v", multed, expected)
	}

	divided := v2.Divi(2)
	expected = mathf.NewVector4i(2, 3, 3, 4)
	if !vec4iEqual(divided, expected) {
		t.Errorf("Divi failed: got %v, expected %v", divided, expected)
	}

	// Test dot product
	dot := v1.Dot(v2)
	expectedDot := float64(70) // 1*5 + 2*6 + 3*7 + 4*8
	if !almostEqual(dot, expectedDot) {
		t.Errorf("Dot failed: got %v, expected %v", dot, expectedDot)
	}

	// Test length calculations
	lengthSq := v1.LengthSquared()
	expectedLenSq := 30.0 // 1*1 + 2*2 + 3*3 + 4*4
	if lengthSq != expectedLenSq {
		t.Errorf("LengthSquared failed: got %v, expected %v", lengthSq, expectedLenSq)
	}

	// Test distance calculations
	distSq := v1.DistanceSquaredTo(v2)
	expectedDistSq := 64.0 // (5-1)^2 + (6-2)^2 + (7-3)^2 + (8-4)^2
	if distSq != expectedDistSq {
		t.Errorf("DistanceSquaredTo failed: got %v, expected %v", distSq, expectedDistSq)
	}

	// Test absolute value
	neg := mathf.NewVector4i(-1, -2, -3, -4)
	absV := neg.Abs()
	expected = mathf.NewVector4i(1, 2, 3, 4)
	if !vec4iEqual(absV, expected) {
		t.Errorf("Abs failed: got %v, expected %v", absV, expected)
	}

	// Test clamp
	min := mathf.NewVector4i(0, 0, 0, 0)
	max := mathf.NewVector4i(1, 1, 1, 1)
	toClamp := mathf.NewVector4i(-1, 2, 0, 3)
	clamped := toClamp.Clamp(min, max)
	expected = mathf.NewVector4i(0, 1, 0, 1)
	if !vec4iEqual(clamped, expected) {
		t.Errorf("Clamp failed: got %v, expected %v", clamped, expected)
	}

	// Test negation
	negated := v1.Neg()
	expected = mathf.NewVector4i(-1, -2, -3, -4)
	if !vec4iEqual(negated, expected) {
		t.Errorf("Neg failed: got %v, expected %v", negated, expected)
	}
}
