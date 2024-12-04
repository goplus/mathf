package test

import (
	"math"
	"testing"

	"github.com/godot-ext/mathf"
)

func TestVector2(t *testing.T) {
	// Test construction
	v1 := mathf.NewVector2(1, 2)
	v2 := mathf.NewVector2(4, 5)

	// Test basic arithmetic
	sum := v1.Add(v2)
	expected := mathf.NewVector2(5, 7)
	if !vec2AlmostEqual(sum, expected) {
		t.Errorf("Add failed: got %v, expected %v", sum, expected)
	}

	diff := v2.Sub(v1)
	expected = mathf.NewVector2(3, 3)
	if !vec2AlmostEqual(diff, expected) {
		t.Errorf("Sub failed: got %v, expected %v", diff, expected)
	}

	prod := v1.Mul(v2)
	expected = mathf.NewVector2(4, 10)
	if !vec2AlmostEqual(prod, expected) {
		t.Errorf("Mul failed: got %v, expected %v", prod, expected)
	}

	// Test scalar operations
	scaled := v1.Mulf(2)
	expected = mathf.NewVector2(2, 4)
	if !vec2AlmostEqual(scaled, expected) {
		t.Errorf("Mulf failed: got %v, expected %v", scaled, expected)
	}

	// Test dot product
	dot := v1.Dot(v2)
	expectedDot := float64(14) // 1*4 + 2*5
	if dot != expectedDot {
		t.Errorf("Dot failed: got %v, expected %v", dot, expectedDot)
	}

	// Test length
	length := v1.Length()
	expectedLen := math.Sqrt(5) // sqrt(1*1 + 2*2)
	if !almostEqual(length, expectedLen) {
		t.Errorf("Length failed: got %v, expected %v", length, expectedLen)
	}

	// Test normalization
	normalized := v1.Normalize()
	expectedNorm := mathf.NewVector2(1/math.Sqrt(5), 2/math.Sqrt(5))
	if !vec2AlmostEqual(normalized, expectedNorm) {
		t.Errorf("Normalize failed: got %v, expected %v", normalized, expectedNorm)
	}

	// Test lerp
	lerped := v1.Lerp(v2, 0.5)
	expectedLerp := mathf.NewVector2(2.5, 3.5)
	if !vec2AlmostEqual(lerped, expectedLerp) {
		t.Errorf("Lerp failed: got %v, expected %v", lerped, expectedLerp)
	}
}

func TestVector2Additional(t *testing.T) {
	v1 := mathf.NewVector2(1, 2)
	v2 := mathf.NewVector2(4, 5)

	// Test division
	div := v2.Div(v1)
	expected := mathf.NewVector2(4, 2.5)
	if !vec2AlmostEqual(div, expected) {
		t.Errorf("Div failed: got %v, expected %v", div, expected)
	}

	// Test scalar operations
	added := v1.Addf(2)
	expected = mathf.NewVector2(3, 4)
	if !vec2AlmostEqual(added, expected) {
		t.Errorf("Addf failed: got %v, expected %v", added, expected)
	}

	subbed := v1.Subf(0.5)
	expected = mathf.NewVector2(0.5, 1.5)
	if !vec2AlmostEqual(subbed, expected) {
		t.Errorf("Subf failed: got %v, expected %v", subbed, expected)
	}

	divided := v1.Divf(2)
	expected = mathf.NewVector2(0.5, 1)
	if !vec2AlmostEqual(divided, expected) {
		t.Errorf("Divf failed: got %v, expected %v", divided, expected)
	}

	// Test cross product
	cross := v1.Cross(v2)
	expectedCross := float64(-3) // 1*5 - 2*4
	if !almostEqual(cross, expectedCross) {
		t.Errorf("Cross failed: got %v, expected %v", cross, expectedCross)
	}

	// Test length squared
	lengthSq := v1.LengthSquared()
	expectedLenSq := float64(5) // 1*1 + 2*2
	if !almostEqual(lengthSq, expectedLenSq) {
		t.Errorf("LengthSquared failed: got %v, expected %v", lengthSq, expectedLenSq)
	}

	// Test distance calculations
	dist := v1.DistanceTo(v2)
	expectedDist := math.Sqrt(18) // sqrt((4-1)^2 + (5-2)^2)
	if !almostEqual(dist, expectedDist) {
		t.Errorf("DistanceTo failed: got %v, expected %v", dist, expectedDist)
	}

	distSq := v1.DistanceSquaredTo(v2)
	expectedDistSq := float64(18) // (4-1)^2 + (5-2)^2
	if !almostEqual(distSq, expectedDistSq) {
		t.Errorf("DistanceSquaredTo failed: got %v, expected %v", distSq, expectedDistSq)
	}

	// Test normalization checks
	normalized := v1.Normalize()
	if !normalized.IsNormalized() {
		t.Errorf("IsNormalized failed for normalized vector")
	}

	// Test finite check
	if !v1.IsFinite() {
		t.Errorf("IsFinite failed for finite vector")
	}

	// Test zero check
	zero := mathf.NewVector2(0.0000001, -0.0000001)
	if !zero.IsApproximatelyZero() {
		t.Errorf("IsApproximatelyZero failed for near-zero vector")
	}

	// Test absolute value
	neg := mathf.NewVector2(-1, -2)
	absV := neg.Abs()
	expected = mathf.NewVector2(1, 2)
	if !vec2AlmostEqual(absV, expected) {
		t.Errorf("Abs failed: got %v, expected %v", absV, expected)
	}

	// Test ceil/floor/round
	decimal := mathf.NewVector2(1.6, 2.2)
	ceiled := decimal.Ceil()
	expected = mathf.NewVector2(2, 3)
	if !vec2AlmostEqual(ceiled, expected) {
		t.Errorf("Ceil failed: got %v, expected %v", ceiled, expected)
	}

	floored := decimal.Floor()
	expected = mathf.NewVector2(1, 2)
	if !vec2AlmostEqual(floored, expected) {
		t.Errorf("Floor failed: got %v, expected %v", floored, expected)
	}

	rounded := decimal.Round()
	expected = mathf.NewVector2(2, 2)
	if !vec2AlmostEqual(rounded, expected) {
		t.Errorf("Round failed: got %v, expected %v", rounded, expected)
	}

	// Test sign
	mixed := mathf.NewVector2(-1.6, 2.2)
	sign := mixed.Sign()
	expected = mathf.NewVector2(-1, 1)
	if !vec2AlmostEqual(sign, expected) {
		t.Errorf("Sign failed: got %v, expected %v", sign, expected)
	}

	// Test clamp
	min := mathf.NewVector2(0, 0)
	max := mathf.NewVector2(1, 1)
	toClamp := mathf.NewVector2(-0.5, 1.5)
	clamped := toClamp.Clamp(min, max)
	expected = mathf.NewVector2(0, 1)
	if !vec2AlmostEqual(clamped, expected) {
		t.Errorf("Clamp failed: got %v, expected %v", clamped, expected)
	}

	// Test negation
	negated := v1.Neg()
	expected = mathf.NewVector2(-1, -2)
	if !vec2AlmostEqual(negated, expected) {
		t.Errorf("Neg failed: got %v, expected %v", negated, expected)
	}
}
