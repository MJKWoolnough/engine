package gl33 // import "vimagination.zapto.org/engine/graphics/gl33"

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"vimagination.zapto.org/engine"
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
	return "GLv3.3"
}
