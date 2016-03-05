package engine

const noEngine = "no engine registered"

type engine interface {
	Loop(Config, func(int, int, float64) bool) error
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

var registered engine = none{}

func Register(e engine) {
	switch registered.(type) {
	case none:
		registered = e
	default:
		panic("cannot register multiple engines")
	}
}
