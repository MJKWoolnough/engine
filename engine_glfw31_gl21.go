package main

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()
}

var keyMap = map[glfw.Key]key{
	glfw.KeyEscape: KeyEscape,
	glfw.KeyUp:     KeyUp,
	glfw.KeyDown:   KeyDown,
	glfw.KeyLeft:   KeyLeft,
	glfw.KeyRight:  KeyRight,
}

func loop(c Config) error {
	if err := glfw.Init(); err != nil {
		return err
	}
	defer glfw.Terminate()
	window, err := glfw.CreateWindow(c.Width, c.Height, c.Title, nil, nil)
	if err != nil {
		return err
	}
	defer window.Destroy()

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		eKey, ok := keyMap[key]
		if action == glfw.Press {
			Keys.Down(eKey)
		} else if action == glfw.Release {
			Keys.Up(eKey)
		}
	})

	gl.Init()

	for !window.ShouldClose() {
		width, height := window.GetSize()
		ratio := float64(width) / float64(height)

		gl.Viewport(0, 0, int32(width), int32(height))
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Ortho(-ratio, ratio, -1, 1, 1, -1)
		gl.MatrixMode(gl.MODELVIEW)

		gl.LoadIdentity()
		gl.Rotated(glfw.GetTime()*50, 0, 0, 1)

		gl.Begin(gl.TRIANGLES)
		gl.Color3f(1, 0, 0)
		gl.Vertex3f(-0.6, -0.4, 0.)
		gl.Color3f(0, 1, 0)
		gl.Vertex3f(0.6, -0.4, 0)
		gl.Color3f(0, 0, 1)
		gl.Vertex3f(0, 0.6, 0)
		gl.End()

		window.SwapBuffers()
		glfw.PollEvents()
	}
	return nil
}
