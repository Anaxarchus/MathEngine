package primitive

import "math"

type Vector2 struct {
	X float64
	Y float64
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

type Vector4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

type Normal2 Vector2
type Normal3 Vector3
type Normal4 Vector4

type Euler float64
type Degree float64
type Radian float64

// addition
func (v Vector2) Add(rh Vector2) Vector2 {
	return Vector2{v.X + rh.X, v.Y + rh.Y}
}

func (v Vector3) Add(rh Vector3) Vector3 {
	return Vector3{v.X + rh.X, v.Y + rh.Y, v.Z + rh.Z}
}

func (v Vector4) Add(rh Vector4) Vector4 {
	return Vector4{v.X + rh.X, v.Y + rh.Y, v.Z + rh.Z, v.W + rh.W}
}

func (v Vector2) Addf(rh float64) Vector2 {
	return Vector2{v.X + rh, v.Y + rh}
}

func (v Vector3) Addf(rh float64) Vector3 {
	return Vector3{v.X + rh, v.Y + rh, v.Z + rh}
}

func (v Vector4) Addf(rh float64) Vector4 {
	return Vector4{v.X + rh, v.Y + rh, v.Z + rh, v.W + rh}
}

// subtraction
func (v Vector2) Sub(rh Vector2) Vector2 {
	return Vector2{v.X - rh.X, v.Y - rh.Y}
}

func (v Vector3) Sub(rh Vector3) Vector3 {
	return Vector3{v.X - rh.X, v.Y - rh.Y, v.Z - rh.Z}
}

func (v Vector4) Sub(rh Vector4) Vector4 {
	return Vector4{v.X - rh.X, v.Y - rh.Y, v.Z - rh.Z, v.W - rh.W}
}

func (v Vector2) Subf(rh float64) Vector2 {
	return Vector2{v.X - rh, v.Y - rh}
}

func (v Vector3) Subf(rh float64) Vector3 {
	return Vector3{v.X - rh, v.Y - rh, v.Z - rh}
}

func (v Vector4) Subf(rh float64) Vector4 {
	return Vector4{v.X - rh, v.Y - rh, v.Z - rh, v.W - rh}
}

// multiplication
func (v Vector2) Mul(rh Vector2) Vector2 {
	return Vector2{v.X * rh.X, v.Y * rh.Y}
}

func (v Vector3) Mul(rh Vector3) Vector3 {
	return Vector3{v.X * rh.X, v.Y * rh.Y, v.Z * rh.Z}
}

func (v Vector4) Mul(rh Vector4) Vector4 {
	return Vector4{v.X * rh.X, v.Y * rh.Y, v.Z * rh.Z, v.W * rh.W}
}

func (v Vector2) Mulf(rh float64) Vector2 {
	return Vector2{v.X * rh, v.Y * rh}
}

func (v Vector3) Mulf(rh float64) Vector3 {
	return Vector3{v.X * rh, v.Y * rh, v.Z * rh}
}

func (v Vector4) Mulf(rh float64) Vector4 {
	return Vector4{v.X * rh, v.Y * rh, v.Z * rh, v.W * rh}
}

// division
func (v Vector2) Div(rh Vector2) Vector2 {
	return Vector2{v.X / rh.X, v.Y / rh.Y}
}

func (v Vector3) Div(rh Vector3) Vector3 {
	return Vector3{v.X / rh.X, v.Y / rh.Y, v.Z / rh.Z}
}

func (v Vector4) Div(rh Vector4) Vector4 {
	return Vector4{v.X / rh.X, v.Y / rh.Y, v.Z / rh.Z, v.W / rh.W}
}

func (v Vector2) Divf(rh float64) Vector2 {
	return Vector2{v.X / rh, v.Y / rh}
}

func (v Vector3) Divf(rh float64) Vector3 {
	return Vector3{v.X / rh, v.Y / rh, v.Z / rh}
}

func (v Vector4) Divf(rh float64) Vector4 {
	return Vector4{v.X / rh, v.Y / rh, v.Z / rh, v.W / rh}
}

// Dot product
func (v Vector2) Dot(rh Vector2) float64 {
	return v.X*rh.X + v.Y*rh.Y
}

func (v Vector3) Dot(rh Vector3) float64 {
	return v.X*rh.X + v.Y*rh.Y + v.Z*rh.Z
}

func (v Vector4) Dot(rh Vector4) float64 {
	return v.X*rh.X + v.Y*rh.Y + v.Z*rh.Z + v.W*rh.W
}

// Cross product
func (v Vector2) Cross(rh Vector2) float64 {
	return v.X*rh.Y - v.Y*rh.X
}

func (v Vector3) Cross(rh Vector3) Vector3 {
	return Vector3{
		v.Y*rh.Z - v.Z*rh.Y,
		v.Z*rh.X - v.X*rh.Z,
		v.X*rh.Y - v.Y*rh.X,
	}
}

// Length
func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector4) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalize
func (v Vector2) Normalize() Normal2 {
	l := v.Length()
	return Normal2{v.X / l, v.Y / l}
}

func (v Vector3) Normalize() Normal3 {
	l := v.Length()
	return Normal3{v.X / l, v.Y / l, v.Z / l}
}

func (v Vector4) Normalize() Normal4 {
	l := v.Length()
	return Normal4{v.X / l, v.Y / l, v.Z / l, v.W / l}
}

// Angle
func (v Vector2) Angle(rh Vector2) Radian {
	return Radian(math.Acos(v.Dot(rh) / (v.Length() * rh.Length())))
}

func (v Vector3) Angle(rh Vector3) Radian {
	return Radian(math.Acos(v.Dot(rh) / (v.Length() * rh.Length())))
}

func (v Vector4) Angle(rh Vector4) Radian {
	return Radian(math.Acos(v.Dot(rh) / (v.Length() * rh.Length())))
}

// Rotate
func (v Vector2) Rotate(angle Radian) Vector2 {
	return Vector2{
		v.X*math.Cos(float64(angle)) - v.Y*math.Sin(float64(angle)),
		v.X*math.Sin(float64(angle)) + v.Y*math.Cos(float64(angle)),
	}
}

func (v Vector3) Rotate(angle Radian) Vector3 {
	return Vector3{
		v.X*math.Cos(float64(angle)) - v.Y*math.Sin(float64(angle)),
		v.X*math.Sin(float64(angle)) + v.Y*math.Cos(float64(angle)),
		v.Z,
	}
}

func (v Vector4) Rotate(angle Radian) Vector4 {
	return Vector4{
		v.X*math.Cos(float64(angle)) - v.Y*math.Sin(float64(angle)),
		v.X*math.Sin(float64(angle)) + v.Y*math.Cos(float64(angle)),
		v.Z,
		v.W,
	}
}

// Direction to
func (v Vector2) DirectionTo(rh Vector2) Radian {
	return v.Angle(rh)
}

func (v Vector3) DirectionTo(rh Vector3) Radian {
	return v.Angle(rh)
}

func (v Vector4) DirectionTo(rh Vector4) Radian {
	return v.Angle(rh)
}

// Distance to
func (v Vector2) DistanceTo(rh Vector2) float64 {
	return math.Sqrt((v.X-rh.X)*(v.X-rh.X) + (v.Y-rh.Y)*(v.Y-rh.Y))
}

func (v Vector3) DistanceTo(rh Vector3) float64 {
	return math.Sqrt((v.X-rh.X)*(v.X-rh.X) + (v.Y-rh.Y)*(v.Y-rh.Y) + (v.Z-rh.Z)*(v.Z-rh.Z))
}

func (v Vector4) DistanceTo(rh Vector4) float64 {
	return math.Sqrt((v.X-rh.X)*(v.X-rh.X) + (v.Y-rh.Y)*(v.Y-rh.Y) + (v.Z-rh.Z)*(v.Z-rh.Z) + (v.W-rh.W)*(v.W-rh.W))
}
