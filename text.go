package engine

import (
	"io"
	"io/ioutil"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type TTF struct {
	Advances []fixed.Int26_6
	Boxes    []fixed.Rectangle26_6
	Pos      [][2]int
	Coords   []float32
}

func DecodeTTF(r io.Reader, start, end rune) (*TTF, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}

	var g truetype.GlyphBuf
	coords := make([]float32, 0, 1024)
	pos := make([][2]int, 0, end-start+1)
	advances := make([]fixed.Int26_6, 0, end-start+1)
	boxes := make([]fixed.Rectangle26_6, 0, end-start+1)

	for r := start; r <= end; r++ {
		err = g.Load(f, fixed.I(100), f.Index(r), font.HintingNone)
		if err != nil {
			return nil, err
		}
		advances = append(advances, g.AdvanceWidth)
		boxes = append(boxes, g.Bounds)

	}

	return &TTF{
		Advances: advances,
		Boxes:    boxes,
		Pos:      pos,
		Coords:   coords,
	}, nil
}
