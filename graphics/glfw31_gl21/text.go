package glfw31_gl21

import (
	"io"
	"io/ioutil"

	"golang.org/x/image/math/fixed"

	"github.com/MJKWoolnough/engine"
	"github.com/golang/freetype/truetype"
)

type font struct {
	min, max rune
	points   [][]truetype.Point
}

func (g *glengine) LoadFont(r io.Reader, min, max rune) (engine.Font, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}
	g := new(truetype.GlyphBuf)
	p := make([][]truetype.Point, max-min+1)
	for r := min; r <= max; r++ {
		err = g.Load(f, fixed.I(100), font.Index(r), 0)
		if err != nil {
			return nil, err
		}
		p[r-min] = g.Points
	}
	return font{min, max, p}, nil
}
