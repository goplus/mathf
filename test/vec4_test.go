package test

import (
	"math"
	"testing"

	"github.com/goplus/mathf"
)

func TestVec4(t *testing.T) {
	v1 := mathf.NewVec4(1, 2, 3, 4)
	v2 := mathf.NewVec4(5, 6, 7, 8)

	// Test basic arithmetic
	sum := v1.Add(v2)
	expected := mathf.NewVec4(6, 8, 10, 12)
	if !vec4AlmostEqual(sum, expected) {
		t.Errorf("Add failed: got %v, expected %v", sum, expected)
	}

	diff := v2.Sub(v1)
	expected = mathf.NewVec4(4, 4, 4, 4)
	if !vec4AlmostEqual(diff, expected) {
		t.Errorf("Sub failed: got %v, expected %v", diff, expected)
	}

	prod := v1.Mul(v2)
	expected = mathf.NewVec4(5, 12, 21, 32)
	if !vec4AlmostEqual(prod, expected) {
		t.Errorf("Mul failed: got %v, expected %v", prod, expected)
	}

	div := v2.Div(v1)
	expected = mathf.NewVec4(5, 3, 7.0/3.0, 2)
	if !vec4AlmostEqual(div, expected) {
		t.Errorf("Div failed: got %v, expected %v", div, expected)
	}

	// Test scalar operations
	added := v1.Addf(2)
	expected = mathf.NewVec4(3, 4, 5, 6)
	if !vec4AlmostEqual(added, expected) {
		t.Errorf("Addf failed: got %v, expected %v", added, expected)
	}

	subbed := v1.Subf(0.5)
	expected = mathf.NewVec4(0.5, 1.5, 2.5, 3.5)
	if !vec4AlmostEqual(subbed, expected) {
		t.Errorf("Subf failed: got %v, expected %v", subbed, expected)
	}

	multed := v1.Mulf(2)
	expected = mathf.NewVec4(2, 4, 6, 8)
	if !vec4AlmostEqual(multed, expected) {
		t.Errorf("Mulf failed: got %v, expected %v", multed, expected)
	}

	divided := v2.Divf(2)
	expected = mathf.NewVec4(2.5, 3, 3.5, 4)
	if !vec4AlmostEqual(divided, expected) {
		t.Errorf("Divf failed: got %v, expected %v", divided, expected)
	}

	// Test dot product
	dot := v1.Dot(v2)
	expectedDot := float64(70) // 1*5 + 2*6 + 3*7 + 4*8
	if !almostEqual(dot, expectedDot) {
		t.Errorf("Dot failed: got %v, expected %v", dot, expectedDot)
	}

	// Test length calculations
	lengthSq := v1.LengthSquared()
	expectedLenSq := float64(30) // 1*1 + 2*2 + 3*3 + 4*4
	if !almostEqual(lengthSq, expectedLenSq) {
		t.Errorf("LengthSquared failed: got %v, expected %v", lengthSq, expectedLenSq)
	}

	length := v1.Length()
	expectedLen := math.Sqrt(30)
	if !almostEqual(length, expectedLen) {
		t.Errorf("Length failed: got %v, expected %v", length, expectedLen)
	}

	// Test distance calculations
	dist := v1.DistanceTo(v2)
	expectedDist := math.Sqrt(64) // sqrt((5-1)^2 + (6-2)^2 + (7-3)^2 + (8-4)^2)
	if !almostEqual(dist, expectedDist) {
		t.Errorf("DistanceTo failed: got %v, expected %v", dist, expectedDist)
	}

	distSq := v1.DistanceSquaredTo(v2)
	expectedDistSq := float64(64) // (5-1)^2 + (6-2)^2 + (7-3)^2 + (8-4)^2
	if !almostEqual(distSq, expectedDistSq) {
		t.Errorf("DistanceSquaredTo failed: got %v, expected %v", distSq, expectedDistSq)
	}

	// Test normalization
	normalized := v1.Normalize()
	if !normalized.IsNormalized() {
		t.Errorf("Normalize failed: result is not normalized")
	}

	// Test finite check
	if !v1.IsFinite() {
		t.Errorf("IsFinite failed for finite Vec")
	}

	// Test zero check
	zero := mathf.NewVec4(0.0000001, -0.0000001, 0.0000001, -0.0000001)
	if !zero.IsApproximatelyZero() {
		t.Errorf("IsApproximatelyZero failed for near-zero Vec")
	}

	// Test absolute value
	neg := mathf.NewVec4(-1, -2, -3, -4)
	absV := neg.Abs()
	expected = mathf.NewVec4(1, 2, 3, 4)
	if !vec4AlmostEqual(absV, expected) {
		t.Errorf("Abs failed: got %v, expected %v", absV, expected)
	}

	// Test ceil/floor/round
	decimal := mathf.NewVec4(1.6, 2.2, -2.7, 3.5)
	ceiled := decimal.Ceil()
	expected = mathf.NewVec4(2, 3, -2, 4)
	if !vec4AlmostEqual(ceiled, expected) {
		t.Errorf("Ceil failed: got %v, expected %v", ceiled, expected)
	}

	floored := decimal.Floor()
	expected = mathf.NewVec4(1, 2, -3, 3)
	if !vec4AlmostEqual(floored, expected) {
		t.Errorf("Floor failed: got %v, expected %v", floored, expected)
	}

	rounded := decimal.Round()
	expected = mathf.NewVec4(2, 2, -3, 4)
	if !vec4AlmostEqual(rounded, expected) {
		t.Errorf("Round failed: got %v, expected %v", rounded, expected)
	}

	// Test sign
	sign := neg.Sign()
	expected = mathf.NewVec4(-1, -1, -1, -1)
	if !vec4AlmostEqual(sign, expected) {
		t.Errorf("Sign failed: got %v, expected %v", sign, expected)
	}

	// Test clamp
	min := mathf.NewVec4(0, 0, 0, 0)
	max := mathf.NewVec4(1, 1, 1, 1)
	toClamp := mathf.NewVec4(-0.5, 1.5, 0.5, 2)
	clamped := toClamp.Clamp(min, max)
	expected = mathf.NewVec4(0, 1, 0.5, 1)
	if !vec4AlmostEqual(clamped, expected) {
		t.Errorf("Clamp failed: got %v, expected %v", clamped, expected)
	}

	// Test negation
	negated := v1.Neg()
	expected = mathf.NewVec4(-1, -2, -3, -4)
	if !vec4AlmostEqual(negated, expected) {
		t.Errorf("Neg failed: got %v, expected %v", negated, expected)
	}

	// Test lerp
	lerped := v1.Lerp(v2, 0.5)
	expected = mathf.NewVec4(3, 4, 5, 6)
	if !vec4AlmostEqual(lerped, expected) {
		t.Errorf("Lerp failed: got %v, expected %v", lerped, expected)
	}
}
