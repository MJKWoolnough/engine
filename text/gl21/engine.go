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
	shader                 *graphics.Program
	transform, colour, pos int32
}

func (g *glengine) TextInit() error {
	var err error
	g.shader, err = graphics.NewProgram(glyphVertexShader, glyphFragmentShader)
	if err != nil {
		return err
	}
	g.transform, err = g.shader.GetUniformLocation("transform")
	if err != nil {
		return err
	}
	g.colour, err = g.shader.GetUniformLocation("colour")
	if err != nil {
		return err
	}
	g.pos, err = g.shader.GetAttribLocation("pos")
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

	// do frame buffer stuff
	return &font{
		engine:       g,
		vertexBuffer: vb,
		first:        ' ',
		advances:     t.Advances,
		points:       t.Pos,
	}, nil
}
