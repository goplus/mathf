package test

import (
	"testing"

	"github.com/godot-ext/mathf"
	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	// Test basic color creation
	c1 := mathf.Color{R: 1.0, G: 0.5, B: 0.25, A: 1.0}
	assert.Equal(t, float32(1.0), c1.R, "Red component mismatch")
	assert.Equal(t, float32(0.5), c1.G, "Green component mismatch")
	assert.Equal(t, float32(0.25), c1.B, "Blue component mismatch")
	assert.Equal(t, float32(1.0), c1.A, "Alpha component mismatch")

	// Test string representation
	expectedStr := "(1.000000, 0.500000, 0.250000, 1.000000)"
	assert.Equal(t, expectedStr, c1.String(), "String representation mismatch")

	// Test color addition
	c2 := mathf.Color{R: 0.5, G: 0.25, B: 0.125, A: 0.5}
	sum := c1.Add(c2)
	assert.InDelta(t, float32(1.5), sum.R, 0.0001, "Color addition: Red component mismatch")
	assert.InDelta(t, float32(0.75), sum.G, 0.0001, "Color addition: Green component mismatch")
	assert.InDelta(t, float32(0.375), sum.B, 0.0001, "Color addition: Blue component mismatch")
	assert.InDelta(t, float32(1.5), sum.A, 0.0001, "Color addition: Alpha component mismatch")

	// Test color subtraction
	diff := c1.Sub(c2)
	assert.InDelta(t, float32(0.5), diff.R, 0.0001, "Color subtraction: Red component mismatch")
	assert.InDelta(t, float32(0.25), diff.G, 0.0001, "Color subtraction: Green component mismatch")
	assert.InDelta(t, float32(0.125), diff.B, 0.0001, "Color subtraction: Blue component mismatch")
	assert.InDelta(t, float32(0.5), diff.A, 0.0001, "Color subtraction: Alpha component mismatch")

	// Test color multiplication
	prod := c1.Mul(c2)
	assert.InDelta(t, float32(0.5), prod.R, 0.0001, "Color multiplication: Red component mismatch")
	assert.InDelta(t, float32(0.125), prod.G, 0.0001, "Color multiplication: Green component mismatch")
	assert.InDelta(t, float32(0.03125), prod.B, 0.0001, "Color multiplication: Blue component mismatch")
	assert.InDelta(t, float32(0.5), prod.A, 0.0001, "Color multiplication: Alpha component mismatch")

	// Test scalar multiplication
	scalar := float32(2.0)
	scalarProd := c1.Mulf(scalar)
	assert.InDelta(t, float32(2.0), scalarProd.R, 0.0001, "Scalar multiplication: Red component mismatch")
	assert.InDelta(t, float32(1.0), scalarProd.G, 0.0001, "Scalar multiplication: Green component mismatch")
	assert.InDelta(t, float32(0.5), scalarProd.B, 0.0001, "Scalar multiplication: Blue component mismatch")
	assert.InDelta(t, float32(2.0), scalarProd.A, 0.0001, "Scalar multiplication: Alpha component mismatch")

	// Test color clamping
	unclamped := mathf.Color{R: 1.5, G: -0.5, B: 2.0, A: -0.1}
	clamped := unclamped.Clamp()
	assert.Equal(t, float32(1.0), clamped.R, "Clamping: Red component mismatch")
	assert.Equal(t, float32(0.0), clamped.G, "Clamping: Green component mismatch")
	assert.Equal(t, float32(1.0), clamped.B, "Clamping: Blue component mismatch")
	assert.Equal(t, float32(0.0), clamped.A, "Clamping: Alpha component mismatch")

	// Test color inversion
	inverted := c1.Invert()
	assert.InDelta(t, float32(0.0), inverted.R, 0.0001, "Inversion: Red component mismatch")
	assert.InDelta(t, float32(0.5), inverted.G, 0.0001, "Inversion: Green component mismatch")
	assert.InDelta(t, float32(0.75), inverted.B, 0.0001, "Inversion: Blue component mismatch")
	assert.InDelta(t, float32(1.0), inverted.A, 0.0001, "Inversion: Alpha component mismatch")

	// Test color lerp
	lerped := c1.Lerp(c2, 0.5)
	assert.InDelta(t, float32(0.75), lerped.R, 0.0001, "Lerp: Red component mismatch")
	assert.InDelta(t, float32(0.375), lerped.G, 0.0001, "Lerp: Green component mismatch")
	assert.InDelta(t, float32(0.1875), lerped.B, 0.0001, "Lerp: Blue component mismatch")
	assert.InDelta(t, float32(0.75), lerped.A, 0.0001, "Lerp: Alpha component mismatch")
}
