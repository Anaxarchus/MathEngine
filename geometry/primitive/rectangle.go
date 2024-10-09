package primitive

type Rectangle [4]Vector2

// GetArea calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return (r[2].X - r[0].X) * (r[2].Y - r[0].Y)
}
