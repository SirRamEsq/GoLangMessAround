package state

import (
	"lengine/entity/manager"
	"lengine/event"
	"lengine/system/systemMovement"
)

type State interface {
	Init()
	Update()
	HandleEvent(*event.Event)
}

type PrimaryState struct {
	entityMan       manager.EntityManager
	eventDispatcher dispatcher.EventDispatcher
}

func (state *PrimaryState) Init() {
	state.entityMan = manager.NewManager()
	move := systemMovement.SystemMovement{}
	state.entityMan.AddSystem(&move)
}

func (state *PrimaryState) Update() {
	state.entityMan.Update()
	state.Render()
}

func (state *PrimaryState) Render() {

}

func (state *PrimaryState) HandleEvent(e *event.Event) {

}
