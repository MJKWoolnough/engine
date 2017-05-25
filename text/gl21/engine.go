package gl21

import (
	"io"
	"unsafe"

	"github.com/MJKWoolnough/engine"
	graphics "github.com/MJKWoolnough/engine/graphics/gl21"
	"github.com/go-gl/gl/v2.1/gl"
)

func init() {
	engine.RegisterText(new(glengine))
}

type glengine struct {
	glyphs, text                                  *graphics.Program
	glyphTransform, glyphColour, glyphPos         int32
	textRect, textColour, textCoords, textTexture int32
}

func (g *glengine) TextInit() error {
	var err error
	g.glyphs, err = graphics.NewProgram(glyphVertexShader, glyphFragmentShader)
	if err != nil {
		return err
	}
	g.glyphTransform, err = g.glyphs.GetUniformLocation("transform")
	if err != nil {
		return err
	}
	g.glyphColour, err = g.glyphs.GetUniformLocation("colour")
	if err != nil {
		return err
	}
	g.glyphPos, err = g.glyphs.GetAttribLocation("pos")
	if err != nil {
		return err
	}
	g.text, err = graphics.NewProgram(textVertexShader, textFragmentShader)
	if err != nil {
		return err
	}
	g.textRect, err = g.text.GetUniformLocation("rect")
	if err != nil {
		return err
	}
	g.textColour, err = g.text.GetUniformLocation("colour")
	if err != nil {
		return err
	}
	g.textTexture, err = g.text.GetUniformLocation("texture")
	if err != nil {
		return err
	}
	g.textCoords, err = g.text.GetAttribLocation("coords")
	if err != nil {
		return err
	}
	return nil
}

func (g *glengine) TextUninit() error {
	return nil
}

func (g *glengine) LoadFont(r io.Reader) (engine.Font, error) {
	t, err := engine.DecodeTTF(r, ' ', '~')
	if err != nil {
		return nil, err
	}

	var vb uint32
	gl.GenBuffers(1, &vb)
	gl.BindBuffer(gl.ARRAY_BUFFER, vb)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vb)
	gl.BufferData(gl.ARRAY_BUFFER, len(t.Coords)*int(unsafe.Sizeof(t.Coords[0])), unsafe.Pointer(&t.Coords[0]), gl.STATIC_DRAW)

	var fb uint32
	gl.GenFramebuffers(1, &fb)
	gl.BindFramebuffer(gl.FRAMEBUFFER, fb)

	// do frame buffer stuff
	return &font{
		engine:       g,
		frameBuffer:  fb,
		vertexBuffer: vb,
		first:        ' ',
		advances:     t.Advances,
		points:       t.Pos,
	}, nil
}
