package gl21

import "github.com/MJKWoolnough/engine"

func init() {

}

type glengine struct {
}

func (g *glengine) Init() error {
}

func (g *glengine) Uninit() error {
	return nil
}

func (g *glengine) LoadFont(glyphs, advances []float32, points []int, areas [][4]float32, first rune) (engine.Font, error) {
}
