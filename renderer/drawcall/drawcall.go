package drawcall

type RenderLayer uint8

const (
	WORLD RenderLayer = iota
	GUI
	SKYBOX
)

type DrawCall struct {
	//Is opague or partially transparent?
	IsTranslucent bool

	Layer    RenderLayer
	Depth    uint16
	Material uint16
}

func (d *DrawCall) Draw() {

}

/*
First, render opague from front to back
Read and write to Z-buffer
this will enable early "Z-Kill" to simply not render
geometry that is not visible

Secondly, render transparent geometry back to front
Read from the z buffer, but do not write to it
This *should* enable the transparent geometry to render only
what is visible and not prevent other transparent geometry from
being rendered...
...I think

If this works then, for opague geometry, rendering may also be able to be
streamed instead of batched
Can immedietley render draw calls if the GPU isn't doing anything
and Z-sorting will take care of the rest
*/
