# Mathf - Mathematical Utilities for Go

A comprehensive mathematical utility library for Go, providing implementations of common geometric types and operations. This library is designed to be efficient, easy to use, and suitable for game development and graphics applications.

## Features

- **Vector Types**
  - `Vector2`, `Vector2i` - 2D vectors (float and integer)
  - `Vector3`, `Vector3i` - 3D vectors (float and integer)
  - `Vector4`, `Vector4i` - 4D vectors (float and integer)

- **Rectangle Types**
  - `Rect2`, `Rect2i` - 2D rectangles (float and integer)
  - `AABB` - Axis-Aligned Bounding Box

- **Quaternion**
  - Full quaternion implementation for 3D rotations

## Installation

```bash
go get github.com/realdream-ai/mathf
```

## Usage

### Vectors

```go
import "github.com/realdream-ai/mathf"

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
import "github.com/realdream-ai/mathf"

// Create rectangles
rect := mathf.NewRect2(0, 0, 100, 100)  // x, y, width, height
rect2 := mathf.NewRect2i(0, 0, 100, 100)

// Check if point is inside
point := mathf.NewVector2(50, 50)
isInside := rect.HasPoint(point)

// Rectangle operations
intersection := rect.Intersection(rect2)
merged := rect.Merge(rect2)
grown := rect.Grow(10)
```

### Quaternions

```go
import "github.com/realdream-ai/mathf"

// Create quaternion
q := mathf.NewQuaternion(x, y, z, w)

// Operations
rotated := q.Rotate(vector)
inverse := q.Inverse()
```

## Testing

The library includes comprehensive test coverage. Run the tests using:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Apache License 2.0

## Acknowledgments

This library is inspired by Godot Engine's math utilities and the awesome library: [xy](https://github.com/grow-graphics/xy), and designed to provide similar functionality in Go.  