package main

import "github.com/go-gl/gl/v2.1/gl"

func (t triangle) Render(width, height int, time float64) {
	ratio := float64(width) / float64(height)

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-ratio, ratio, -1, 1, 1, -1)
	gl.MatrixMode(gl.MODELVIEW)

	gl.LoadIdentity()
	gl.Rotated(time*50, 0, 0, 1)

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-0.6, -0.4, 0.)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(0.6, -0.4, 0)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(0, 0.6, 0)
	gl.End()
}
