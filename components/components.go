package components

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl64"
)

// Application data
type App struct {
	width      int
	height     int
	title      string
	fullscreen bool
	shaders    []shaderInfo // List of shaders for this program
}

// Shader data
type shaderInfo struct {
	stype  uint32 // OpenGL shader type (VERTEX_SHADER, FRAGMENT_SHADER, or GEOMETRY_SHADER)
	source string // Shader source code
}

// Global object data
type Entity struct {
	Name      string
	Renderer  Renderer
	Transform Transform
}

// Scene data
type Scene struct {
	isLoaded bool
	entities []Entity
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
	DrawType uint32 // GL_STATIC_DRAW | GL_DYNAMIC_DRAW | GL_STREAM_DRAW
	Mode     uint32 //GL_TRIANGLES, etc.
	Xtype    uint32 // gl.UNSIGNED_INT, etc.
	Vertices []float32
	Indices  []uint32
}

func (renderer *Renderer) GetVertices() []float32 {
	return renderer.Vertices
}

func (renderer *Renderer) GetDrawType() uint32 {
	return renderer.DrawType
}

func (renderer *Renderer) Draw() {

	if len(renderer.Indices) > 1 {
		gl.DrawElements(renderer.Mode, int32(len(renderer.Indices)), renderer.Xtype, nil)
	} else {
		gl.DrawArrays(renderer.Mode, 0, int32(len(renderer.Vertices)/3))
	}

	gl.BindVertexArray(0) //TODO No need un-bind every time

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
