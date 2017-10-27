package systemMovement

import (
	"lengine/component/movement"
	"lengine/component/position"
	"lengine/entity"
	"lengine/event"
	"lengine/system"
	"log"
	"reflect"
)

type SystemMovement struct {
	system.BaseSystem
}

func (sys *SystemMovement) ValidateEntity(ent entity.IEntity) {
	v := reflect.ValueOf(ent)
	t := v.Type()

	interfacePos := reflect.TypeOf((*position.Interface)(nil)).Elem()
	if !t.Implements(interfacePos) {
		return
	}

	interfaceMovement := reflect.TypeOf((*movement.Interface)(nil)).Elem()
	if !t.Implements(interfaceMovement) {
		return
	}

	sys.AddEntity(ent)
}

func (sys *SystemMovement) Update() {
	for _, value := range sys.Entities {
		pos, ok := value.(position.Interface)
		if !ok {
			log.Println("POS - SERIOUS BUG")
			return
		}

		move, ok := value.(movement.Interface)
		if !ok {
			log.Println("MOVEMENT - SERIOUS BUG")
			return
		}

		sys.UpdateMovement(pos, move)
	}
}

func (sys *SystemMovement) UpdateMovement(pos position.Interface, movement movement.Interface) {
	newPosition := pos.GetPosition()

	velocity := movement.GetVelocity()
	newPosition = newPosition.AddVec3(velocity)
	velocity = velocity.AddVec3(movement.GetAcceleration())

	movement.SetVelocity(velocity)

}

func (sys *SystemMovement) HandleEvent(e *event.Event) {

}
