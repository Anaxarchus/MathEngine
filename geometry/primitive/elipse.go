package primitive

import "math"

type Ellipse struct {
	Center  Vector2
	RadiusX float64
	RadiusY float64
}

func (e *Ellipse) Area() float64 {
	return math.Pi * e.RadiusX * e.RadiusY
}

func (e *Ellipse) BoundingBox() Rectangle {
	return Rectangle{
		{e.Center.X - e.RadiusX, e.Center.Y - e.RadiusY},
		{e.Center.X + e.RadiusX, e.Center.Y + e.RadiusY},
	}
}

func (e *Ellipse) Centroid() Vector2 {
	return e.Center
}

func (e *Ellipse) DistanceTo(v Vector2) float64 {
	return math.Abs(e.Center.DistanceTo(v) - e.RadiusX)
}
