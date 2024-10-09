package primitive

type Shape interface {
	Area() float64
	BoundingBox() Rectangle
	Centroid() Vector2
	DistanceTo(point Vector2) float64
}

type Ngon interface {
	GetPointCount() int
	GetEdgeCount() int
	GetVertexAngles() []Radian
	GetVertexAngle(index int) Radian
	GetVertexNormals() []Normal2
	GetVertexNormal(index int) Normal2
	GetEdgeNormals() []Normal2
	GetEdgeNormal(index int) Normal2
	GetEdges() PolyLine
	GetEdge(index int) Line
	GetNearestEdge(point Vector2) Vector2
}

type Path interface {
	Length() float64
	Walk(distance float64, wrap bool) []Vector2
	BoundingBox() Rectangle
	IsClosed() bool
	IsClockwise() bool
	Invert()
	Offset(distance float64) Path
	GetPointAtDistance(distance float64) Vector2
	GetLineAtDistance(distance float64) Line
}

func NewPolyline(lines ...Line) PolyLine {
	return PolyLine(lines)
}
