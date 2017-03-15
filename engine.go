package engine

import "io"

const noEngine = "no engine registered"

type window interface {
	Loop(Config, func(int, int, float64) bool) error
	Close()
}

type graphics interface {
	Init() error
	Uninit() error
}

type Sound interface {
}

type audio interface {
	Play(Sound)
}

type Font interface {
	Render(float64, float64, string)
}

type input interface {
	KeyPressed(Key) bool
	CursorPos() (float64, float64)
}

type text interface {
	LoadFont(io.Reader) Font
}

type none struct{}

func (none) Loop(Config, func(int, int, float64) bool) error {
	panic(noEngine)
}

func (none) Close() {
	panic(noEngine)
}

func (none) Init() error {
	panic(noEngine)
}

func (none) Uninit() error {
	panic(noEngine)
}

func (none) KeyPressed(Key) bool {
	panic(noEngine)
}

func (none) CursorPos() (float64, float64) {
	panic(noEngine)
}

func (none) Play(Sound) {
	panic(noEngine)
}

func (none) LoadFont(io.Reader) Font {
	panic(noEngine)
}

var (
	registeredGraphics graphics = none{}
	registeredWindow   window   = none{}
	registeredAudio    audio    = none{}
	registeredInput    input    = none{}
	registeredText     text     = none{}
)

func RegisterWindow(w window) {
	switch registeredWindow.(type) {
	case none:
		registeredWindow = w
	default:
		panic("cannot register multiple window engines")
	}
}

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

func RegisterText(t text) {
	switch registeredText.(type) {
	case none:
		registeredText = t
	default:
		panic("cannot register multiple text engines")
	}
}

func Loop(c Config, run func(int, int, float64) bool) error {
	return registeredWindow.Loop(c, run)
}

func Close() {
	registeredWindow.Close()
}

func GLInit() error {
	return registeredGraphics.Init()
}

func GLUninit() error {
	return registeredGraphics.Uninit()
}

func KeyPressed(k Key) bool {
	return registeredInput.KeyPressed(k)
}

func CursorPos() (float64, float64) {
	return registeredInput.CursorPos()
}

func PlaySound(s Sound) {
	registeredAudio.Play(s)
}

func LoadFont(r io.Reader) Font {
	return registeredText.LoadFont(r)
}
