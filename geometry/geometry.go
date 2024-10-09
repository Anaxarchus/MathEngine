package geometry

import "github.com/anaxarchus/MathEngine/geometry/primitive"

func Vector2(point [2]float64) primitive.Vector2 {
	return primitive.Vector2{X: point[0], Y: point[1]}
}

// Lines
func Line(points [2][2]float64) primitive.Line {
	return primitive.Line([2]primitive.Vector2{Vector2(points[0]), Vector2(points[1])})
}

// Triangles
func Triangle(points [3][2]float64) primitive.Triangle {
	return primitive.Triangle([3]primitive.Vector2{Vector2(points[0]), Vector2(points[1]), Vector2(points[2])})
}

// Rectangles
func Rectangle(points [4][2]float64) primitive.Rectangle {
	return primitive.Rectangle([4]primitive.Vector2{Vector2(points[0]), Vector2(points[1]), Vector2(points[2]), Vector2(points[3])})
}

// PolyPoints
func PolyPoint(points ...[2]float64) primitive.PolyPoint {
	pp := make(primitive.PolyPoint, len(points))
	for i, p := range points {
		pp[i] = Vector2(p)
	}
	return pp
}

func PolyLine(lines ...[2][2]float64) primitive.PolyLine {
	pl := make(primitive.PolyLine, len(lines))
	for i, l := range lines {
		pl[i] = primitive.Line([2]primitive.Vector2{Vector2(l[0]), Vector2(l[1])})
	}
	return pl
}

func Circle(center [2]float64, radius float64) primitive.Circle {
	return primitive.Circle{Center: Vector2(center), Radius: radius}
}

func Ellipse(center [2]float64, radiusX, radiusY float64) primitive.Ellipse {
	return primitive.Ellipse{Center: Vector2(center), RadiusX: radiusX, RadiusY: radiusY}
}

func Arc(center [2]float64, radius, startAngle, endAngle float64) primitive.Arc {
	return primitive.Arc{Center: Vector2(center), Radius: radius, StartAngle: startAngle, EndAngle: endAngle}
}
