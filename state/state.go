package state

import (
	"lengine/entity/manager"
	"lengine/event"
)

type State interface {
	Update()
	Render()
	HandleEvent(*event.Event)
}

type PrimaryState struct {
	entityMan manager.EntityManager
}
