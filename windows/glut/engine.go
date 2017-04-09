package glut

import (
	"runtime"

	"github.com/MJKWoolnough/engine"
	"github.com/vbsw/glut"
)

type glutengine struct {
	window        int
	width, height int
	runFunc       func(int, int, float64) bool
	keys          [256]bool
}

func init() {
	runtime.LockOSThread()
	g := new(glutengine)
	engine.RegisterWindow(g)
	engine.RegisterInput(g)
	glut.Init()
}

func (g *glutengine) WindowInit(c engine.Config) error {
	g.width = c.Width
	g.height = c.Height
	glut.InitDisplayMode(glut.DOUBLE | glut.RGBA)
	glut.InitWindowPosition(0, 0)
	glut.InitWindowSize(c.Width, c.Height)
	g.window = glut.CreateWindow(c.Title)
	if c.Monitor != nil {
		glut.FullScreen()
	}
	glut.IdleFunc(glut.PostRedisplay)
	glut.DisplayFunc(g.loop)
	glut.KeyboardFunc(g.keyboardDown)
	glut.KeyboardUpFunc(g.keyboardUp)
	glut.IgnoreKeyRepeat(1)
	glut.MouseFunc(g.mouse)
	glut.ReshapeFunc(g.reshape)
	return nil
}

func (g *glutengine) Loop(run func(int, int, float64) bool) {
	g.runFunc = run
	glut.MainLoop()
}

func (g *glutengine) Close() {
	glut.DestroyWindow(g.window)
	g.window = 0
}

func (g *glutengine) WindowUninit() error {
	return nil
}

func (g *glutengine) GetMonitors() []*engine.Monitor {
	return []*engine.Monitor{engine.NewMonitor("Only", 0)}
}

func (g *glutengine) GetModes(monitor interface{}) []engine.Mode {
	return []engine.Mode{
		{Width: glut.Get(glut.SCREEN_WIDTH), Height: glut.Get(glut.SCREEN_HEIGHT), Refresh: 60},
	}
}

func (g *glutengine) SetMode(monitor interface{}, mode engine.Mode) {

}

func (g *glutengine) Poll() {}

func (g *glutengine) InputInit() error {
	return nil
}

func (g *glutengine) InputUninit() error {
	return nil
}

func (g *glutengine) CursorPos() (x, y float64) {
	return 0, 0
}

func (g *glutengine) KeyPressed(key engine.Key) bool {
	return g.keys[key]
}

func (g *glutengine) loop() {
	t := float64(glut.Get(glut.ELAPSED_TIME)) / 1000
	g.runFunc(g.width, g.height, t)
	if g.window > 0 {
		glut.SwapBuffers()
	}
}

func (g *glutengine) reshape(width, height int) {
	g.width = width
	g.height = height
}

func (g *glutengine) keyboardDown(key uint8, x, y int) {
	if key == 27 {
		g.keys[engine.KeyEscape] = true
	}
}

func (g *glutengine) keyboardUp(key uint8, x, y int) {
	if key == 27 {
		g.keys[engine.KeyEscape] = false
	}
}

func (g *glutengine) mouse(button, state, x, y int) {

}
