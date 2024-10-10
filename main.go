package main

import (
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
	"github.com/anaxarchus/MathEngine/geometry/primitive"
)

type Color [4]float64

func main() {
	mesh := primitive.NewMesh("temp/out.png", 1000, 1000)
	mesh.LineWidth = 3.0
	//mesh.DrawRectangle(primitive.NewRectangle(200, 200, 600, 600), Color{1, 0, 0, 1}, primitive.Outline, 5, 0, 0)
	//mesh.DrawRectangle(primitive.NewRectangle(100, 100, 300, 300), Color{1, 1, 1, 1}, primitive.Outline, 3, 0, 0)
	polygon := primitive.NewPolygon(
		vector2.New(200, 200),
		vector2.New(400, 200),
		vector2.New(400, 350),
		vector2.New(350, 350),
		vector2.New(350, 400),
		vector2.New(200, 400),
	)

	circle := primitive.NewCircle(350, 350, 50)

	mesh.AddPolygon(polygon, Color{1, 0, 0, 1}, primitive.Outline, 3.0, 0, 0)
	mesh.SubtractCircle(circle, Color{1, 0, 0, 1}, primitive.Outline, 3.0, 0, 0)
	mesh.AddCircle(primitive.NewCircle(200, 300, 50), Color{1, 0, 0, 1}, primitive.Outline, 3.0, 0, 0)
	mesh.Contour(10.0)
	//mesh.Pocket(8.0)
}
