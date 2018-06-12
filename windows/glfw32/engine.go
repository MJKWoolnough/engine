package glfw32 // import "vimagination.zapto.org/engine/windows/glfw32"

import (
	"runtime"
	"sort"

	"github.com/go-gl/glfw/v3.2/glfw"
	"vimagination.zapto.org/engine"
)

type glfwengine struct {
	window *glfw.Window
}

func init() {
	runtime.LockOSThread()
	g := &glfwengine{}
	engine.RegisterWindow(g)
	engine.RegisterInput(g)
	if err := glfw.Init(); err != nil {
		panic(err)
	}
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

func (g *glfwengine) WindowInit(c engine.Config) error {
	var monitor *glfw.Monitor
	if c.Monitor != nil {
		if m, ok := c.Monitor.Data().(*glfw.Monitor); ok {
			monitor = m
		}
	}
	window, err := glfw.CreateWindow(int(c.Width), int(c.Height), c.Title, monitor, nil)
	if err != nil {
		return err
	}
	g.window = window

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	return nil
}

func (g *glfwengine) Loop(run func(int, int, float64) bool) {
	for !g.window.ShouldClose() {
		width, height := g.window.GetSize()
		if run(width, height, glfw.GetTime()) {
			g.window.SwapBuffers()
			glfw.PollEvents()
		} else {
			glfw.WaitEventsTimeout(0.125)
		}
	}
}

func (g *glfwengine) WindowUninit() error {
	g.window.Destroy()
	g.window = nil
	glfw.Terminate()
	return nil
}

func (g *glfwengine) Close() {
	g.window.SetShouldClose(true)
}

func (g *glfwengine) InputInit() error {
	return nil
}

func (g *glfwengine) InputUninit() error {
	return nil
}

func (g *glfwengine) Poll() {}

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

type modes []*glfw.VidMode

func (m modes) Len() int {
	return len(m)
}

func (m modes) Less(i, j int) bool {
	if m[i].Width < m[j].Width {
		return true
	} else if m[i].Width > m[j].Width {
		return false
	}
	if m[i].Height < m[j].Height {
		return true
	} else if m[i].Height > m[j].Height {
		return false
	}
	if m[i].RefreshRate < m[j].RefreshRate {
		return true
	} else if m[i].RefreshRate > m[j].RefreshRate {
		return false
	}
	if m[i].RedBits < m[j].RedBits {
		return true
	} else if m[i].RedBits > m[j].RedBits {
		return false
	}
	if m[i].GreenBits < m[j].GreenBits {
		return true
	} else if m[i].GreenBits > m[j].GreenBits {
		return false
	}
	return m[i].BlueBits < m[j].BlueBits
}

func (m modes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (g *glfwengine) GetMonitors() []*engine.Monitor {
	pm := glfw.GetPrimaryMonitor()
	monitors := glfw.GetMonitors()
	em := make([]*engine.Monitor, 1, len(monitors))
	pmn := pm.GetName()
	em[0] = engine.NewMonitor(pmn, pm)
	for _, m := range monitors {
		if m.GetName() != pmn {
			em = append(em, engine.NewMonitor(m.GetName(), m))
		}
	}
	return em
}

func (g *glfwengine) GetModes(m interface{}) []engine.Mode {
	monitor, ok := m.(*glfw.Monitor)
	if !ok {
		return nil
	}
	vm := monitor.GetVideoModes()
	sort.Sort(modes(vm))
	modes := make([]engine.Mode, len(vm))
	var lastWidth, lastHeight, lastRefresh, i int
	for _, mode := range vm {
		modes[i].Width = mode.Width
		modes[i].Height = mode.Height
		modes[i].Refresh = mode.RefreshRate
		if mode.Width != lastWidth || mode.Height != lastHeight || mode.RefreshRate != lastRefresh {
			lastWidth = mode.Width
			lastHeight = mode.Height
			lastRefresh = mode.RefreshRate
			i++
		}
	}
	return modes[:i:i]
}

func (g *glfwengine) SetMode(m interface{}, mode engine.Mode) {
	var monitor *glfw.Monitor
	if m != nil {
		var ok bool
		monitor, ok = m.(*glfw.Monitor)
		if !ok {
			return
		}
	}
	g.window.SetMonitor(monitor, 0, 0, mode.Width, mode.Height, mode.Refresh)
}
