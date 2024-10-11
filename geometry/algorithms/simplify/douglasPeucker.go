package simplify

import (
	"math"
)

// DouglasPeucker simplifies a given set of points using the Douglas-Peucker algorithm
func DouglasPeucker(points [][2]float64, epsilon float64) [][2]float64 {
	// If the points are fewer than 3, return the original points
	if len(points) < 3 {
		return points
	}

	// Get the first and last point
	start := points[0]
	end := points[len(points)-1]

	// Find the index of the point with the maximum distance from the line segment
	index := -1
	maxDistance := 0.0

	for i := 1; i < len(points)-1; i++ {
		distance := perpendicularDistance(points[i], start, end)
		if distance > maxDistance {
			maxDistance = distance
			index = i
		}
	}

	// If the maximum distance is greater than the tolerance, recursively simplify
	var result [][2]float64
	if maxDistance > epsilon {
		// Recursively simplify the segments before and after the index
		firstLine := DouglasPeucker(points[:index+1], epsilon)
		secondLine := DouglasPeucker(points[index:], epsilon)

		// Combine the results, excluding the last point of the first segment to avoid duplication
		result = append(result, firstLine[:len(firstLine)-1]...)
		result = append(result, secondLine...)
	} else {
		// If the max distance is within the tolerance, return the endpoints
		result = [][2]float64{start, end}
	}

	return result
}

// perpendicularDistance calculates the distance from a point to a line segment defined by two endpoints
func perpendicularDistance(pt, start, end [2]float64) float64 {
	// Calculate the length of the line segment
	lineLength := math.Sqrt(math.Pow(end[0]-start[0], 2) + math.Pow(end[1]-start[1], 2))
	if lineLength == 0 {
		return math.Sqrt(math.Pow(pt[0]-start[0], 2) + math.Pow(pt[1]-start[1], 2)) // Distance to start point
	}

	// Calculate the projection of point onto the line segment
	t := ((pt[0]-start[0])*(end[0]-start[0]) + (pt[1]-start[1])*(end[1]-start[1])) / (lineLength * lineLength)

	// Clamp the projection value to the line segment
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	// Calculate the nearest point on the line segment
	nearest := [2]float64{
		start[0] + t*(end[0]-start[0]),
		start[1] + t*(end[1]-start[1]),
	}

	// Return the distance from the point to the nearest point on the line segment
	return math.Sqrt(math.Pow(pt[0]-nearest[0], 2) + math.Pow(pt[1]-nearest[1], 2))
}
