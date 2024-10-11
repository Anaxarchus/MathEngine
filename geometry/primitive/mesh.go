package primitive

import (
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/anaxarchus/MathEngine/geometry/algorithms/simplify"
	"github.com/fogleman/contourmap"
	"github.com/fogleman/gg"
)

type Mesh struct {
	Polygon      Polygon
	Position     vector2.Vector2
	Color        [4]float64
	OutlineWidth float64
	OutlineColor [4]float64
	Filled       bool
}

// Constructors
func NewMesh() *Mesh {
	return &Mesh{}
}

func (m Mesh) Scale(factor float64) Mesh {
	for i, p := range m.Polygon {
		m.Polygon[i] = p.Mulf(factor)
	}
	return m
}

func (m *Mesh) GetBoundingBox() vector2.Vector2 {
	minX, minY := m.Polygon[0].X, m.Polygon[0].Y
	maxX, maxY := m.Polygon[0].X, m.Polygon[0].Y

	for _, p := range m.Polygon {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return vector2.Vector2{X: maxX - minX, Y: maxY - minY}
}

// Member functions
func (m *Mesh) Draw(context *gg.Context) {
	if m.Filled {
		m.Polygon.DrawFilled(context, m.Color)
	}
	m.Polygon.Draw(context, m.OutlineColor, m.OutlineWidth)
}

func (m *Mesh) SignedDistance(x, y int) float64 {
	return m.Polygon.SignedDistance(x, y)
}

func (m *Mesh) AddShape(s Shape) {
	bg := BooleanGroup([]Shape{m.Polygon, s})
	m.Polygon = Polygon(bg.GetContour(0, bg.UnionDistance))
}

func (m *Mesh) SubtractShape(s Shape) {
	bg := BooleanGroup([]Shape{m.Polygon, s})
	m.Polygon = Polygon(bg.GetContour(0, bg.DifferenceDistance))
}

func (m *Mesh) IntersectShape(s Shape) {
	bg := BooleanGroup([]Shape{m.Polygon, s})
	m.Polygon = Polygon(bg.GetContour(0, bg.IntersectionDistance))
}

func simplifyContours(c []contourmap.Contour, epsilon float64) [][]vector2.Vector2 {
	var result [][]vector2.Vector2
	for _, c := range c {

		floats := contourPointsToFloatSlice(c)
		simplified := simplify.DouglasPeucker(floats, epsilon)
		result = append(result, floatsToVector2Slice(simplified))
	}
	return result
}

func floatsToVector2Slice(f [][2]float64) []vector2.Vector2 {
	var points []vector2.Vector2
	for _, p := range f {
		points = append(points, vector2.Vector2{X: p[0], Y: p[1]})
	}
	return points
}

func contourPointsToFloatSlice(c []contourmap.Point) [][2]float64 {
	var points [][2]float64
	for _, p := range c {
		points = append(points, contourPointToFloatSlice(p))
	}
	return points
}

func contourPointToFloatSlice(c contourmap.Point) [2]float64 {
	return [2]float64{float64(c.X), float64(c.Y)}
}
