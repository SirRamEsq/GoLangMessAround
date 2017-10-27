package state

import "lengine/event"

type State interface {
	Update()
	Render()
	HandleEvent(*event.Event)
}
