package shader

import (
	"errors"
	"lengine/logger"
	"lengine/resources"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
)

const (
	VERTEX   = gl.VERTEX_SHADER
	FRAGMENT = gl.FRAGMENT_SHADER
	GEOMETRY
)

type Shader struct {
	SourceCode string
	IsUsable   bool
	glID       uint32
	shaderType uint32
}

type ShaderProgram struct {
}

func CheckShaderCompileErrors(handle uint32) (bool, error) {
	var success int32
	gl.GetProgramiv(handle, gl.LINK_STATUS, &success)
	if success == 0 {
		var logLength int32
		gl.GetProgramiv(handle, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(handle, logLength, nil, gl.Str(log))

		return false, errors.New(log)
	}
	return true, nil
}

func NewShader(sourceCode string, t uint32) *Shader {
	shader := Shader{SourceCode: sourceCode, shaderType: t}

	handle := gl.CreateShader(t)

	csources, free := gl.Strs(sourceCode)
	gl.ShaderSource(handle, 1, csources, nil)
	free()
	gl.CompileShader(handle)

	// Check if Usable
	usable, err := CheckShaderCompileErrors(handle)
	if err != nil {
		errorString := "Shader is unusable"
		errorString += "\r\n  Reason is: '" + err.Error() + "'"
		errorString += "\r\n  Src is: \r\n"
		errorString += sourceCode
		logger.Error(errorString)
	}

	shader.IsUsable = usable
	shader.glID = handle

	return &shader
}

func NewShaderFromBinaryData(data resources.BinaryData, t uint32) *Shader {
	sourceCode := data.String()
	return NewShader(sourceCode, t)
}
