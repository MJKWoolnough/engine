package gl21

import (
	"unsafe"

	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/gl/v2.1/gl"
)

var jitter = [...][2]float32{
	{-1.0 / 12, -5.0 / 12},
	{1.0 / 12, 1.0 / 12},
	{3.0 / 12, -1.0 / 12},
	{5.0 / 12, 5.0 / 12},
	{7.0 / 12, -3.0 / 12},
	{9.0 / 12, 3.0 / 12},
}

type font struct {
	engine       *glengine
	frameBuffer  uint32
	vertexBuffer uint32
	first        rune
	points       [][2]int
	advances     []float32
	areas        [][4]float32
}

func (f *font) Render(x1, y1, x2, y2 float64, text string) {
	var longestLine, currLine float32
	lines := 1
	for _, g := range text {
		if g == '\n' {
			if currLine > longestLine {
				longestLine = currLine
			}
			lines++
			currLine = 0
			continue
		}
		i := int(g - f.first)
		if i < 0 || i >= len(f.points) {
			continue
		}
		currLine += f.advances[i]
	}
	if currLine > longestLine {
		longestLine = currLine
	}
	var (
		transform engine.Transform2D
		emH, emV  float32
		advance   float32
	)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.frameBuffer)
	gl.BlendFunc(gl.ONE, gl.ONE)
	gl.BlendEquation(gl.FUNC_SUBTRACT)
	f.engine.shader.Use()
	transform.Scale(emH, emv)
	for _, g := range text {
		if g == '\n' {
			transform.Translate(-advance, emV)
			continue
		}
		i := int(g - f.first)
		if i < 0 || i >= len(f.points) {
			continue
		}

		gl.BindBuffer(gl.ARRAY_BUFFER, f.engine.pos)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, f.engine.pos)
		gl.VertexAttribPointer(uint32(pos), 4, gl.FLOAT, false, 0, unsafe.Pointer(uintptr(f.points[i][0])))
		gl.EnableVertexAttribArray(uint32(pos))

		nPoints := int32(f.points[i][1]) / 4

		for n, j := range jitter {
			trans := transform
			trans.Translate(j[0], j[1])
			switch n {
			case 0:
				gl.Uniform4f(f.engine.colour, 1, 0, 0, 1)
			case 2:
				gl.Uniform4f(f.engine.colour, 0, 1, 0, 1)
			case 4:
				gl.Uniform4f(f.engine.colour, 0, 0, 1, 1)
			}
			gl.UniformMatrix3fv(f.engine.transform, 1, true, &trans[0])
			gl.DrawArrays(gl.TRIANGLES, 0, nPoints)
		}

		gl.DisableVertexAttribArray(uint32(pos))

		advance += f.advances[i]
		transform.Translate(f.advances[i], 0)
	}
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.BlendFunc(gl.ZERO, gl.SRC_COLOR)

}

func (f *font) Length(text string) float32 {
	var length float32
	for _, g := range text {
		i := int(g - f.first)
		if i < 0 || i >= len(f.points) {
			continue
		}
		length += f.advances[i]
	}
	return length
}
