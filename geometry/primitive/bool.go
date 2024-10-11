package primitive

import (
	"fmt"
	"math"

	"github.com/Anaxarchus/zero-gdscript/pkg/rect2"
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/fogleman/contourmap"
)

type BooleanGroup []Shape

func (bg BooleanGroup) GetBoundingBox() rect2.Rect2 {
	bb := bg[0].GetBoundingBox()
	for _, shape := range bg {
		bb = bb.Merge(shape.GetBoundingBox())
	}
	return bb
}

func (bg *BooleanGroup) UnionDistance(x, y int) float64 {
	min := math.Inf(1) // Initialize to positive infinity
	for _, shape := range *bg {
		d := shape.SignedDistance(x, y)
		if d < min {
			min = d
		}
	}
	return min
}

func (bg *BooleanGroup) IntersectionDistance(x, y int) float64 {
	max := math.Inf(-1) // Initialize to negative infinity
	for _, shape := range *bg {
		d := shape.SignedDistance(x, y)
		if d > max {
			max = d
		}
	}
	return max
}

func (bg *BooleanGroup) DifferenceDistance(x, y int) float64 {
	min := math.Inf(1) // Initialize to positive infinity
	for _, shape := range *bg {
		d := shape.SignedDistance(x, y)
		if d < min {
			min = d
		}
	}
	return -min // Return the negated minimum distance
}

// Utility functions
func (bg *BooleanGroup) GetContour(z float64, dFunc contourmap.Function) Polygon {
	const Scale = 10.0
	r := (*bg)[0].GetBoundingBox()
	for _, shape := range *bg {
		r = r.Merge(shape.GetBoundingBox())
	}
	padding := 20.0
	r = r.Grow(padding)
	for i, _ := range *bg {
		(*bg)[i] = (*bg)[i].Translate(-r.Position.X+padding, -r.Position.Y+padding).Scale(Scale)
		printRect((*bg)[i].GetBoundingBox())
	}

	fmt.Println("Bounding box, size:", r.Size.X, ", ", r.Size.Y, ", position:", r.Position.X, ", ", r.Position.Y)
	//panic("Not implemented")

	cRes := [2]int{int(math.Ceil(r.Size.X * Scale)), int(math.Ceil(r.Size.Y * Scale))}

	cMap := contourmap.FromFunction(cRes[0], cRes[1], dFunc).Closed()
	contours := cMap.Contours(z)
	println("contour size: ", len(contours))
	simplified := simplifyContours(contours, 0.1)

	boundaryTolerance := 0.001 // Larger tolerance for boundary detection
	minBoundaryPoints := 0.2   // Discard if more than 20% of points are on the boundary

	var largestContour []vector2.Vector2

	for _, s := range simplified {
		boundaryPoints := 0

		// Count how many points are close to the boundary
		for _, pt := range s {
			if pt.X < boundaryTolerance || pt.Y < boundaryTolerance ||
				pt.X > float64(cRes[0])-boundaryTolerance || pt.Y > float64(cRes[1])-boundaryTolerance {
				boundaryPoints++
			}
		}

		boundaryRatio := float64(boundaryPoints) / float64(len(s))

		// Discard contours with too many boundary points
		if boundaryRatio < minBoundaryPoints {
			if len(s) > len(largestContour) {
				largestContour = s // Keep the largest valid contour
			}
		} else {
			println("Ignoring boundary contour with size:", len(s))
		}
	}

	if len(largestContour) > 0 {
		p := Polygon(largestContour).Scale(1.0/Scale).Translate(r.Position.X, r.Position.Y).(Polygon)
		return p
	}

	// Return an empty slice if no valid contours were found
	return []vector2.Vector2{}
}

func printRect(rect rect2.Rect2) {
	fmt.Println("Rect: ", rect.Position.X, ", ", rect.Position.Y, ", ", rect.Size.X, ", ", rect.Size.Y)
}
