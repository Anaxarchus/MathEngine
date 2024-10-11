package primitive

import (
	"github.com/Anaxarchus/zero-gdscript/pkg/rect2"
	"github.com/fogleman/gg"
)

type Shape interface {
	Draw(dc *gg.Context, color [4]float64, lwidth float64)
	DrawDashed(dc *gg.Context, color [4]float64, lwidth float64, dashLength float64, dashInterval float64)
	DrawFilled(dc *gg.Context, color [4]float64)
	SignedDistance(x, y int) float64
	GetBoundingBox() rect2.Rect2
	Scale(factor float64) Shape
	Translate(offsetX, offsetY float64) Shape
}

type DrawStyle int

const (
	Outline DrawStyle = iota
	Filled
	Dashed
)

type BooleanOperation int

const (
	Union BooleanOperation = iota
	Difference
	Intersection
)
