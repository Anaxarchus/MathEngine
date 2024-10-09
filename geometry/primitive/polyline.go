package primitive

type PolyLine []Line

// GetPoints extracts all unique points from the PolyLine
func (p PolyLine) GetPoints() []Vector2 {
	points := make([]Vector2, 0)
	if len(p) == 0 {
		return points
	}

	// Add the first point of the first line
	points = append(points, p[0][0])

	// Add the second point of each line
	for _, l := range p {
		points = append(points, l[1])
	}
	return points
}

// GetArea calculates the area of the closed polygon created by the polyline points
func (p PolyLine) Area() float64 {
	// Convert the points of the PolyLine to a PolyPoint and calculate the area
	points := p.GetPoints()

	// Ensure the polyline is closed by adding the first point to the end if necessary
	if points[0] != points[len(points)-1] {
		points = append(points, points[0])
	}

	return PolyPoint(points).Area()
}
