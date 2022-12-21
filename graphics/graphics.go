package graphics

import (
	"kotongo/components"

	"github.com/go-gl/gl/v2.1/gl"
)

// TODO Replace into a shapes file
type ShapeType2D int16

const (
	RightTriangle ShapeType2D = iota
	IsoscelesTriangle
	Square
	Rectangle
	Circle
	Ellipse
	Line
)

var (
	IsoscelesTrianglePoints = []float32{
		0, 0.1, 0, // top
		-0.1, -0.1, 0, // left
		0.1, -0.1, 0, // right
	}
	RightTrianglePoints = []float32{
		0, 0.1, 0, // top
		0, 0, 0, // left
		0.1, 0, 0, // right

	}
	RectanglePoints = components.Polygon{
		Vertices: []float32{
			0.1, 0.1, 0, //top right
			0.1, 0, 0, //bottom right
			0, 0, 0, //bottom left
			0, 0.1, 0, //top left
		},
		Indices: []uint32{
			0, 1, 3, //first triangle
			1, 2, 3, //second triangle
		},
	}
)

// TODO Disconnect VAO (glBindVertexArray(0);)
// Create vertex array object
func VAO(renderers ...components.Renderer) uint32 {

	//Vertex array object
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	for _, obj := range renderers {
		obj.VBO()

		if len(obj.Indices) > 1 {
			obj.EBO()
		}
	}

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	return vao
}
