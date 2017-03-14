package glfw31

import (
	"runtime"

	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type glfwengine struct {
	window *glfw.Window
}

func init() {
	runtime.LockOSThread()
	g := &glfwengine{}
	engine.RegisterWindow(g)
	engine.RegisterInput(g)
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

func (g *glfwengine) Loop(c engine.Config, run func(int, int, float64) bool) (err error) {
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

	if err := engine.GLInit(); err != nil {
		return err
	}

	for !window.ShouldClose() {
		width, height := window.GetSize()
		if !run(width, height, glfw.GetTime()) {
			window.SetShouldClose(true)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
	return engine.GLUninit()
}

func (g *glfwengine) KeyPressed(k engine.Key) bool {
	if mk, ok := mouseMap[k]; ok {
		return g.window.GetMouseButton(mk) == glfw.Press
	} else if kk, ok := keyMap[k]; ok {
		return g.window.GetKey(kk) == glfw.Press
	}
	return false
}

func (g *glfwengine) CursorPos() (x, y float64) {
	return g.window.GetCursorPos()
}
