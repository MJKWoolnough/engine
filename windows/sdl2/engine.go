package sdl2 // import "vimagination.zapto.org/engine/windows/sdl2"

import (
	"runtime"
	"sort"
	"sync/atomic"

	"github.com/veandco/go-sdl2/sdl"
	"vimagination.zapto.org/engine"
	_ "vimagination.zapto.org/engine/input/sdl2"
)

type sdlengine struct {
	window   *sdl.Window
	context  *sdl.GLContext
	renderer *sdl.Renderer
	quit     uint32
}

func init() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER); err != nil {
		panic(err)
	}
	s := new(sdlengine)
	engine.RegisterWindow(s)
}

func (s *sdlengine) WindowInit(c engine.Config) error {
	graphicsID := engine.GLID()
	var (
		flags uint32
		gl    bool
		x     int = sdl.WINDOWPOS_UNDEFINED
		y     int = sdl.WINDOWPOS_UNDEFINED
	)
	if len(graphicsID) >= 2 && graphicsID[0] == 'G' && graphicsID[1] == 'L' {
		sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)
		sdl.GL_SetAttribute(sdl.GL_RED_SIZE, 8)
		sdl.GL_SetAttribute(sdl.GL_BLUE_SIZE, 8)
		sdl.GL_SetAttribute(sdl.GL_GREEN_SIZE, 8)
		sdl.GL_SetAttribute(sdl.GL_ALPHA_SIZE, 8)
		flags |= sdl.WINDOW_OPENGL
		gl = true
	}
	if c.Monitor != nil {
		flags |= sdl.WINDOW_FULLSCREEN
		m, _ := c.Monitor.Data().(monitor)
		var r sdl.Rect
		if err := sdl.GetDisplayBounds(int(m), &r); err != nil {
			return err
		}
		x = int(r.X)
		y = int(r.Y)
	}
	window, err := sdl.CreateWindow(c.Title, x, y, c.Mode.Width, c.Mode.Height, flags)

	if err != nil {
		return err
	}
	if gl {
		context, err := sdl.GL_CreateContext(window)
		if err != nil {
			return err
		}
		s.context = &context
		if err = sdl.GL_SetSwapInterval(1); err != nil {
			return err
		}
	} else {
		renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
		if err != nil {
			return err
		}
		s.renderer = renderer
	}
	s.window = window
	return nil
}

func (s *sdlengine) Loop(run func(int, int, float64) bool) {
	for atomic.LoadUint32(&s.quit) == 0 {
		engine.PollInput()
		w, h := s.window.GetSize()
		t := float64(sdl.GetTicks()) / 1000
		if run(w, h, t) {
			if s.context != nil {
				sdl.GL_SwapWindow(s.window)
			} else {
				s.renderer.Present()
			}
		}
	}
}

func (s *sdlengine) Close() {
	atomic.StoreUint32(&s.quit, 1)
}

func (s *sdlengine) WindowUninit() error {
	if s.context != nil {
		sdl.GL_DeleteContext(*s.context)
	}
	s.window.Destroy()
	sdl.Quit()
	return nil
}

type monitor int

func (s *sdlengine) GetMonitors() []*engine.Monitor {
	n, err := sdl.GetNumVideoDisplays()
	if err != nil {
		return nil
	}
	monitors := make([]*engine.Monitor, n)
	for i := 0; i < n; i++ {
		monitors[i] = engine.NewMonitor(sdl.GetDisplayName(i), monitor(i))
	}
	return monitors
}

type modes []engine.Mode

func (m modes) Len() int {
	return len(m)
}

func (m modes) Less(i, j int) bool {
	if m[i].Width < m[j].Width {
		return true
	} else if m[j].Width < m[i].Width {
		return false
	}
	if m[i].Height < m[j].Height {
		return true
	} else if m[j].Height < m[i].Height {
		return false
	}
	return m[i].Refresh < m[j].Refresh
}

func (m modes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (s *sdlengine) GetModes(m interface{}) []engine.Mode {
	monitor, _ := m.(monitor)
	n, err := sdl.GetNumDisplayModes(int(monitor))
	if err != nil {
		return nil
	}
	ms := make([]engine.Mode, n)
	var mode sdl.DisplayMode
	for i := 0; i < n; i++ {
		if err := sdl.GetDisplayMode(int(monitor), i, &mode); err != nil {
			return nil
		}
		ms[i].Width = int(mode.W)
		ms[i].Height = int(mode.H)
		ms[i].Refresh = int(mode.RefreshRate)
	}
	sort.Sort(modes(ms))
	return ms
}

func (s *sdlengine) SetMode(m interface{}, mode engine.Mode) {
	mon, ok := m.(monitor)
	if ok {
		_ = mon
	}
	//s.window.SetDisplayMode
}
