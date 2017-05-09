package engine

import "io"

const noEngine = "no engine registered"

type Monitor struct {
	Name string
	data interface{}
}

func NewMonitor(name string, data interface{}) *Monitor {
	return &Monitor{
		Name: name,
		data: data,
	}
}

func (m *Monitor) GetModes() []Mode {
	return registeredWindow.GetModes(m.data)
}

func (m *Monitor) Data() interface{} {
	return m.data
}

type Mode struct {
	Width, Height int
	Refresh       int
}

type window interface {
	WindowInit(Config) error
	Loop(func(int, int, float64) bool)
	WindowUninit() error
	GetMonitors() []*Monitor
	GetModes(interface{}) []Mode
	SetMode(interface{}, Mode)
	Close()
}

type graphics interface {
	GLInit() error
	GLUninit() error
	ID() string
}

type Sound interface {
}

type audio interface {
	Play(Sound)
}

type Font interface {
	Render(x1, y1, x2, y2 float64, text string)
}

type input interface {
	Poll()
	KeyPressed(Key) bool
	CursorPos() (float64, float64)
	InputInit() error
	InputUninit() error
}

type text interface {
	TextInit() error
	TextUninit() error
	LoadFont(glyphs []float32, points []int) (Font, error)
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

func Loop(run func(int, int, float64) bool) {
	registeredWindow.Loop(run)
}

func Close() {
	registeredWindow.Close()
}

func Init(c Config) error {
	if err := registeredWindow.WindowInit(c); err != nil {
		return err
	}
	if err := registeredGraphics.GLInit(); err != nil {
		return err
	}
	if err := registeredInput.InputInit(); err != nil {
		return err
	}
	if err := registeredText.TextInit(); err != nil {
		return err
	}
	return nil
}

func Uninit() error {
	if err := registeredText.TextUninit(); err != nil {
		return err
	}
	if err := registeredInput.InputUninit(); err != nil {
		return err
	}
	if err := registeredGraphics.GLUninit(); err != nil {
		return err
	}
	if err := registeredWindow.WindowUninit(); err != nil {
		return err
	}
	return nil
}

func GLID() string {
	return registeredGraphics.ID()
}

func GetMonitors() []*Monitor {
	return registeredWindow.GetMonitors()
}

func SetMode(m *Monitor, mode Mode) {
	var data interface{}
	if m != nil {
		data = m.data
	}
	registeredWindow.SetMode(data, mode)
}

func PollInput() {
	registeredInput.Poll()
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
