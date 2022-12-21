package components

import (
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

// Render object data
type Renderer struct {
	DrawType uint32
	Points   []float32
}

func (renderer *Renderer) GetPoints() []float32 {
	return renderer.Points
}

func (renderer *Renderer) GetDrawType() uint32 {
	return renderer.DrawType
}

// Create vertex buffer object
func (renderer *Renderer) VBO() uint32 {

	var vbo uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(renderer.Points), gl.Ptr(renderer.Points), renderer.DrawType)

	return vbo
}
