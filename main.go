package main

import (
	"fmt"

	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/anaxarchus/MathEngine/geometry/primitive"
	"github.com/fogleman/gg"
)

type Color [4]float64

func main() {

}

func DumbShit() {
	cc := gg.NewContext(1000, 1000)

	mesh := primitive.NewMesh()
	mesh.OutlineWidth = 3.0
	mesh.OutlineColor = Color{1, 0, 0, 1}
	mesh.Filled = false
	mesh.Polygon = primitive.NewPolygon(
		vector2.New(200.129, 200.129),
		vector2.New(400.433, 200.129),
		vector2.New(400.433, 350.912),
		vector2.New(350.912, 350.912),
		vector2.New(350.912, 400.433),
		vector2.New(200.129, 400.433),
	)

	for _, pt := range mesh.Polygon {
		fmt.Println("point: ", pt.X, ", ", pt.Y)
	}

	mesh.Draw(cc)

	circle := primitive.NewCircle(350, 350, 50)
	mesh.AddShape(circle)
	mesh.OutlineColor = Color{0, 1, 0, 1}
	mesh.Draw(cc)

	for _, pt := range mesh.Polygon {
		fmt.Println("point: ", pt.X, ", ", pt.Y)
	}

	cc.SavePNG("temp/out.png")
	//mesh.Pocket(8.0)
}
