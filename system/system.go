package system

import (
	"lengine/entity"
	"lengine/event"
)

//System interface
type ISystem interface {
	Update()
	HandleEvent(*event.Event)
	ValidateAddEntity(entity entity.IEntity)
	RemoveEntity(entity entity.EID)
}

type BaseSystem struct {
	Entities map[entity.EID]entity.IEntity
}

func (sys *BaseSystem) AddEntity(ent entity.IEntity) {
	sys.Entities[ent.EID()] = ent
}

func (sys *BaseSystem) Remove(eid entity.EID) {
	sys.Entities[eid] = nil
}
