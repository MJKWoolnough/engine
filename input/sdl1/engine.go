package sdl1

import (
	"errors"
	"runtime"

	"github.com/MJKWoolnough/engine"
	"github.com/banthar/Go-SDL/sdl"
)

type sdlengine struct {
	keys [256]bool
}

const eventSubsystem = 0x4000

func init() {
	runtime.LockOSThread()
	if err := sdl.Init(eventSubsystem); err != 0 {
		panic(errors.New(sdl.GetError()))
	}
	engine.RegisterInput(new(sdlengine))
}

func (s *sdlengine) InputInit() error {
	return nil
}

func (s *sdlengine) InputUninit() error {
	sdl.QuitSubSystem(eventSubsystem)
	return nil
}

func (s *sdlengine) Poll() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e := e.(type) {
		case *sdl.KeyboardEvent:
			s.setKey(e.Keysym, e.State == 1)
		case *sdl.MouseButtonEvent:
			s.setMouse(e.Button, e.State == sdl.MOUSEBUTTONDOWN)
			//case sdl.MouseMotionEvent:
		}
	}
}

func (s *sdlengine) KeyPressed(k engine.Key) bool {
	return s.keys[k]
}

func (s *sdlengine) CursorPos() (float64, float64) {
	var x, y int
	sdl.GetMouseState(&x, &y)
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
