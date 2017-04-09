package sdl1

import (
	"runtime"
	"sort"
	"sync/atomic"

	"github.com/MJKWoolnough/engine"
	_ "github.com/MJKWoolnough/engine/input/sdl1"
	"github.com/banthar/Go-SDL/sdl"
)

type sdlengine struct {
	surface *sdl.Surface
	quit    uint32
	gl      bool
}

func init() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER); err != 0 {
		panic(sdl.GetError())
	}
	s := new(sdlengine)
	engine.RegisterWindow(s)
}

func (s *sdlengine) WindowInit(c engine.Config) error {
	graphicsID := engine.GLID()
	var (
		flags uint32
	)
	if len(graphicsID) >= 2 && graphicsID[0] == 'G' && graphicsID[1] == 'L' {
		sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)
		sdl.GL_SetAttribute(sdl.GL_RED_SIZE, 8)
		sdl.GL_SetAttribute(sdl.GL_BLUE_SIZE, 8)
		sdl.GL_SetAttribute(sdl.GL_GREEN_SIZE, 8)
		sdl.GL_SetAttribute(sdl.GL_ALPHA_SIZE, 8)
		flags |= sdl.OPENGL
		s.gl = true
	}
	if c.Monitor != nil {
		flags |= sdl.FULLSCREEN | sdl.HWACCEL
	}

	s.surface = sdl.SetVideoMode(c.Mode.Width, c.Mode.Height, 32, flags)

	return nil
}

func (s *sdlengine) Loop(run func(int, int, float64) bool) {
	for atomic.LoadUint32(&s.quit) == 0 {
		engine.PollInput()
		vi := sdl.GetVideoInfo()
		t := float64(sdl.GetTicks()) / 1000
		if run(int(vi.Current_w), int(vi.Current_h), t) && s.gl {
			sdl.GL_SwapBuffers()
		}
	}
}

func (s *sdlengine) Close() {
	atomic.StoreUint32(&s.quit, 1)
}

func (s *sdlengine) WindowUninit() error {
	s.surface.Free()
	sdl.Quit()
	return nil
}

func (s *sdlengine) GetMonitors() []*engine.Monitor {
	return []*engine.Monitor{engine.NewMonitor("DEFAULT", 0)}
}

type modeList []engine.Mode

func (m modeList) Len() int {
	return len(m)
}

func (m modeList) Less(i, j int) bool {
	if m[i].Width < m[j].Width {
		return true
	} else if m[j].Width < m[i].Width {
		return false
	}
	return m[i].Height < m[j].Height
}

func (m modeList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (s *sdlengine) GetModes(m interface{}) []engine.Mode {
	ms := sdl.ListModes(nil, sdl.FULLSCREEN|sdl.HWACCEL)
	modes := make([]engine.Mode, len(ms))
	for n, mode := range ms {
		modes[n].Width = int(mode.W)
		modes[n].Height = int(mode.H)
		modes[n].Refresh = 60
	}
	sort.Sort(modeList(modes))
	return modes
}

func (s *sdlengine) SetMode(m interface{}, mode engine.Mode) {
	//s.window.SetDisplayMode
}
