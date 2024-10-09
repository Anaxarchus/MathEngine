package primitive

type Triangle [3]Vector2

// GetArea calculates the area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * ((t[1].X-t[0].X)*(t[2].Y-t[0].Y) - (t[2].X-t[0].X)*(t[1].Y-t[0].Y))
}
