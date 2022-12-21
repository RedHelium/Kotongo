package components

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl64"
)

// Application data
type App struct {
	width   int
	height  int
	title   string
	shaders []shaderInfo // List of shaders for this program
}

// Shader data
type shaderInfo struct {
	stype  uint32 // OpenGL shader type (VERTEX_SHADER, FRAGMENT_SHADER, or GEOMETRY_SHADER)
	source string // Shader source code
}

// Global object data
type Entity struct {
	name      string
	renderer  Renderer
	transform Transform
}

// Transformation data
type Transform struct {
	position mgl64.Vec3
	rotation mgl64.Quat
	scale    mgl64.Vec3
}

type Polygon struct {
	Vertices []float32
	Indices  []uint32
}

// Render object data
type Renderer struct {
	DrawType uint32 // GL_STATIC_DRAW, GL_DYNAMIC_DRAW, GL_STREAM_DRAW
	Vertices []float32
	Indices  []uint32
}

func (renderer *Renderer) GetVertices() []float32 {
	return renderer.Vertices
}

func (renderer *Renderer) GetDrawType() uint32 {
	return renderer.DrawType
}

// Create vertex buffer object
func (renderer *Renderer) VBO() uint32 {

	var vbo uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(renderer.Vertices), gl.Ptr(renderer.Vertices), renderer.DrawType)

	return vbo
}

// Create Element buffer object
func (renderer *Renderer) EBO() uint32 {

	var ebo uint32

	if len(renderer.Indices) <= 0 {
		log.Fatal("Renderer indices is empty!")
	}

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(renderer.Indices), gl.Ptr(renderer.Indices), renderer.DrawType)

	return ebo
}
