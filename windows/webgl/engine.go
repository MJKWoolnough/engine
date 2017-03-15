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

var raf, uraf *js.Object

func init() {
	raf = js.Global.Get("requestAnimationFrame")
	uraf = js.Global.Get("cancelAnimationFrame")
	if raf != nil {
		w := &webglengine{}
		engine.RegisterWindow(w)
		engine.RegisterInput(w)
	}
}

type webglengine struct {
	canvas         *dom.HTMLCanvasElement
	context        *webgl.Context
	fn             func(int, int, float64)
	rafID          int
	mouseX, mouseY float64
}

func (w *webglengine) Loop(c engine.Config, run func(int, int, float64) bool) error {

	canvas := xdom.Canvas()
	canvas.Width = c.Width
	canvas.Height = c.Height
	ctx, err := webgl.NewContext(canvas)

	if err != nil {
		return err
	}

	if tDoc, ok := dom.GetWindow().Document().(dom.HTMLDocument); ok {
		tDoc.SetTitle(c.Title)
	}
	xjs.Body().AppendChild(canvas)

	w.canvas = canvas
	w.context = ctx
	w.fn = run

	// set up mouse and keyboard event handlers

	w.rafID = raf.Invoke(w.loop).Int()
}

func (w *webglengine) Close() {
	uraf.Invoke(w.rafID)
	xjs.Body().RemoveChild(w.canvas)
	w.canvas = nil
	w.rafID = 0
	w.context = nil
}

func (w *webglengine) loop(t float64) { // DOMHighResTimeStamp ??
	w.fn(w.canvas.Width, w.canvas.Height, t)
	w.rafID = raf.Invoke(w.loop).Int()
}

func (w *webglengine) KeyPressed(k engine.Key) bool {
	return false
}

func (w *webglengine) CursorPos() (x, y float64) {
	return w.mouseX, w.mouseY
}
