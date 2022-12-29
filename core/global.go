package core

import (
	"kotongo/components"
	"kotongo/graphics"
	"kotongo/shaders"
	"log"
	"runtime"

	"github.com/RedHelium/toolz"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

// TODO Transfer app struct in arg
func InitWindow(width, height int, title string, backgroundColor *toolz.Color) {

	runtime.LockOSThread()

	window := initGlfw(width, height, title)
	program := initOpenGL()

	//TODO Extract compile shaders from initOpenGL

	for !window.ShouldClose() {
		Draw(window, program, backgroundColor)
	}

	defer glfw.Terminate()
}

// TODO Add clear shaders function
// Initialize OpenGL
func initOpenGL() uint32 {

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	//TODO Create abstract in loop
	//Compile shaders
	vertex := compileShader(shaders.VertexShaderSource, gl.VERTEX_SHADER)
	fragment := compileShader(shaders.FragmentShaderSource, gl.FRAGMENT_SHADER)

	prog := gl.CreateProgram()

	attachShaders(prog, vertex, fragment)

	gl.LinkProgram(prog)
	return prog
}

func compileShader(source string, shaderType uint32) uint32 {

	shader, error := shaders.CompileShader(source, shaderType)
	if error != nil {
		panic(error)
	}

	return shader
}

// Attach all app shaders
func attachShaders(program uint32, shaders ...uint32) {

	for i := 0; i < len(shaders); i++ {
		gl.AttachShader(program, shaders[i])
	}
}

// Draw graphics
func Draw(window *glfw.Window, program uint32, backgroundColor *toolz.Color) {

	gl.ClearColor(backgroundColor.R, backgroundColor.G, backgroundColor.B, backgroundColor.A)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(program)

	//TODO THIS IS TEST!!! REFACT!!!
	r := components.Renderer{
		DrawType: gl.STATIC_DRAW,
		Mode:     gl.TRIANGLES,
		Xtype:    gl.UNSIGNED_INT,
		Vertices: graphics.RectanglePoints.Vertices,
		Indices:  graphics.RectanglePoints.Indices,
	}

	gl.BindVertexArray(graphics.VAO(r))

	//XXX This is change draw type geometry. For debuggin, we can use wireframe mode (gl.LINE)
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	//Callback all connected events
	glfw.PollEvents()
	//Update color buffer
	window.SwapBuffers()
}

// Initialize GLFW Window
func initGlfw(width, height int, title string) *glfw.Window {

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	//Set window state
	glfw.WindowHint(glfw.Resizable, glfw.False)
	//Set required OpenGL versions
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(setFrameBufferSize)

	return window
}

// Create GL Viewport
func setFrameBufferSize(window *glfw.Window, width, height int) {

	gl.Viewport(0, 0, int32(width), int32(height))
}
