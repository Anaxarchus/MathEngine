package primitive

import "math"

type Circle struct {
	Center Vector2
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) BoundingBox() Rectangle {
	return Rectangle{
		{c.Center.X - c.Radius, c.Center.Y - c.Radius},
		{c.Center.X + c.Radius, c.Center.Y + c.Radius},
	}
}

func (c *Circle) Centroid() Vector2 {
	return c.Center
}

func (c *Circle) DistanceTo(v Vector2) float64 {
	return math.Abs(c.Center.DistanceTo(v) - c.Radius)
}
