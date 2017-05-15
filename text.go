package engine

import (
	"io"
	"io/ioutil"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type TTF struct {
	Advances []float32
	Boxes    [][4]float32
	Pos      [][2]int
	Coords   []float32
}

const em = 8192 // 128 << 6 // for fixed.Int26_6

func DecodeTTF(r io.Reader, start, end rune) (*TTF, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}

	var (
		g truetype.GlyphBuf
		v VertexBuilder
	)
	pos := make([][2]int, 0, end-start+1)
	advances := make([]float32, 0, end-start+1)
	boxes := make([][4]float32, 0, end-start+1)

	for r := start; r <= end; r++ {
		err = g.Load(f, em, f.Index(r), font.HintingNone)
		if err != nil {
			return nil, err
		}
		advances = append(advances, float32(g.AdvanceWidth)/em)
		boxes = append(boxes, [4]float32{
			float32(g.Bounds.Min.X) / em,
			float32(g.Bounds.Min.Y) / em,
			float32(g.Bounds.Max.X) / em,
			float32(g.Bounds.Max.Y) / em,
		})
		firstPos := len(v.Vertices)
		prevPos := 0
		for _, e := range g.Ends {
			lastPoint := g.Points[prevPos]
			v.MoveTo(lastPoint.X, lastPoint.Y)
			for _, point := range g.Points[prevPos+1 : e] {
				if point.Flags&1 == 0 && lastPoint.Flags&1 == 0 {
					v.CurveTo(lastPoint.X, lastPoint.Y, (point.X+lastPoint.X)/2, (point.Y+lastPoint.Y)/2)
					lastPoint.X = (point.X + lastPoint.X) / 2
					lastPoint.Y = (point.Y + lastPoint.Y) / 2
				}
				if point.Flags&1 == 1 {
					if lastPoint.Flags&1 == 1 {
						v.LineTo(point.X, point.Y)
					} else {
						v.CurveTo(lastPoint.X, lastPoint.Y, point.X, point.Y)
					}
				}
				lastPoint = point
			}
			prevPos += e
		}
		pos = append(pos, [2]int{firstPos, len(v.Vertices) - firstPos})
	}

	return &TTF{
		Advances: advances,
		Boxes:    boxes,
		Pos:      pos,
		Coords:   v.Vertices,
	}, nil
}

type VertexBuilder struct {
	start, current [2]fixed.Int26_6
	count          int
	Vertices       []float32
}

func (v *VertexBuilder) MoveTo(x, y fixed.Int26_6) {
	v.start = [2]fixed.Int26_6{x, y}
	v.current = v.start
	v.count = 0
}

func (v *VertexBuilder) LineTo(x, y fixed.Int26_6) {
	v.count++
	if v.count >= 2 {
		v.addTriangle(v.start[0], v.start[1], v.current[0], v.current[1], x, y)
	}
	v.current[0] = x
	v.current[1] = y
}

func (v *VertexBuilder) CurveTo(cx, cy, x, y fixed.Int26_6) {
	v.count++
	if v.count >= 2 {
		v.addTriangle(v.start[0], v.start[1], v.current[0], v.current[1], x, y)
	}
	v.addCurve(v.current[0], v.current[1], cx, cy, x, y)
	v.current[0] = x
	v.current[1] = y
}

func (v *VertexBuilder) addTriangle(ax, ay, bx, by, cx, cy fixed.Int26_6) {
	v.addVertex(ax, ay, 0, 1)
	v.addVertex(bx, by, 0, 1)
	v.addVertex(cx, cy, 0, 1)
}

func (v *VertexBuilder) addCurve(ax, ay, bx, by, cx, cy fixed.Int26_6) {
	v.addVertex(ax, ay, 0, 0)
	v.addVertex(bx, by, 0.5, 0)
	v.addVertex(cx, cy, 1, 1)
}

func (v *VertexBuilder) addVertex(x, y fixed.Int26_6, s, t float32) {
	v.Vertices = append(v.Vertices, float32(x)/em, float32(y)/em, s, t)
}
