package main

import (
	"log"
	"time"

	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"schneemelcher.com/gopengl/internal/input"
	"schneemelcher.com/gopengl/internal/settings"
	"schneemelcher.com/gopengl/internal/shaders"
	"schneemelcher.com/gopengl/internal/utils"
)

var (
	triangleX, triangleY int
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window := utils.CreateWindow(settings.WindowWidth, settings.WindowHeight)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	programLoop(window)
}

const targetFps = 60

func programLoop(window *glfw.Window) {
	shaderProgram, err := utils.CreateShaderProgram(shaders.VertexShaderSource, shaders.FragmentShaderSource)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(input.KeyCallback)

	previousTime := glfw.GetTime()
	for !window.ShouldClose() {
		currentTime := glfw.GetTime()
		deltaTime := currentTime - previousTime
		previousTime = currentTime

		log.Println(1 / deltaTime)

		if input.UpPressed {
			triangleY -= settings.MoveSpeed
		}
		if input.DownPressed {
			triangleY += settings.MoveSpeed
		}
		if input.LeftPressed {
			triangleX -= settings.MoveSpeed
		}
		if input.RightPressed {
			triangleX += settings.MoveSpeed
		}

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Render the triangle using the VAO
		gl.UseProgram(shaderProgram)
		utils.DrawSquare(triangleX, triangleY, 20, 400)

		// Swap buffers, poll events, etc.
		window.SwapBuffers()
		glfw.PollEvents()

		targetFrameTime := 1.0 / targetFps
		sleepTime := targetFrameTime - deltaTime

		if sleepTime > 0 {
			time.Sleep(time.Duration(sleepTime * float64(time.Second)))
		}
	}
}
