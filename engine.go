package engine

const noEngine = "no engine registered"

type engine interface {
	Pressed(Key) bool
}

type none struct{}

func (none) Pressed(Key) bool {
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

type Key uint

const (
	KeyEscape Key = iota
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	MouseLeft
	MouseMiddle
	MouseRight
	MouseScrollUp
	MouseScrollDown
	NumKeys
)
