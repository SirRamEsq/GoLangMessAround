package renderable

import "lengine/renderer/drawcall"

type Renderable interface {
	Render() *drawcall.DrawCall
}
