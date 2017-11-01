package manager

import (
	"lengine/entity"
	"lengine/event"
	"lengine/system"
)

type EntityManager struct {
	systems     []system.ISystem
	maxInUseEID entity.EID
}

func (eMan *EntityManager) AddSystem(sys system.ISystem) {
	eMan.systems = append(eMan.systems, sys)
}

func (eMan *EntityManager) Init() {
	eMan.maxInUseEID = entity.EID_MIN
}

func (eMan *EntityManager) NewEID() entity.EID {
	returnValue := eMan.maxInUseEID
	eMan.maxInUseEID += 1
	return returnValue
}

//NewEntity will try to add the passed entity to each subsystem
func (eMan *EntityManager) AddEntity(entity entity.IEntity) {
	for _, value := range eMan.systems {
		value.ValidateAddEntity(entity)
	}
}

func (eMan *EntityManager) Update() {
	for _, value := range eMan.systems {
		value.Update()
	}
}

func (eMan *EntityManager) HandleEvent(event *event.Event) {
	for _, value := range eMan.systems {
		value.HandleEvent(event)
	}
}

func NewManager() EntityManager {
	mngr := EntityManager{}
	mngr.Init()
	return mngr
}
