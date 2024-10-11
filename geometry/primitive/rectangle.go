package primitive

import (
	"github.com/Anaxarchus/zero-gdscript/pkg/rect2"
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/fogleman/gg"
)

type Rectangle struct {
	Offset vector2.Vector2
	Size   vector2.Vector2
}

func NewRectangle(offsetX, offsetY, sizeX, sizeY float64) Rectangle {
	return Rectangle{
		Offset: vector2.Vector2{X: offsetX, Y: offsetY},
		Size:   vector2.Vector2{X: sizeX, Y: sizeY},
	}
}

func (r Rectangle) Translate(offsetX, offsetY float64) Shape {
	r.Offset = r.Offset.Add(vector2.Vector2{X: offsetX, Y: offsetY})
	return r
}

func (r Rectangle) Scale(factor float64) Shape {
	r.Offset = r.Offset.Mulf(factor)
	r.Size = r.Size.Mulf(factor)
	return r
}

func (r Rectangle) GetBoundingBox() rect2.Rect2 {
	return rect2.New(r.Offset, r.Size)
}

func (r Rectangle) Draw(dc *gg.Context, color [4]float64, lwidth float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawRectangle(r.Offset.X, r.Offset.Y, r.Size.X, r.Size.Y)
	dc.Stroke()
	dc.Pop()
}

func (r Rectangle) DrawDashed(dc *gg.Context, color [4]float64, lwidth, dashLength, dashInterval float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawRectangle(r.Offset.X, r.Offset.Y, r.Size.X, r.Size.Y)
	dc.SetDash(dashLength, dashInterval)
	dc.Stroke()
	dc.Pop()
}

func (r Rectangle) DrawFilled(dc *gg.Context, color [4]float64) {
	dc.Push()
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawRectangle(r.Offset.X, r.Offset.Y, r.Size.X, r.Size.Y)
	dc.Fill()
	dc.Pop()
}

func (r Rectangle) SignedDistance(x, y int) float64 {
	// Calculate the position of the point relative to the rectangle center.
	// Assuming r.Center is the center of the rectangle.
	p := vector2.Vector2{X: float64(x), Y: float64(y)}
	q := p.Sub(r.Offset.Add(r.Size.Mulf(0.5))).ABS().Sub(r.Size.Mulf(0.5)) // Subtract half-extents (rectangle size / 2)

	// Calculate the distance to the rectangle.
	outsideDist := q.Maxf(0.0).Length()   // Distance when outside rectangle
	insideDist := min(max(q.X, q.Y), 0.0) // Negative distance when inside rectangle

	// Return the total signed distance.
	return outsideDist + insideDist
}
