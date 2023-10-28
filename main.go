package main

import (
	"fmt"
	"log"

	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth  = 640
	windowHeight = 480
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window := createWindow()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	programLoop(window)
}

func createWindow() *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}

var vertices = []float32{
	-0.5, -0.5,
	0.5, -0.5,
	0.0, 0.5,
}

const vertexShaderSource = `
	#version 410 core

	layout(location = 0) in vec3 inPosition;

	out vec3 FragPos;

	void main()
	{
	    gl_Position = vec4(inPosition, 1.0);
	    FragPos = inPosition;
	}
`

const fragmentShaderSource = `
	#version 410 core

	out vec4 FragColor;
	in vec3 FragPos;

	void main()
	{
	    // Calculate the color based on the fragment's position
	    vec3 color = vec3(abs(FragPos.x), abs(FragPos.y), abs(FragPos.z));
    
	    // Output the color
	    FragColor = vec4(color, 1.0);
	}
`

// createShaderProgram links vertex and fragment shaders into a shader program.
func createShaderProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {

	// Compile the vertex shader
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	// Compile the fragment shader
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		gl.DeleteShader(vertexShader)
		return 0, err
	}

	// Create a shader program and link the shaders
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	// Check for linking errors
	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, &log[0])
		return 0, fmt.Errorf("failed to link shaders: %v", string(log))
	}

	// Shaders are no longer needed after linking
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return shaderProgram, nil
}

// compileShader compiles a single shader and returns its ID.
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csource, nil)
	free()
	gl.CompileShader(shader)

	// Check for compilation errors
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &log[0])
		return 0, fmt.Errorf("failed to compile shader: %v", string(log))
	}

	return shader, nil
}

// createVAOn creates a Vertex Array Object and returns its ID.
func createVAO(vertices []float32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Define the vertex attributes
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 2*4, 0)
	gl.EnableVertexAttribArray(0)

	// Unbind the VBO and VAO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return vao
}

func programLoop(window *glfw.Window) {
	shaderProgram, err := createShaderProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		panic(err)
	}
	vao := createVAO(vertices)

	for !window.ShouldClose() {
		// Render the triangle using the VAO
		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/2))
		gl.BindVertexArray(0)

		// Swap buffers, poll events, etc.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
