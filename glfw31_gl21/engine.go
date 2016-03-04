package glfw31_gl21

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

func loop(c Config, run func(int, int, float64) bool) error {
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
		if !ok {
			return
		}
		if action == glfw.Press {
			Keys.Down(eKey)
		} else if action == glfw.Release {
			Keys.Up(eKey)
		}
	})

	gl.Init()

	for !window.ShouldClose() {
		width, height := window.GetSize()
		if !run(width, height, glfw.GetTime()) {
			window.SetShouldClose(true)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
	return nil
}
