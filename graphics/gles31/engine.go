package gles2

import (
	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/gl/v3.1/gles2"
)

func init() {
	engine.RegisterGraphics(glengine{})
}

type glengine struct{}

func (glengine) Init() error {
	return gles2.Init()
}

func (glengine) Uninit() error {
	return nil
}
