package engine

const noEngine = "no engine registered"

type graphics interface {
	Loop(Config, func(int, int, float64) bool) error
}

type audio interface {
}

type input interface {
	KeyPressed(Key) bool
	CursorPos() (float64, float64)
}

type none struct{}

func (none) Loop(Config, func(int, int, float64) bool) error {
	panic(noEngine)
}

func (none) KeyPressed(Key) bool {
	panic(noEngine)
}

func (none) CursorPos() (float64, float64) {
	panic(noEngine)
}

var (
	registeredGraphics graphics = none{}
	registeredAudio    audio    = none{}
	registeredInput    input    = none{}
)

func RegisterGraphics(g graphics) {
	switch registeredGraphics.(type) {
	case none:
		registeredGraphics = g
	default:
		panic("cannot register multiple graphics engines")
	}
}

func RegisterAudio(a audio) {
	switch registeredAudio.(type) {
	case none:
		registeredAudio = a
	default:
		panic("cannot register multiple audio engines")
	}
}

func RegisterInput(i input) {
	switch registeredInput.(type) {
	case none:
		registeredInput = i
	default:
		panic("cannot register multiple input engines")
	}
}

func Loop(c Config, run func(int, int, float64) bool) error {
	return registeredGraphics.Loop(c, run)
}

func KeyPressed(k Key) bool {
	return registeredInput.KeyPressed(k)
}

func CursorPos() (float64, float64) {
	return registeredInput.CursorPos()
}
