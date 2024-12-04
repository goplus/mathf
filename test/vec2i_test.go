package test

import (
	"testing"

	"github.com/realdream-ai/mathf"

	"github.com/stretchr/testify/assert"
)

func TestVec2i(t *testing.T) {
	// Test construction
	v1 := mathf.NewVec2i(1, 2)
	v2 := mathf.NewVec2i(3, 4)

	// Test addition
	sum := v1.Add(v2)
	assert.Equal(t, mathf.NewVec2i(4, 6), sum, "Addition failed")

	// Test subtraction
	diff := v2.Sub(v1)
	assert.Equal(t, mathf.NewVec2i(2, 2), diff, "Subtraction failed")

	// Test multiplication
	prod := v1.Mul(v2)
	assert.Equal(t, mathf.NewVec2i(3, 8), prod, "Multiplication failed")

	// Test division
	quot := v2.Div(v1)
	assert.Equal(t, mathf.NewVec2i(3, 2), quot, "Division failed")

	// Test scalar operations
	scalar := 2
	scalarAdd := v1.Addi(scalar)
	assert.Equal(t, mathf.NewVec2i(3, 4), scalarAdd, "Scalar addition failed")

	scalarSub := v1.Subi(scalar)
	assert.Equal(t, mathf.NewVec2i(-1, 0), scalarSub, "Scalar subtraction failed")

	scalarMul := v1.Muli(scalar)
	assert.Equal(t, mathf.NewVec2i(2, 4), scalarMul, "Scalar multiplication failed")

	scalarDiv := v2.Divi(scalar)
	assert.Equal(t, mathf.NewVec2i(1, 2), scalarDiv, "Scalar division failed")

	// Test dot product
	dot := v1.Dot(v2)
	assert.InDelta(t, float64(11), dot, 0.0001, "Dot product failed")

	// Test length calculations
	length := v1.Length()
	assert.InDelta(t, 2.236068, length, 0.0001, "Length calculation failed")

	lengthSquared := v1.LengthSquared()
	assert.Equal(t, 5.0, lengthSquared, "Length squared calculation failed")

	// Test distance calculations
	distance := v1.DistanceTo(v2)
	assert.InDelta(t, 2.828427, distance, 0.0001, "Distance calculation failed")

	distanceSquared := v1.DistanceSquaredTo(v2)
	assert.Equal(t, 8.0, distanceSquared, "Distance squared calculation failed")

	// Test abs
	negVec := mathf.NewVec2i(-1, -2)
	absVec := negVec.Abs()
	assert.Equal(t, v1, absVec, "Abs failed")

	// Test sign
	signVec := negVec.Sign()
	assert.Equal(t, mathf.NewVec2i(-1, -1), signVec, "Sign failed")

	// Test clamp
	min := mathf.NewVec2i(0, 0)
	max := mathf.NewVec2i(5, 5)
	clamped := v1.Clamp(min, max)
	assert.Equal(t, v1, clamped, "Clamp failed")

	// Test negation
	neg := v1.Neg()
	assert.Equal(t, mathf.NewVec2i(-1, -2), neg, "Negation failed")
}
