package camera

import (
	"lengine/coordinates/matrix"
	"lengine/coordinates/vector"
	"lengine/renderer/drawcall"
	"lengine/renderer/renderable"
	"lengine/resources/texture"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type ICamera interface {
	Bind(CameraUBO uint32)
	GetProjection()
	GetModelView()
	Render(*[]renderable.Renderable) *[]*drawcall.DrawCall
}

type Frustrum struct {
	Near float64
	Far  float64
	FOV  float64 //Range 0 - 180
}

type Camera struct {
	Rotation vector.Vec3
	Scale    vector.Vec3
	Pos      vector.Vec3
	Viewport Frustrum

	//Gl
	FrameBufferTexture texture.Texture
	FrameBufferID      uint32
}

func (cam *Camera) GetModelView() matrix.Matrix {
	T := matrix.NewIdentity()
	T = T.Translate(cam.Pos)

	R := matrix.NewIdentity()
	//R = R.RotateX(cam.rotation.X)
	//R = R.RotateY(cam.rotation.Y)
	R = R.RotateZ(cam.Rotation.Z)

	S := matrix.NewIdentity()

	modelView := T.MultiplyMatrix(R.MultiplyMatrix(S))
	return modelView
}

func (cam *Camera) GetProjection() matrix.Matrix {
	viewport := &cam.Viewport
	projection := matrix.NewProjection(viewport.FOV, cam.Aspect(), viewport.Near, viewport.Far)
	return projection
}

func (cam *Camera) Aspect() float64 {
	return 16 / 9
	//return 4 / 3
}

/*
Bind will bind the current camera FBO and insert its camera data into
the passed uniform buffer
*/
func (cam *Camera) Bind(CameraUBO uint32) {
	//figure this out...
	//communicate over channles?
	//pass back a matrix or pointer to matrix?
	//go cam.GetModelView()
	//go cam.GetProjection()

	modelView := cam.GetModelView()
	projection := cam.GetProjection()
	sizeofMatrix := int(unsafe.Sizeof(modelView))

	ptrModelView := unsafe.Pointer(&modelView)
	ptrProjection := unsafe.Pointer(&projection)

	gl.BindBuffer(gl.UNIFORM_BUFFER, CameraUBO)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 0, sizeofMatrix, ptrModelView)
	gl.BufferSubData(gl.UNIFORM_BUFFER, sizeofMatrix, sizeofMatrix, ptrProjection)
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)

	// Push viewport bit
	gl.PushAttrib(gl.VIEWPORT_BIT)
	// Setup frame buffer render
	gl.BindFramebuffer(gl.FRAMEBUFFER, cam.FrameBufferID)
	// Set Render Viewport
	gl.Viewport(0, 0, cam.FrameBufferTexture.Width(),
		cam.FrameBufferTexture.Height())

	// Clear Background
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Atatch buffer texture
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D,
		cam.FrameBufferTexture.ID(), 0)
}

func (cam *Camera) Render(renderables *[]renderable.Renderable) *[]*drawcall.DrawCall {
	//Has capacity of len(renderables)
	drawCalls := make([]*drawcall.DrawCall, 0, len(*renderables))
	for _, v := range *renderables {
		drawCalls = append(drawCalls, v.Render())
	}
	return &drawCalls
}
