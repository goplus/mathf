# Mathf - Mathematical Utilities for Go

A comprehensive mathematical utility library for Go that provides efficient implementations of common geometric types and operations. This library is designed to be performant, easy to use, and particularly well-suited for game development and graphics applications.

## Features

- **Vector Types**
  - `Vector2`, `Vector2i` - 2D vectors (float and integer versions)
  - `Vector3`, `Vector3i` - 3D vectors (float and integer versions)
  - `Vector4`, `Vector4i` - 4D vectors (float and integer versions)

- **Rectangle Types**
  - `Rect2`, `Rect2i` - 2D rectangles (float and integer versions)
  - `AABB` - Axis-Aligned Bounding Box for 3D collision detection

- **Quaternion**
  - Complete quaternion implementation for 3D rotations

## Installation

```bash
go get github.com/godot-ext/mathf
```

## Usage

### Vectors

```go
import "github.com/godot-ext/mathf"

// Create vectors
v1 := mathf.NewVector2(1.0, 2.0)
v2 := mathf.NewVector2i(1, 2)

// Basic operations
sum := v1.Add(v2)
diff := v1.Sub(v2)
scaled := v1.Mul(2.0)

// Vector operations
dot := v1.Dot(v2)
length := v1.Length()
normalized := v1.Normalized()
```

### Rectangles

```go
import "github.com/godot-ext/mathf"

// Create rectangles (x, y, width, height)
rect := mathf.NewRect2(0, 0, 100, 100)
rect2 := mathf.NewRect2i(0, 0, 100, 100)

// Check if a point is inside
point := mathf.NewVector2(50, 50)
isInside := rect.HasPoint(point)

// Rectangle operations
intersection := rect.Intersection(rect2)  // Get intersection area
merged := rect.Merge(rect2)              // Combine rectangles
grown := rect.Grow(10)                   // Expand by 10 units
```

### Quaternions

```go
import "github.com/godot-ext/mathf"

// Create a quaternion
q := mathf.NewQuaternion(x, y, z, w)

// Common operations
rotated := q.Rotate(vector)    // Rotate a vector
inverse := q.Inverse()         // Get inverse rotation
```

## Testing

The library includes comprehensive test coverage for all components. Run the tests using:

```bash
go test ./...
```

## Contributing

We welcome contributions! Feel free to:
- Submit bug reports and feature requests
- Propose improvements to documentation
- Create pull requests for bug fixes or new features

Please ensure your code follows the existing style and includes appropriate tests.

## License

Apache License 2.0

## Acknowledgments

This library draws inspiration from:
- [Godot Engine's](https://godotengine.org/) mathematical utilities
- [xy](https://github.com/grow-graphics/xy) - An awesome graphics library

Both have significantly influenced our design and implementation choices.