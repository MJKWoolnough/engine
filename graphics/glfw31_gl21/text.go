package glfw31_gl21

import (
	"io"

	"github.com/MJKWoolnough/engine"
)

type font struct {
}

func (g *glengine) LoadFont(r io.Reader) engine.Font {
	return font{}
}
