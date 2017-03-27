package sdl2

import (
	"runtime"

	"github.com/MJKWoolnough/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type sdlengine struct {
	keys [256]bool
}

const eventSubsystem = 0x4000

func init() {
	runtime.LockOSThread()
	if err := sdl.Init(eventSubsystem); err != nil {
		panic(err)
	}
	engine.RegisterInput(new(sdlengine))
}

func (s *sdlengine) Poll() {
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
}

func (s *sdlengine) Close() {
	sdl.QuitSubSystem(eventSubsystem)
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
