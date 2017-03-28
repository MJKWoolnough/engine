package engine

import "io"

type none struct{}

func (none) Loop(func(int, int, float64)) {}

func (none) GetMonitors() []*Monitor {
	return nil
}

func (none) GetModes(interface{}) []Mode {
	return nil
}

func (none) SetMode(interface{}, Mode) {}

func (none) Close() {}

func (none) WindowInit(Config) error {
	return nil
}

func (none) Uninit() error {
	return nil
}

func (none) GLInit() error {
	return nil
}

func (none) GLUninit() error {
	return nil
}

func (none) ID() string {
	return "NONE"
}

func (none) Poll() {}

func (none) KeyPressed(Key) bool {
	return false
}

func (none) CursorPos() (float64, float64) {
	return 0, 0
}

func (none) Play(Sound) {}

func (none) LoadFont(io.Reader) Font {
	return nil
}
