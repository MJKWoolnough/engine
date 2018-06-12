package egl // import "vimagination.zapto.org/engine/windows/egl"

import (
	"runtime"
	"sync/atomic"
	"time"

	"github.com/remogatto/egl"
	"vimagination.zapto.org/engine"
)

type eglengine struct {
	close   uint32
	display egl.Display
	surface egl.Surface
	context egl.Context
	start   time.Time
}

func init() {
	runtime.LockOSThread()
	engine.RegisterWindow(new(eglengine))
}

func (e *eglengine) WindowInit(c engine.Config) error {
	/*
		e.display = egl.GetDisplay(egl.DEFAULT_DISPLAY)
		egl.Initialize(e.display, nil, nil)
		var (
			configs   egl.Config
			numConfig int32
		)
		//egl.GetConfigs(e.display, &configs, configSize, &numConfig)
		egl.ChooseConfig(e.display, nil, &configs, 1, &numConfig)
		e.context = egl.CreateContext(e.display, configs, egl.NO_CONTEXT, nil)
		e.surface = egl.CreateWindowSurface(e.display, configs, nil, nil)
		egl.MakeCurrent(e.display, e.surface, e.surface, e.context)
	*/
	/*
		x, err := xgbutil.NewConn()
		if err != nil {
			return err
		}
		win, err := xwindow.Generate(x)
		if err != nil {
			return err
		}
		win.Create(x.RootWin(), 0, 0, 800, 600, xproto.CwBackPixel|xproto.CwEventMask, 0, xproto.EventMaskButtonRelease)

		win.WMGracefulClose(func(w *xwindow.Window) {

		})
		win.Map()
		go xevent.Main(x)
		state := xorg.Initialize(egl.NativeWindowType(uintptr(win.Id)), xorg.DefaultConfigAttributes, xorg.DefaultContextAttributes)
		e.display = state.Display
		e.surface = state.Surface
		e.context = state.Context
		e.start = time.Now()
		egl.MakeCurrent(e.display, e.surface, e.surface, e.context)
	*/
	state := initEGL(800, 600)
	e.display = state.Display
	e.surface = state.Surface
	e.context = state.Context
	e.start = time.Now()
	egl.MakeCurrent(e.display, e.surface, e.surface, e.context)
	return nil
}

func (e *eglengine) Loop(run func(int, int, float64) bool) {
	for atomic.LoadUint32(&e.close) == 0 {
		var w, h int32
		egl.QuerySurface(e.display, e.surface, egl.WIDTH, &w)
		egl.QuerySurface(e.display, e.surface, egl.HEIGHT, &h)
		if run(int(w), int(h), time.Now().Sub(e.start).Seconds()) {
			egl.SwapBuffers(e.display, e.surface)
		}
	}
}

func (e *eglengine) Close() {
	atomic.StoreUint32(&e.close, 1)
}

func (e *eglengine) WindowUninit() error {
	egl.DestroySurface(e.display, e.surface)
	egl.DestroyContext(e.display, e.context)
	egl.Terminate(e.display)
	return nil
}

func (e *eglengine) GetMonitors() []*engine.Monitor {
	return []*engine.Monitor{
		engine.NewMonitor("DEFAULT", egl.DEFAULT_DISPLAY),
	}
}

func (e *eglengine) GetModes(monitor interface{}) []engine.Mode {
	return []engine.Mode{
		{Width: 100, Height: 100, Refresh: 60},
	}
}

func (e *eglengine) SetMode(monitor interface{}, mode engine.Mode) {

}
