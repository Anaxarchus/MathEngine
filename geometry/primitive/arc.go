package primitive

import (
	"math"

	"github.com/Anaxarchus/zero-gdscript/pkg/rect2"
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/fogleman/gg"
)

type Arc struct {
	Circle               Circle
	AngleStart, AngleEnd float64
}

func NewArc(centerX, centerY, radius, angleStart, angleEnd float64) Arc {
	return Arc{
		Circle:     NewCircle(centerX, centerY, radius),
		AngleStart: angleStart,
		AngleEnd:   angleEnd,
	}
}

func ArcFromPoints(points []vector2.Vector2) *Arc {
	origin, radius := fitArc(points)
	a1 := points[0]
	a2 := points[len(points)-1]
	return &Arc{
		Circle: Circle{
			Center: origin,
			Radius: radius,
		},
		AngleStart: origin.AngleToPoint(a1),
		AngleEnd:   origin.AngleToPoint(a2),
	}
}

func (a Arc) Translate(offsetX, offsetY float64) Shape {
	a.Circle = a.Circle.Translate(offsetX, offsetY).(Circle)
	return a
}

func (a Arc) Scale(factor float64) Shape {
	a.Circle = a.Circle.Scale(factor).(Circle)
	return a
}

func (a Arc) GetBoundingBox() rect2.Rect2 {
	return rect2.Rect2{}
}

func (a Arc) Draw(dc *gg.Context, color [4]float64, lwidth float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawArc(a.Circle.Center.X, a.Circle.Center.Y, a.Circle.Radius, a.AngleStart, a.AngleEnd)
	dc.Stroke()
	dc.Pop()
}

func (a Arc) DrawDashed(dc *gg.Context, color [4]float64, lwidth, dashLength, dashInterval float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawArc(a.Circle.Center.X, a.Circle.Center.Y, a.Circle.Radius, a.AngleStart, a.AngleEnd)
	dc.SetDash(dashLength, dashInterval)
	dc.Stroke()
	dc.Pop()
}

func (a Arc) DrawFilled(dc *gg.Context, color [4]float64) {
	dc.Push()
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawArc(a.Circle.Center.X, a.Circle.Center.Y, a.Circle.Radius, a.AngleStart, a.AngleEnd)
	dc.Fill()
	dc.Pop()
}

func (a Arc) SignedDistance(x, y int) float64 {
	return 0.0
}

func (a Arc) GetArcBetweenPoints(start, end vector2.Vector2) *Arc {
	a.AngleStart = a.AngleToPoint(start)
	a.AngleStart = a.AngleToPoint(end)
	return &a
}

func (a *Arc) Project(point vector2.Vector2) vector2.Vector2 {
	return a.Circle.Center.Add(a.Circle.Center.DirectionTo(point).Mulf(a.Circle.Radius))
}

func (a *Arc) AngleToPoint(point vector2.Vector2) float64 {
	return a.Circle.Center.AngleToPoint(point)
}

// ArcDirection determines the direction of the arc: 1 for clockwise, -1 for counter-clockwise.
// It takes lastPosition, currentPosition and arcOrigin, all as Vector3.
func ArcDirection(lastPosition, position, arcOrigin vector2.Vector2) int {
	// Convert to coordinates relative to arcOrigin
	lastRel := lastPosition.Sub(arcOrigin)
	currentRel := position.Sub(arcOrigin)

	// Calculate angles from arcOrigin
	lastAngle := lastRel.Angle()
	currentAngle := currentRel.Angle()

	// Calculate angle difference and determine direction
	angleDifference := currentAngle - lastAngle

	// Normalize the angle to be within the range -π to π
	if angleDifference > math.Pi {
		angleDifference -= 2 * math.Pi
	} else if angleDifference < -math.Pi {
		angleDifference += 2 * math.Pi
	}

	// Determine the direction based on the angle difference
	if angleDifference > 0 {
		return -1 // CCW
	} else {
		return 1 // CW
	}
}

func (a *Arc) Discretize(maxInterval float64, minSteps int) []vector2.Vector2 {
	// Calculate the total angle span of the arc
	totalAngle := a.AngleEnd - a.AngleStart
	if totalAngle < 0 {
		totalAngle += 2 * math.Pi
	}

	// Calculate the total length of the arc
	arcLength := a.Circle.Radius * totalAngle

	// Calculate the minimum number of steps required based on maxInterval
	minStepsBasedOnInterval := int(math.Ceil(arcLength / maxInterval))

	// Ensure the number of steps is at least minSteps and is odd
	numSteps := minStepsBasedOnInterval
	if numSteps < minSteps {
		numSteps = minSteps
	}
	if numSteps%2 == 0 {
		numSteps++
	}

	// Calculate the angle step
	angleStep := totalAngle / float64(numSteps-1)

	// Generate the points along the arc
	points := make([]vector2.Vector2, 0, numSteps)
	for i := 0; i < numSteps; i++ {
		angle := a.AngleStart + float64(i)*angleStep
		x := a.Circle.Center.X + a.Circle.Radius*math.Cos(angle)
		y := a.Circle.Center.Y + a.Circle.Radius*math.Sin(angle)
		points = append(points, vector2.Vector2{X: x, Y: y})
	}

	return points
}

// ProcessArcs calculates the center and radius of a circle that passes through the first,
// last, and a nearest point to the midpoint of the first and last points from a given
// slice of Vector2 points that represent an arc.
func fitArc(arcPoints []vector2.Vector2) (vector2.Vector2, float64) {
	e1 := arcPoints[0]                // Start point of the arc
	e2 := arcPoints[len(arcPoints)-1] // End point of the arc

	// Calculate the midpoint of the line segment from e1 to e2
	mp := e1.Add(e2).Divf(2.0)

	// Find the point nearest to the midpoint among the arc points
	minDist := math.MaxFloat64
	var np vector2.Vector2
	for _, point := range arcPoints {
		dist := mp.DistanceTo(point)
		if dist < minDist {
			minDist = dist
			np = point
		}
	}

	// Calculate the distances of the sides of the triangle formed by e1, e2, and np
	a := e1.DistanceTo(e2) // Distance between end points
	b := e1.DistanceTo(np) // Distance from start point to nearest point
	c := e2.DistanceTo(np) // Distance from end point to nearest point

	// Calculate the radius of the circumscribed circle using Heron's formula
	// and the area of a triangle calculation
	product := (a + b + c) * (a + b - c) * (a - b + c) * (b + c - a)
	if product <= 0 {
		return np, 0 // If product is non-positive, radius calculation fails
	}

	abc := a * b * c

	rad := abc / math.Sqrt(product)

	// Calculate the center of the circumscribed circle
	A := e1.Sub(np)
	B := e2.Sub(np)
	//C := e1.Sub(e2)
	D := A.Dot(e1.Add(np).Divf(2.0))
	E := B.Dot(e2.Add(np).Divf(2.0))

	center := vector2.Vector2{
		X: (D*B.Y - E*A.Y) / (A.X*B.Y - B.X*A.Y),
		Y: (A.X*E - B.X*D) / (A.X*B.Y - B.X*A.Y),
	}

	return center, rad
}
