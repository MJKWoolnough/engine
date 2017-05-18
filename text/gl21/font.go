package gl21

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
	vertexBuffer uint32
	first        rune
	points       []int
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
		i := g - f.first
		if i < 0 || i >= len(f.points) {
			continue
		}
		currLine += f.advances[i]
	}
	if currLine > longestLine {
		longestLine = currLine
	}
	f.engine.shader.Use()
}

func (f *font) Length(text string) float32 {
	var length float32
	for _, g := range text {
		i := g - f.first
		if i < 0 || i >= len(f.points) {
			continue
		}
		length += f.advances[i]
	}
	return length
}
