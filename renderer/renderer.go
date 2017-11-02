package renderer

import (
	"lengine/renderer/camera"
	"lengine/renderer/drawcall"
	"lengine/renderer/renderable"
)

//Order in which renderables are rendered
type Order int
type CameraID int
type RenderID int

//Renderer is responsible for rendering everything in a scene
type Renderer struct {
	cameras         map[CameraID]camera.ICamera
	renderables     []renderable.Renderable
	highestCameraID CameraID
}

func (r *Renderer) Init() {
	r.cameras = make(map[CameraID]camera.ICamera)
	r.highestCameraID = 1
}

func (r *Renderer) AddCamera(cam camera.ICamera) CameraID {
	camID := r.highestCameraID
	r.cameras[camID] = cam
	r.highestCameraID++
	return camID
}

func (r *Renderer) RemoveCamera(id CameraID) {
	delete(r.cameras, id)
}

func (r *Renderer) GenerateDrawCalls() []*drawcall.DrawCall {
	drawCalls := make(map[CameraID]*[]*drawcall.DrawCall)
	for k, v := range r.cameras {
		v.Bind()
		/*
			Renderables are read only
			should be able to process
			all cameras concurrently

			if sub cameras are a thing (they should be)
			they can in turn be processed concurrently as well
		*/
		drawCalls[k] = v.Render(&r.renderables)
	}
	//consolodate into a single array and return
}

func (r *Renderer) SortDrawCalls(calls []*drawcall.DrawCall) {
	/*
		Can use many go routines here
			Sort by Layer
				go Opague first
					go Sort by Material
						go Sort by Depth (front to back)
				go Translucent last
					go Sort by Depth (back to front)

				go Opague first
					go Sort by Depth (back to front)
				go Translucent last
					go Sort by Depth (back to front)
	*/

}

func (r *Renderer) Render() {
	drawCalls := (r.GenerateDrawCalls())
	r.SortDrawCalls(drawCalls)

	for _, v := range drawCalls {
		v.Draw()
	}
}
