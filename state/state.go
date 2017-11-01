package state

import (
	"lengine/entity"
	"lengine/entity/manager"
	"lengine/event"
	"lengine/event/dispatcher"
	"lengine/system/systemMovement"
)

type State interface {
	Init()
	Resume()
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

	state.eventDispatcher = dispatcher.EventDispatcher{}
	state.eventDispatcher.Init()
	state.eventDispatcher.Register(entity.EID_STATE, state)
	dispatcher.SetActiveDispatcher(&state.eventDispatcher)
}

func (state *PrimaryState) Update() {
	state.entityMan.Update()
	state.Render()
}

func (state *PrimaryState) Render() {

}

//Resume is called after a state pushed after this one has been popped
func (state *PrimaryState) Resume() {
	dispatcher.SetActiveDispatcher(&state.eventDispatcher)
}

func (state *PrimaryState) HandleEvent(e event.Event) {

}
