package engine

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
)
