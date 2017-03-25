package gl33

import (
	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func init() {
	engine.RegisterGraphics(glengine{})
}

type glengine struct{}

func (glengine) GLInit() error {
	return gl.Init()
}

func (glengine) GLUninit() error {
	return nil
}
