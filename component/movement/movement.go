package movement

import "lengine/coordinates/vector"

type Interface interface {
	SetVelocity(vector.Vec3)
	GetVelocity() vector.Vec3
	SetAcceleration(vector.Vec3)
	GetAcceleration() vector.Vec3
}

type Movement struct {
	velocity     vector.Vec3
	acceleration vector.Vec3
}

func (move *Movement) SetVelocity(newSpeed *vector.Vec3) {
	move.velocity = *newSpeed
}
func (move *Movement) SetAcceleration(newSpeed *vector.Vec3) {
	move.acceleration = *newSpeed
}

func (move *Movement) GetVelocity() vector.Vec3 {
	return move.velocity
}
func (move *Movement) GetAcceleration() vector.Vec3 {
	return move.acceleration
}
