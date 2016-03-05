package glfw31_gl21

import (
	"runtime"

	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type glengine struct {
	window *glfw.Window
}

func init() {
	runtime.LockOSThread()
	engine.Register(&glengine{})
}

var keyMap = map[engine.Key]glfw.Key{
	engine.KeyEscape: glfw.KeyEscape,
	engine.KeyUp:     glfw.KeyUp,
	engine.KeyDown:   glfw.KeyDown,
	engine.KeyLeft:   glfw.KeyLeft,
	engine.KeyRight:  glfw.KeyRight,
}

var mouseMap = map[engine.Key]glfw.MouseButton{
	engine.MouseLeft:   glfw.MouseButtonLeft,
	engine.MouseMiddle: glfw.MouseButtonMiddle,
	engine.MouseRight:  glfw.MouseButtonRight,
}

func (g *glengine) Loop(c engine.Config, run func(int, int, float64) bool) error {
	if err := glfw.Init(); err != nil {
		return err
	}
	defer glfw.Terminate()
	window, err := glfw.CreateWindow(int(c.Width), int(c.Height), c.Title, nil, nil)
	if err != nil {
		return err
	}
	defer window.Destroy()
	g.window = window
	defer func() {
		g.window = nil
	}()

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

	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		eKey, ok := mouseMap[button]
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
		mouseX, mouseY := window.GetCursorPos()
		Cursor.SetPos(int(mouseX), int(mouseY))
		width, height := window.GetSize()
		if !run(width, height, glfw.GetTime()) {
			window.SetShouldClose(true)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
	return nil
}

func (g *glengine) KeyPressed(k engine.Key) bool {
	if mk, ok := mouseMap[k]; ok {
		return g.window.GetMouseButton(mk) == glfw.Press
	} else if kk, ok := keyMap[k]; ok {
		return kk == glfw.Press
	}
	return false
}

func (g *glengine) CursorPos() (x, y float64) {
	return g.window.GetCursorPos()
}
