package systemMovement

import (
	"lengine/component/movement"
	"lengine/component/position"
	"lengine/entity"
	"lengine/event"
	"lengine/system"
	"reflect"
)

type SystemMovement struct {
	system.BaseSystem
}

func (sys *SystemMovement) ValidateEntity(ent entity.IEntity) bool {
	v := reflect.ValueOf(ent)
	t := v.Type()

	interfacePos := reflect.TypeOf((*position.Interface)(nil)).Elem()
	if !t.Implements(interfacePos) {
		return false
	}

	interfaceMovement := reflect.TypeOf((*movement.Interface)(nil)).Elem()
	if !t.Implements(interfaceMovement) {
		return false
	}

	return true
}

func (sys *SystemMovement) ValidateAddEntity(ent entity.IEntity) bool {
	if !sys.ValidateEntity(ent) {
		return false
	}

	sys.AddEntity(ent)
	return true
}

func (sys *SystemMovement) Update() {
	for _, value := range sys.Entities {
		pos, _ := value.(position.Interface)
		move, _ := value.(movement.Interface)

		sys.UpdateMovement(pos, move)
	}
}

func (sys *SystemMovement) UpdateMovement(pos position.Interface, movement movement.Interface) {
	newPosition := pos.GetPosition()

	//update position
	velocity := movement.GetVelocity()
	newPosition = newPosition.AddVec3(velocity)
	pos.SetPosition(newPosition)

	//update velocity
	velocity = velocity.AddVec3(movement.GetAcceleration())

	//set and clamp velocity
	movement.SetVelocity(velocity)
}

func (sys *SystemMovement) HandleEvent(e *event.Event) {

}
