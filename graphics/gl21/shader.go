package gl21

import (
	"errors"

	"github.com/go-gl/gl/v2.1/gl"
)

type Program struct {
	pid, vid, fid uint32
}

func createShader(typ uint32, source *byte, length int32) (uint32, error) {
	id := gl.CreateShader(typ)
	switch gl.GetError() {
	case gl.INVALID_ENUM:
		return 0, errors.New("Invalid Shader Type")
	}
	gl.ShaderSource(id, 1, &source, &length)
	switch gl.GetError() {
	case gl.INVALID_OPERATION:
		return 0, errors.New("Shader Compile not supported")
	case gl.INVALID_VALUE:
		return 0, errors.New("Invalid Shader ID")
	}
	gl.CompileShader(id)
	var status int32
	gl.GetShaderiv(id, gl.COMPILE_STATUS, &status)
	if status != gl.TRUE {
		return 0, errors.New(log(id, gl.GetShaderiv, gl.GetShaderInfoLog))
	}
	return id, nil
}

func NewProgram(vertexShader, fragmentShader []byte) (*Program, error) {
	vs, err := createShader(gl.VERTEX_SHADER, &vertexShader[0], int32(len(vertexShader)))
	if err != nil {
		return nil, err
	}
	fs, err := createShader(gl.FRAGMENT_SHADER, &fragmentShader[0], int32(len(fragmentShader)))
	if err != nil {
		return nil, err
	}
	pid := gl.CreateProgram()
	gl.AttachShader(pid, vs)
	gl.AttachShader(pid, fs)
	gl.LinkProgram(pid)
	return &Program{
		pid: pid,
		vid: vs,
		fid: fs,
	}, nil
}

func (p *Program) Use() {
	gl.UseProgram(p.pid)
}

func (p *Program) GetUniformLocation(uName string) (int32, error) {
	return getLocation(uName, p.pid, gl.GetUniformLocation)
}

func (p *Program) GetAttribLocation(aName string) (int32, error) {
	return getLocation(aName, p.pid, gl.GetAttribLocation)
}

func getLocation(lname string, pid uint32, lf func(uint32, *uint8) int32) (int32, error) {
	name := make([]byte, len(lName)+1)
	copy(name, lName)
	r := lf(pid, &name[0])
	if r < 0 {
		switch gl.GetError() {
		case gl.INVALID_VALUE:
			return r, errors.New("Invalid Var: " + lName)
		case gl.INVALID_OPERATION:
			return r, errors.New("Invalid Op: " + lName)
		default:
			return r, errors.New("Unknown Var: " + lName)
		}
	}
	return r, nil
}

func (p *Program) VertexLog() string {
	return log(p.vid, gl.GetShaderiv, gl.GetShaderInfoLog)
}

func (p *Program) FragmentLog() string {
	return log(p.fid, gl.GetShaderiv, gl.GetShaderInfoLog)
}

func (p *Program) ProgramLog() string {
	return log(p.pid, gl.GetProgramiv, gl.GetProgramInfoLog)
}

func log(id uint32, ll func(uint32, uint32, *int32), lf func(uint32, int32, *int32, *uint8)) string {
	var length int32
	ll(id, gl.INFO_LOG_LENGTH, &length)
	if length == 0 {
		return ""
	}
	buf := make([]byte, length)
	lf(id, length, &length, &buf[0])
	return string(buf)
}
