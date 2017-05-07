package engine

import "math"

var identity2D = Transform2D{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
}

type Transform2D [9]float32

func (t *Transform2D) Reset() {
	*t = identity2D
}

func (t *Transform2D) Translate(x, y float32) {
	t[2] += x*t[0] + y*t[1]
	t[5] += x*t[3] + y*t[4]
	t[8] += x*t[6] + y*t[7]
}

func (t *Transform2D) Rotate(rad float32) {
	s := float32(math.Sin(float64(rad)))
	c := float32(math.Cos(float64(rad)))
	/*
		t.Multiply(Transform2D{
			c, -s, 0,
			s, c, 0,
			0, 0, 1,
		})
	*/
	t[0], t[1] = c*t[0]+s*t[1], c*t[1]-s*t[0]
	t[3], t[4] = c*t[3]+s*t[4], c*t[4]-s*t[3]
}

func (t *Transform2D) Scale(x, y float32) {
	t[0] *= x
	t[1] *= y
	t[2] *= x
	t[3] *= x
	t[4] *= y
	t[5] *= y
}

func (t *Transform2D) Multiply(s Transform2D) {
	*t = Transform2D{
		t[0]*s[0] + t[1]*s[3] + t[2]*s[6], t[0]*s[1] + t[1]*s[4] + t[2]*s[7], t[0]*s[2] + t[1]*s[5] + t[2]*s[8],
		t[3]*s[0] + t[4]*s[3] + t[5]*s[6], t[3]*s[1] + t[4]*s[4] + t[5]*s[7], t[3]*s[2] + t[4]*s[5] + t[5]*s[8],
		t[6]*s[0] + t[7]*s[3] + t[8]*s[6], t[6]*s[1] + t[7]*s[4] + t[8]*s[7], t[6]*s[2] + t[7]*s[5] + t[8]*s[8],
	}
}
