package sdl

import (
	"github.com/MJKWoolnough/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type sdlengine struct {
	window  *sdl.Window
	context sdl.GLContext
}

func init() {
	s := new(sdlengine)
	engine.RegisterWindow(s)
}

func (s *sdlengine) Init(c engine.Config) error {
	sdl.Init(sdl.INIT_VIDEO)
	sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	sdl.GL_SetAttribute(sdl.GL_DEPTH_SIZE, 24)
	window, err := sdl.CreateWindow(c.Title, 0, 0, c.Mode.Width, c.Mode.Height, sdl.WINDOW_OPENGL|sdl.WINDOW_FULLSCREEN)
	if err != nil {
		return err
	}
	context, err := sdl.GL_CreateContext(window)
	return engine.GLInit()
}

func (s *sdlengine) Loop(run func(int, int, float64)) {

}

func (s *sdlengine) Uninit() error {
	return engine.GLUninit()
}

func (s *sdlengine) GetMonitors() []*engine.Monitor {
	return nil
}

func (s *sdlengine) GetModes(m interface{}) []engine.Mode {
	return nil
}

func (s *sdlengine) SetMode(m interface{}, mode engine.Mode) {

}

func (s *sdlengine) Close() {

}
