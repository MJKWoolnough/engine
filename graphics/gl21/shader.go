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
		return 0, errors.New(log(id))
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
	name := make([]byte, len(uName)+1)
	copy(name, uName)
	r := gl.GetUniformLocation(p.pid, &name[0])
	if r < 0 {
		switch gl.GetError() {
		case gl.INVALID_VALUE:
			return r, errors.New("Invalid Var: " + uName)
		case gl.INVALID_OPERATION:
			return r, errors.New("Invalid Op: " + uName)
		default:
			return r, errors.New("Unknown Var: " + uName)
		}
	}
	return r, nil
}

func (p *Program) GetAttribLocation(aName string) (int32, error) {
	name := make([]byte, len(aName)+1)
	copy(name, aName)
	r := gl.GetAttribLocation(p.pid, &name[0])
	if r < 0 {
		switch gl.GetError() {
		case gl.INVALID_VALUE:
			return r, errors.New("Invalid Var: " + aName)
		case gl.INVALID_OPERATION:
			return r, errors.New("Invalid Op: " + aName)
		default:
			return r, errors.New("Unknown Var: " + aName)
		}
	}
	return r, nil
}

func (p *Program) VertexLog() string {
	return log(p.vid)
}

func (p *Program) FragmentLog() string {
	return log(p.fid)
}

func (p *Program) ProgramLog() string {
	var length int32
	gl.GetProgramiv(p.pid, gl.INFO_LOG_LENGTH, &length)
	if length == 0 {
		return ""
	}
	buf := make([]byte, length)
	gl.GetProgramInfoLog(p.pid, length, &length, &buf[0])
	return string(buf)
}

func log(id uint32) string {
	var length int32
	gl.GetShaderiv(id, gl.INFO_LOG_LENGTH, &length)
	if length == 0 {
		return ""
	}
	buf := make([]byte, length)
	gl.GetShaderInfoLog(id, length, &length, &buf[0])
	return string(buf)
}
