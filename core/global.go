package core

import (
	"kotongo/core/extensions"
	"kotongo/shaders"
	"kotongo/shapes"
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func InitWindow(width, height int, title string, backgroundColor *extensions.Color) {

	runtime.LockOSThread()

	window := initGlfw(width, height, title)
	defer glfw.Terminate()

	program := initOpenGL(backgroundColor)

	//TODO Init started shapes
	sh := shapes.CreateSquare()
	//sh2 := shapes.Create(shapes.IsoscelesTrianglePoints)
	shapes := []*shapes.Shape{}
	shapes = append(shapes, sh)
	//shapes = append(shapes, sh2)

	for !window.ShouldClose() {
		Draw(window, program, shapes)
	}
}

func initGlfw(width, height int, title string) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL(backgroundColor *extensions.Color) uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.ClearColor(backgroundColor.R, backgroundColor.G, backgroundColor.B, backgroundColor.A)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	//TODO Create abstract in loop
	//Compile shaders
	vertexShader, err := shaders.CompileShader(shaders.VertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := shaders.CompileShader(shaders.FragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()

	attachShaders(prog, vertexShader, fragmentShader)

	gl.LinkProgram(prog)
	return prog
}

func attachShaders(program uint32, shaders ...uint32) {

	for i := 0; i < len(shaders); i++ {
		gl.AttachShader(program, shaders[i])
	}
}

// Draw shapes
func Draw(window *glfw.Window, program uint32, shapes []*shapes.Shape) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for i := 0; i < len(shapes); i++ {
		shapes[i].Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
