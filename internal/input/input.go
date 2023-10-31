package input

import "github.com/go-gl/glfw/v3.3/glfw"

var UpPressed, DownPressed, LeftPressed, RightPressed bool

func KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	case glfw.KeyW:
		if action == glfw.Press {
			UpPressed = true
		} else if action == glfw.Release {
			UpPressed = false
		}
	case glfw.KeyS:
		if action == glfw.Press {
			DownPressed = true
		} else if action == glfw.Release {
			DownPressed = false
		}
	case glfw.KeyA:
		if action == glfw.Press {
			LeftPressed = true
		} else if action == glfw.Release {
			LeftPressed = false
		}
	case glfw.KeyD:
		if action == glfw.Press {
			RightPressed = true
		} else if action == glfw.Release {
			RightPressed = false
		}
	}
}
