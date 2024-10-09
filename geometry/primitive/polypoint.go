package primitive

import "math"

type PolyPoint []Vector2

// GetArea calculates the area of the n-gon using the Shoelace theorem
func (p PolyPoint) Area() float64 {
	n := len(p)
	if n < 3 {
		return 0 // Not a valid polygon
	}

	area := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n // The next vertex (with wrapping around)
		area += p[i].X * p[j].Y
		area -= p[i].Y * p[j].X
	}

	return math.Abs(area) / 2.0
}

// BoundingBox returns the minimum bounding rectangle for the n-gon
func (p PolyPoint) BoundingBox() Rectangle {
	if len(p) == 0 {
		return Rectangle{}
	}

	minX, minY := p[0].X, p[0].Y
	maxX, maxY := minX, minY

	for _, v := range p {
		minX = math.Min(minX, v.X)
		minY = math.Min(minY, v.Y)
		maxX = math.Max(maxX, v.X)
		maxY = math.Max(maxY, v.Y)
	}

	return Rectangle{
		{minX, minY},
		{maxX, maxY},
	}
}

// Centroid calculates the centroid of the n-gon
func (p PolyPoint) Centroid() Vector2 {
	n := len(p)
	if n < 3 {
		return Vector2{} // Not a valid polygon
	}

	area := 0.0
	cx, cy := 0.0, 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n // The next vertex (with wrapping around)
		f := p[i].X*p[j].Y - p[j].X*p[i].Y
		area += f
		cx += (p[i].X + p[j].X) * f
		cy += (p[i].Y + p[j].Y) * f
	}

	area /= 2.0
	cx /= 6.0 * area
	cy /= 6.0 * area

	return Vector2{cx, cy}
}

// DistanceTo calculates the distance from the n-gon to a
// point using the minimum distance to any of its edges
func (p PolyPoint) DistanceTo(v Vector2) float64 {
	n := len(p)
	if n < 3 {
		return 0 // Not a valid polygon
	}

	minDist := math.Inf(1)
	for i := 0; i < n; i++ {
		j := (i + 1) % n // The next vertex (with wrapping around)
		dist := p[i].DistanceToSegment(p[j], v)
		minDist = math.Min(minDist, dist)
	}

	return minDist
}
