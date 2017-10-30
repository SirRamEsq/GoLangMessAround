package systemMovement

import (
	"lengine/component/movement"
	"lengine/component/position"
	"lengine/coordinates/vector"
	"lengine/entity"
	"lengine/event"
	log "lengine/logger"
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

	move, _ := ent.(movement.Interface)
	move.SetMaxVelocity(vector.Vec3{16, 16, 16})

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
		pos, ok := value.(position.Interface)
		if !ok {
			log.Error("Cannot get position interface")
			return
		}

		move, ok := value.(movement.Interface)
		if !ok {
			log.Error("Cannot get movement interface")
			return
		}

		sys.UpdateMovement(pos, move)
	}
}

func (sys *SystemMovement) UpdateMovement(pos position.Interface, movement movement.Interface) {
	newPosition := pos.GetPosition()

	maxVelocity := movement.GetMaxVelocity()
	velocity := movement.GetVelocity()

	if velocity.X > maxVelocity.X {
		velocity.X = maxVelocity.X
	}
	if velocity.Y > maxVelocity.Y {
		velocity.Y = maxVelocity.Y
	}
	if velocity.Z > maxVelocity.Z {
		velocity.Z = maxVelocity.Z
	}
	if velocity.X < -maxVelocity.X {
		velocity.X = -maxVelocity.X
	}
	if velocity.Y < -maxVelocity.Y {
		velocity.Y = -maxVelocity.Y
	}
	if velocity.Z < -maxVelocity.Z {
		velocity.Z = -maxVelocity.Z
	}

	newPosition = newPosition.AddVec3(velocity)
	velocity = velocity.AddVec3(movement.GetAcceleration())

	pos.SetPosition(newPosition)
	movement.SetVelocity(velocity)
}

func (sys *SystemMovement) HandleEvent(e *event.Event) {

}
