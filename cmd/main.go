package main

import (
	"log"

	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"schneemelcher.com/gopengl/internal/shaders"
	"schneemelcher.com/gopengl/internal/utils"
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

	window := utils.CreateWindow(windowWidth, windowHeight)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	programLoop(window)
}

var vertices = []float32{
	-0.5, -0.5,
	0.5, -0.5,
	0.0, 0.5,
}

func programLoop(window *glfw.Window) {
	shaderProgram, err := utils.CreateShaderProgram(shaders.VertexShaderSource, shaders.FragmentShaderSource)
	if err != nil {
		panic(err)
	}
	vao := utils.CreateVAO(vertices)

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
