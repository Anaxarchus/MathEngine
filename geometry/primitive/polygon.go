package primitive

import (
	"math"

	zerogdscript "github.com/Anaxarchus/zero-gdscript"
	"github.com/Anaxarchus/zero-gdscript/pkg/rect2"
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/fogleman/gg"
)

type Polygon []vector2.Vector2

func NewPolygon(vertices ...vector2.Vector2) Polygon {
	return Polygon(vertices)
}

func (p Polygon) Translate(offsetX, offsetY float64) Shape {
	for i, v := range p {
		p[i] = v.Add(vector2.Vector2{X: offsetX, Y: offsetY})
	}
	return p
}

func (p Polygon) Scale(factor float64) Shape {
	for i, v := range p {
		p[i] = v.Mulf(factor)
	}
	return p
}

func (p Polygon) GetBoundingBox() rect2.Rect2 {
	position := p[0]
	size := vector2.Vector2{X: 0.0, Y: 0.0}
	rect := rect2.Rect2{Position: position, Size: size}

	for _, pt := range p {
		rect = rect.Expand(pt)
	}

	return rect
}

func (p Polygon) Draw(dc *gg.Context, color [4]float64, lwidth float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.MoveTo(p[0].X, p[0].Y)
	for _, v := range p[1:] {
		dc.LineTo(v.X, v.Y)
	}
	dc.ClosePath()
	dc.Stroke()
	dc.Pop()
}

func (p Polygon) DrawDashed(dc *gg.Context, color [4]float64, lwidth, dashLength, dashInterval float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.MoveTo(p[0].X, p[0].Y)
	for _, v := range p[1:] {
		dc.LineTo(v.X, v.Y)
	}
	dc.ClosePath()
	dc.SetDash(dashLength, dashInterval)
	dc.Stroke()
	dc.Pop()
}

func (p Polygon) DrawFilled(dc *gg.Context, color [4]float64) {
	dc.Push()
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.MoveTo(p[0].X, p[0].Y)
	for _, v := range p[1:] {
		dc.LineTo(v.X, v.Y)
	}
	dc.ClosePath()
	dc.Fill()
	dc.Pop()
}

func (p Polygon) SignedDistance(x, y int) float64 {
	return SdPolygon(p, vector2.Vector2{X: float64(x), Y: float64(y)})
}

func SdPolygon(vertices []vector2.Vector2, p vector2.Vector2) float64 {
	d := p.Sub(vertices[0]).Dot(p.Sub(vertices[0]))
	s := 1.0

	N := len(vertices)
	for i := 0; i < N; i++ {
		j := (i + N - 1) % N              // Previous vertex index
		e := vertices[j].Sub(vertices[i]) // Edge vector
		w := p.Sub(vertices[i])           // Vector from vertex to point
		// Projection of w onto e
		proj := e.Dot(w) / e.Dot(e)
		// Closest point on edge
		b := w.Sub(e.Mulf(zerogdscript.Clampf(proj, 0.0, 1.0)))

		// Update distance
		d = math.Min(d, b.Dot(b))

		// Check for winding number
		c1 := p.Y >= vertices[i].Y
		c2 := p.Y < vertices[j].Y
		c3 := e.X*w.Y > e.Y*w.X
		if (c1 && c2 && c3) || (!c1 && !c2 && !c3) {
			s *= -1.0
		}
	}
	return s * math.Sqrt(d)
}
