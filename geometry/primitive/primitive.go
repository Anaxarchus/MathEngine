package primitive

import (
	"math"
	"sort"

	"github.com/fogleman/colormap"
	"github.com/fogleman/contourmap"
	"github.com/fogleman/gg"
)

type Shape interface {
	Draw(dc *gg.Context, color [4]float64, lwidth float64)
	DrawDashed(dc *gg.Context, color [4]float64, lwidth float64, dashLength float64, dashInterval float64)
	DrawFilled(dc *gg.Context, color [4]float64)
	Sdf(x, y int) float64
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

type MeshItem struct {
	Shape        Shape
	Color        [4]float64
	Style        DrawStyle
	LineWidth    float64
	DashLength   float64
	DashInterval float64
	Operation    BooleanOperation
}

type Mesh struct {
	Output     string
	Resolution [2]int
	Items      []MeshItem
	LineWidth  float64
}

func NewMesh(output string, resX, resY int) Mesh {
	return Mesh{
		Output:     output,
		Resolution: [2]int{resX, resY},
		Items:      []MeshItem{},
		LineWidth:  3,
	}
}

func (m *Mesh) Update() {
	dc := gg.NewContext(m.Resolution[0], m.Resolution[1])
	for _, item := range m.Items {
		switch item.Style {
		case Outline:
			item.Shape.Draw(dc, item.Color, item.LineWidth)
		case Filled:
			item.Shape.DrawFilled(dc, item.Color)
		case Dashed:
			item.Shape.DrawDashed(dc, item.Color, item.LineWidth, item.DashLength, item.DashInterval)
		}
	}
	dc.SavePNG(m.Output)
}

func (m *Mesh) AddRectangle(rect Rectangle, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        rect,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
	})
}

func (m *Mesh) AddCircle(circle Circle, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        circle,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
	})
}

func (m *Mesh) AddPolygon(polygon Polygon, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        polygon,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
	})
}

func (m *Mesh) SubtractRectangle(rect Rectangle, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        rect,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
		Operation:    Difference,
	})
}

func (m *Mesh) SubtractCircle(circle Circle, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        circle,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
		Operation:    Difference,
	})
}

func (m *Mesh) SubtractPolygon(polygon Polygon, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        polygon,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
		Operation:    Difference,
	})
}

func (m *Mesh) IntersectRectangle(rect Rectangle, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        rect,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
		Operation:    Intersection,
	})
}

func (m *Mesh) IntersectCircle(circle Circle, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        circle,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
		Operation:    Intersection,
	})
}

func (m *Mesh) IntersectPolygon(polygon Polygon, color [4]float64, style DrawStyle, lwidth, dashLength, dashInterval float64) {
	m.Items = append(m.Items, MeshItem{
		Shape:        polygon,
		Color:        color,
		Style:        style,
		LineWidth:    lwidth,
		DashLength:   dashLength,
		DashInterval: dashInterval,
		Operation:    Intersection,
	})
}

func (m *Mesh) Sdf(x, y int) float64 {
	// Sort the Items based on the operation order
	sort.Slice(m.Items, func(i, j int) bool {
		return m.Items[i].Operation < m.Items[j].Operation
	})

	s := math.MaxFloat64 // Initialize the signed distance to a large positive value
	// Perform SDF calculations based on the sorted Items
	for _, item := range m.Items {
		switch item.Operation {
		case Union:
			s = m.fUnion(x, y, s, item.Shape)
		case Difference:
			s = m.fDifference(x, y, s, item.Shape)
		case Intersection:
			s = m.fIntersection(x, y, s, item.Shape)
		}
	}
	return s
}

func (ms *Mesh) Contour(z float64) {
	// Step 1: Create a new drawing context with the specified resolution.
	dc := gg.NewContext(ms.Resolution[0], ms.Resolution[1])

	// Step 2: Set the background color of the canvas to white (RGB: 1,1,1).
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Step 3: Generate a contour map from the signed distance function (SDF).
	m := contourmap.FromFunction(ms.Resolution[0], ms.Resolution[1], ms.Sdf).Closed()

	// Step 4: Extract the contours for the specified SDF level (z).
	contours := m.Contours(z)

	// Step 5: For each contour (a list of points), start a new sub-path.
	for _, c := range contours {
		dc.NewSubPath() // Begin a new contour path (i.e., a continuous line)

		println("contour size: ", len(c))
		c := DouglasPeucker(c, 0.5)
		println("decimate size: ", len(c))

		// Step 6: Loop through each point in the contour and add it to the current path.
		for _, p := range c {
			dc.LineTo(p.X, p.Y) // Draw a line to this point (part of the contour)
			dc.DrawCircle(p.X, p.Y, 1.5)
		}
	}

	// Step 7: Set the fill color for the contour.
	// You can use a fixed color or choose based on the distance (optional).
	// Here, let's use a fixed color (e.g., blue).
	dc.SetRGB(0, 0, 1) // Set the color to blue for the contour fill.
	dc.FillPreserve()  // Fill the interior of the contour

	// Step 8: Set the stroke color and line width for the contour outline.
	dc.SetRGB(0, 0, 0)            // Set outline color to black.
	dc.SetLineWidth(ms.LineWidth) // Set the width of the contour outline.
	dc.Stroke()                   // Draw the outline around the contour

	for _, item := range ms.Items {
		switch item.Style {
		case Outline:
			item.Shape.Draw(dc, item.Color, item.LineWidth)
		case Filled:
			item.Shape.DrawFilled(dc, item.Color)
		case Dashed:
			item.Shape.DrawDashed(dc, item.Color, item.LineWidth, item.DashLength, item.DashInterval)
		}
	}

	// Step 9: Save the final image as a PNG file.
	dc.SavePNG(ms.Output)
}

func (ms *Mesh) Pocket(distance float64) {
	// Step 1: Create a new drawing context with the specified resolution
	dc := gg.NewContext(ms.Resolution[0], ms.Resolution[1])

	// Step 2: Set the background color of the canvas to white.
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Step 3: Generate a contour map from the signed distance function (SDF) modified for pocketing.
	m := contourmap.FromFunction(ms.Resolution[0], ms.Resolution[1], func(x, y int) float64 {
		return ms.Sdf(x, y) - distance // Subtract the specified distance for pocketing
	}).Closed()

	// Step 4: Retrieve the minimum and maximum SDF values for the modified function.
	z0 := m.Min // Minimum SDF value (can be negative inside the shape)
	//z1 := m.Max // Maximum SDF value (positive, far outside the shape)

	// Step 5: Calculate the number of contours based on the range of SDF values.
	// Only count levels for SDF values <= 0 (inner contours).
	numContours := int(math.Ceil(math.Abs(z0) / distance)) // Adjust N as needed
	println("numContours: ", numContours)

	// Step 6: Loop over the number of contour levels (numContours) to generate multiple contour lines.
	for i := 0; i < numContours; i++ {
		// Step 7: Normalize the current index (i) to a value between 0 and 1.
		t := float64(i) / float64(numContours-1)

		// Step 8: Interpolate the current contour level (z) based on t.
		z := z0 + (distance * float64(i))

		// Only continue if z is less than or equal to 0
		if z > 0-distance {
			continue // Skip contours that are greater than 0
		}

		// Step 9: Extract the contours for the current SDF level (z).
		contours := m.Contours(z)

		// Step 10: For each contour (a list of points), start a new sub-path.
		for _, c := range contours {
			dc.NewSubPath() // Begin a new contour path

			// Step 11: Loop through each point in the contour and add it to the current path.
			for _, p := range c {
				dc.LineTo(p.X, p.Y) // Draw a line to this point
			}
		}

		// Step 12: Set the fill color for the contour based on the normalized t value.
		dc.SetColor(colormap.Viridis.At(t))

		// Step 13: Fill the interior of the contour with the selected color.
		dc.FillPreserve()

		// Step 14: Set the stroke color and line width for the contour outline.
		dc.SetRGB(0, 0, 0) // Set outline color to black
		dc.SetLineWidth(3) // Set the width of the contour outline
		dc.Stroke()        // Draw the outline around the contour
	}

	for _, item := range ms.Items {
		switch item.Style {
		case Outline:
			item.Shape.Draw(dc, item.Color, item.LineWidth)
		case Filled:
			item.Shape.DrawFilled(dc, item.Color)
		case Dashed:
			item.Shape.DrawDashed(dc, item.Color, item.LineWidth, item.DashLength, item.DashInterval)
		}
	}

	// Step 15: Save the final image as a PNG file.
	dc.SavePNG(ms.Output)
}

func (m *Mesh) fUnion(x, y int, v float64, s Shape) float64 {
	// Find the minimum signed distance from the shape SDF
	return math.Min(v, s.Sdf(x, y))
}

func (m *Mesh) fDifference(x, y int, v float64, s Shape) float64 {
	// Here, `v` represents the signed distance of the first shape
	// `s.Sdf(x, y)` gives the signed distance of the second shape
	return math.Max(v, -s.Sdf(x, y))
}

func (m *Mesh) fIntersection(x, y int, v float64, s Shape) float64 {
	// Initialize z to a large negative value
	z := -math.MaxFloat64

	// Loop through each shape to find the maximum signed distance
	for _, shape := range m.Items {
		distance := shape.Shape.Sdf(x, y)
		z = math.Max(z, distance)
	}

	// Return the maximum distance
	return z // Return the maximum signed distance
}

// DouglasPeucker simplifies a given set of points using the Douglas-Peucker algorithm
func DouglasPeucker(points []contourmap.Point, epsilon float64) []contourmap.Point {
	// If the points are fewer than 3, return the original points
	if len(points) < 3 {
		return points
	}

	// Get the first and last point
	start := points[0]
	end := points[len(points)-1]

	// Find the index of the point with the maximum distance from the line segment
	index := -1
	maxDistance := 0.0

	for i := 1; i < len(points)-1; i++ {
		distance := perpendicularDistance(points[i], start, end)
		if distance > maxDistance {
			maxDistance = distance
			index = i
		}
	}

	// If the maximum distance is greater than the tolerance, recursively simplify
	var result []contourmap.Point
	if maxDistance > epsilon {
		// Recursively simplify the segments before and after the index
		firstLine := DouglasPeucker(points[:index+1], epsilon)
		secondLine := DouglasPeucker(points[index:], epsilon)

		// Combine the results, excluding the last point of the first segment to avoid duplication
		result = append(result, firstLine[:len(firstLine)-1]...)
		result = append(result, secondLine...)
	} else {
		// If the max distance is within the tolerance, return the endpoints
		result = []contourmap.Point{start, end}
	}

	return result
}

// perpendicularDistance calculates the distance from a point to a line segment defined by two endpoints
func perpendicularDistance(pt, start, end contourmap.Point) float64 {
	// Calculate the length of the line segment
	lineLength := math.Sqrt(math.Pow(end.X-start.X, 2) + math.Pow(end.Y-start.Y, 2))
	if lineLength == 0 {
		return math.Sqrt(math.Pow(pt.X-start.X, 2) + math.Pow(pt.Y-start.Y, 2)) // Distance to start point
	}

	// Calculate the projection of point onto the line segment
	t := ((pt.X-start.X)*(end.X-start.X) + (pt.Y-start.Y)*(end.Y-start.Y)) / (lineLength * lineLength)

	// Clamp the projection value to the line segment
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	// Calculate the nearest point on the line segment
	nearest := contourmap.Point{
		X: start.X + t*(end.X-start.X),
		Y: start.Y + t*(end.Y-start.Y),
	}

	// Return the distance from the point to the nearest point on the line segment
	return math.Sqrt(math.Pow(pt.X-nearest.X, 2) + math.Pow(pt.Y-nearest.Y, 2))
}
