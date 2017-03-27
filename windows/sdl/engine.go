package sdl

import (
	"runtime"
	"sort"
	"sync/atomic"

	"github.com/MJKWoolnough/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type sdlengine struct {
	window  *sdl.Window
	context *sdl.GLContext
	quit    uint32
	keys    [256]bool
}

func init() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER); err != nil {
		panic(err)
	}
	s := new(sdlengine)
	engine.RegisterWindow(s)
	engine.RegisterInput(s)
}

func (s *sdlengine) Init(c engine.Config) error {
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
	}
	s.window = window
	return engine.GLInit()
}

func (s *sdlengine) Loop(run func(int, int, float64)) {
	for atomic.LoadUint32(&s.quit) == 0 {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e := e.(type) {
			case *sdl.KeyDownEvent:
				s.setKey(e.Keysym, true)
			case *sdl.KeyUpEvent:
				s.setKey(e.Keysym, false)
			case *sdl.MouseButtonEvent:
				s.setMouse(e.Button, e.State == sdl.PRESSED)
				//case sdl.MouseMotionEvent:
			}
		}
		w, h := s.window.GetSize()
		t := float64(sdl.GetTicks()) / 1000
		run(w, h, t)
		if s.context != nil {
			sdl.GL_SwapWindow(s.window)
		}
		sdl.Delay(10)
	}
}

func (s *sdlengine) Close() {
	atomic.StoreUint32(&s.quit, 1)
}

func (s *sdlengine) Uninit() error {
	err := engine.GLUninit()
	if err != nil {
		return err
	}
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

func (s *sdlengine) KeyPressed(k engine.Key) bool {
	return s.keys[k]
}

func (s *sdlengine) CursorPos() (float64, float64) {
	x, y, _ := sdl.GetMouseState()
	return float64(x), float64(y)
}

func (s *sdlengine) setKey(k sdl.Keysym, down bool) {
	switch k.Sym {
	case sdl.K_ESCAPE:
		s.keys[engine.KeyEscape] = down
	}
}

func (s *sdlengine) setMouse(b uint8, down bool) {
	switch b {
	case sdl.BUTTON_LEFT:
	case sdl.BUTTON_RIGHT:
	case sdl.BUTTON_MIDDLE:
	case sdl.BUTTON_X1:
	case sdl.BUTTON_X2:
	}
}
