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
	windowWidth  = 1200
	windowHeight = 800
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

func getTriangle(x, y, width, height int) uint32 {
	triangleX := float32(x)/windowWidth - 1
	triangleY := 1 - float32(y)/windowHeight
	widthFloat := 2 * (float32(width) / windowWidth)
	heightFloat := 2 * (float32(height) / windowHeight)

	var vertices = []float32{
		triangleX, triangleY - heightFloat,
		triangleX + widthFloat, triangleY - heightFloat,
		triangleX + widthFloat/2, triangleY,
	}

	return utils.CreateVAO(vertices)
}

func programLoop(window *glfw.Window) {
	shaderProgram, err := utils.CreateShaderProgram(shaders.VertexShaderSource, shaders.FragmentShaderSource)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(keyCallback)

	for !window.ShouldClose() {
		if upPressed {
			triangleY -= moveSpeed
		}
		if downPressed {
			triangleY += moveSpeed
		}
		if leftPressed {
			triangleX -= moveSpeed
		}
		if rightPressed {
			triangleX += moveSpeed
		}

		triangle := getTriangle(triangleX, triangleY, 300, 200)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Render the triangle using the VAO
		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(triangle)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.BindVertexArray(0)

		// Swap buffers, poll events, etc.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var (
	triangleX, triangleY                              int
	moveSpeed                                         int = 10
	upPressed, downPressed, leftPressed, rightPressed bool
)

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	case glfw.KeyW:
		if action == glfw.Press {
			upPressed = true
		} else if action == glfw.Release {
			upPressed = false
		}
	case glfw.KeyS:
		if action == glfw.Press {
			downPressed = true
		} else if action == glfw.Release {
			downPressed = false
		}
	case glfw.KeyA:
		if action == glfw.Press {
			leftPressed = true
		} else if action == glfw.Release {
			leftPressed = false
		}
	case glfw.KeyD:
		if action == glfw.Press {
			rightPressed = true
		} else if action == glfw.Release {
			rightPressed = false
		}
	}
}
