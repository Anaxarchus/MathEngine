package primitive

import (
	"github.com/Anaxarchus/zero-gdscript/pkg/rect2"
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/fogleman/gg"
)

type Circle struct {
	Center vector2.Vector2
	Radius float64
}

func NewCircle(centerX, centerY, radius float64) Circle {
	return Circle{
		Center: vector2.Vector2{X: centerX, Y: centerY},
		Radius: radius,
	}
}

func (c Circle) Translate(offsetX, offsetY float64) Shape {
	c.Center = c.Center.Add(vector2.Vector2{X: offsetX, Y: offsetY})
	return c
}

func (c Circle) Scale(factor float64) Shape {
	c.Center = c.Center.Mulf(factor)
	c.Radius *= factor
	return c
}

func (c Circle) GetBoundingBox() rect2.Rect2 {
	position := vector2.Vector2{X: 0.0, Y: 0.0}
	size := vector2.Vector2{X: c.Radius * 2, Y: c.Radius * 2}
	rect := rect2.Rect2{Position: position, Size: size}

	return rect
}

func (c Circle) Draw(dc *gg.Context, color [4]float64, lwidth float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawCircle(c.Center.X, c.Center.Y, c.Radius)
	dc.Stroke()
	dc.Pop()
}

func (c Circle) DrawDashed(dc *gg.Context, color [4]float64, lwidth, dashLength, dashInterval float64) {
	dc.Push()
	dc.SetLineWidth(lwidth)
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawCircle(c.Center.X, c.Center.Y, c.Radius)
	dc.SetDash(dashLength, dashInterval)
	dc.Stroke()
	dc.Pop()
}

func (c Circle) DrawFilled(dc *gg.Context, color [4]float64) {
	dc.Push()
	dc.SetRGBA(color[0], color[1], color[2], color[3])
	dc.DrawCircle(c.Center.X, c.Center.Y, c.Radius)
	dc.Fill()
	dc.Pop()
}

func (c Circle) SignedDistance(x, y int) float64 {
	p := vector2.Vector2{X: float64(x), Y: float64(y)}
	return p.Sub(c.Center).Length() - c.Radius
}
