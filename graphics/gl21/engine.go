package gl21

import (
	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/gl/v2.1/gl"
)

func init() {
	engine.RegisterGraphics(glengine{})
}

type glengine struct{}

func (glengine) Init() error {
	return gl.Init()
}

func (glengine) Uninit() error {
	return nil
}
