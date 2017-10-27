package entityManager

import (
	"lengine/entity"
	"lengine/event"
	"lengine/system"
)

type EntityManager struct {
	systems []system.ISystem
}

func (eMan *EntityManager) AddSystem(sys system.ISystem) {
	eMan.systems = append(eMan.systems, sys)
}

//NewEntity will try to add the passed entity to each subsystem
func (eMan *EntityManager) NewEntity(entity entity.IEntity) {
	for _, value := range eMan.systems {
		value.AddEntity(entity)
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
