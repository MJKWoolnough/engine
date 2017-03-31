package gl21

import (
	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/gl/v2.1/gl"
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

func (glengine) ID() string {
	return "GLv2.1"
}
