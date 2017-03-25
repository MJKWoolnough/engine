// +build js

package webgl

import (
	"github.com/MJKWoolnough/engine"
	"github.com/MJKWoolnough/gopherjs/xdom"
	"github.com/MJKWoolnough/gopherjs/xjs"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
	"honnef.co/go/js/dom"
)

var (
	raf, uraf *js.Object
	Instance  webglengine
)

func init() {
	raf = js.Global.Get("requestAnimationFrame")
	uraf = js.Global.Get("cancelAnimationFrame")
	if raf != nil {
		engine.RegisterWindow(&Instance)
		engine.RegisterInput(&Instance)
	}
}

type webglengine struct {
	canvas         *dom.HTMLCanvasElement
	context        *webgl.Context
	fn             func(int, int, float64)
	rafID          int
	mouseX, mouseY float64
	keys           map[string]struct{}
}

func (w *webglengine) Init(c engine.Config) error {
	canvas := xdom.Canvas()
	canvas.Width = c.Width
	canvas.Height = c.Height
	canvas.SetTabIndex(1)
	ctx, err := webgl.NewContext(canvas.Underlying(), &webgl.ContextAttributes{})

	if err != nil {
		return err
	}

	if tDoc, ok := dom.GetWindow().Document().(dom.HTMLDocument); ok {
		tDoc.SetTitle(c.Title)
	}
	xjs.Body().AppendChild(canvas)
	w.SetMonitor(c.Monitor.Data())

	w.canvas = canvas
	w.context = ctx

	canvas.AddEventListener("mousemove", false, func(m dom.Event) {
		w.mouseX = m.Underlying().Get("clientX").Float()
		w.mouseY = m.Underlying().Get("clientY").Float()
		m.StopPropagation()
		m.PreventDefault()
	})

	mb := func(set bool) func(dom.Event) {
		return func(m dom.Event) {
			button := m.Underlying().Get("button").String()
			if set {
				w.keys["mouse_"+button] = struct{}{}
			} else {
				delete(w.keys, "mouse_"+button)
			}
			m.StopPropagation()
			m.PreventDefault()
		}
	}

	canvas.AddEventListener("mousedown", false, mb(true))
	canvas.AddEventListener("mouseup", false, mb(false))

	w.keys = make(map[string]struct{})

	kb := func(set bool) func(dom.Event) {
		return func(k dom.Event) {
			key := k.Underlying().Get("key").String()
			if set {
				w.keys[key] = struct{}{}
			} else {
				delete(w.keys, key)
			}
			k.StopPropagation()
			k.PreventDefault()
		}
	}

	canvas.AddEventListener("keydown", false, kb(true))
	canvas.AddEventListener("keyup", false, kb(false))
	return nil
}

func (w *webglengine) Loop(run func(int, int, float64)) {
	w.fn = run
	w.rafID = raf.Invoke(w.loop).Int()
}

func (w *webglengine) Uninit() error {
	p := w.canvas.ParentNode()
	if p != nil {
		p.RemoveChild(w.canvas)
	}
	w.canvas = nil
	w.context = nil
	return nil
}

func (w *webglengine) Close() {
	uraf.Invoke(w.rafID)
	w.rafID = 0
}

func (w *webglengine) loop(t float64) { // DOMHighResTimeStamp ??
	w.fn(w.canvas.Width, w.canvas.Height, t)
	w.rafID = raf.Invoke(w.loop).Int()
}

var keyMap = map[engine.Key]string{
	engine.MouseLeft:   "mouse_0",
	engine.MouseMiddle: "mouse_1",
	engine.MouseRight:  "mouse_2",
	engine.KeyEscape:   "Escape",
	engine.KeyUp:       "ArrowUp",
	engine.KeyDown:     "ArrowDown",
	engine.KeyLeft:     "ArrowLeft",
	engine.KeyRight:    "ArrowRight",
}

func (w *webglengine) KeyPressed(k engine.Key) bool {
	kn, ok := keyMap[k]
	if !ok {
		return false
	}
	_, keyPressed := w.keys[kn]
	return keyPressed
}

func (w *webglengine) CursorPos() (x, y float64) {
	return w.mouseX, w.mouseY
}

func (w *webglengine) Context() *webgl.Context {
	return w.context
}

func (w *webglengine) GetMonitors() []*engine.Monitor {
	if w.canvas.Get("requestFullscreen") != nil {
		return []*engine.Monitor{
			engine.NewMonitor("Fullscreen", fullscreen),
			engine.NewMonitor("Browser", browser),
		}
	}
	return []*engine.Monitor{
		engine.NewMonitor("Browser", browser),
	}
}

type monitor int

const (
	browser monitor = iota
	fullscreen
)

func (w *webglengine) GetModes(m interface{}) []engine.Mode {
	t, ok := m.(monitor)
	if !ok {
		return nil
	}
	screen := js.Global.Get("screen")
	var width, height int
	if t == browser {
		width = screen.Get("availWidth").Int()
		height = screen.Get("availHeight").Int()
	} else {
		width = screen.Get("width").Int()
		height = screen.Get("height").Int()
	}
	return []engine.Mode{
		engine.Mode{
			Width:   width,
			Height:  height,
			Refresh: 60,
		},
	}
}

func (w *webglengine) SetMode(m interface{}, mode engine.Mode) {
	w.canvas.Width = mode.Width
	w.canvas.Height = mode.Height
	w.SetMonitor(m)
}

func (w *webglengine) SetMonitor(m interface{}) {
	var mon monitor
	if m != nil {
		mon, _ = m.(monitor)
	}
	if mon == fullscreen {
		if w.canvas.Get("requestFullscreen") != nil {
			w.canvas.Call("requestFullscreen")
		}
	} else if w.canvas.Get("exitFullscreen") != nil {
		w.canvas.Call("exitFullscreen")
	}
}
