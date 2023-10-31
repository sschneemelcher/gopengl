package utils

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"schneemelcher.com/gopengl/internal/settings"
)

func getTriangle(x, y, width, height int) uint32 {
	triangleX := float32(x)/settings.WindowWidth - 1
	triangleY := 1 - float32(y)/settings.WindowHeight
	widthFloat := 2 * (float32(width) / settings.WindowWidth)
	heightFloat := 2 * (float32(height) / settings.WindowHeight)

	var vertices = []float32{
		triangleX, triangleY - heightFloat,
		triangleX + widthFloat, triangleY - heightFloat,
		triangleX + widthFloat/2, triangleY,
	}

	return CreateVAO(vertices)
}

func getLeftTriangle(x, y, width, height int) uint32 {
	triangleX := float32(x)/settings.WindowWidth - 1
	triangleY := 1 - float32(y)/settings.WindowHeight
	widthFloat := float32(width) / settings.WindowWidth
	heightFloat := float32(height) / settings.WindowHeight

	var vertices = []float32{
		triangleX, triangleY - heightFloat,
		triangleX + widthFloat, triangleY - heightFloat,
		triangleX, triangleY,
	}

	return CreateVAO(vertices)
}

func getRightTriangle(x, y, width, height int) uint32 {
	triangleX := float32(x)/settings.WindowWidth - 1
	triangleY := 1 - float32(y)/settings.WindowHeight
	widthFloat := float32(width) / settings.WindowWidth
	heightFloat := float32(height) / settings.WindowHeight

	var vertices = []float32{
		triangleX, triangleY,
		triangleX + widthFloat, triangleY - heightFloat,
		triangleX + widthFloat, triangleY,
	}

	return CreateVAO(vertices)
}

func DrawSquare(x, y, width, height int) {
	leftTriangle := getLeftTriangle(2*x, 2*y, 2*width, 2*height)
	rightTriangle := getRightTriangle(2*x, 2*y, 2*width, 2*height)

	gl.BindVertexArray(leftTriangle)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(rightTriangle)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(0)
}
