package movement

import "lengine/coordinates/vector"

type Interface interface {
	SetVelocity(vector.Vec3)
	GetVelocity() vector.Vec3
	SetAcceleration(vector.Vec3)
	GetAcceleration() vector.Vec3

	SetMaxVelocity(vector.Vec3)
	GetMaxVelocity() vector.Vec3
}

type Movement struct {
	velocity     vector.Vec3
	acceleration vector.Vec3
	maxVelocity  vector.Vec3
}

func (move *Movement) SetVelocity(newSpeed vector.Vec3) {
	maxVelocity := move.maxVelocity

	if newSpeed.X > maxVelocity.X {
		newSpeed.X = maxVelocity.X
	}
	if newSpeed.Y > maxVelocity.Y {
		newSpeed.Y = maxVelocity.Y
	}
	if newSpeed.Z > maxVelocity.Z {
		newSpeed.Z = maxVelocity.Z
	}
	if newSpeed.X < -maxVelocity.X {
		newSpeed.X = -maxVelocity.X
	}
	if newSpeed.Y < -maxVelocity.Y {
		newSpeed.Y = -maxVelocity.Y
	}
	if newSpeed.Z < -maxVelocity.Z {
		newSpeed.Z = -maxVelocity.Z
	}
	move.velocity = newSpeed
}
func (move *Movement) SetAcceleration(newSpeed vector.Vec3) {
	move.acceleration = newSpeed
}

func (move *Movement) SetMaxVelocity(newMax vector.Vec3) {
	move.maxVelocity = newMax
}

func (move *Movement) GetVelocity() vector.Vec3 {
	return move.velocity
}
func (move *Movement) GetAcceleration() vector.Vec3 {
	return move.acceleration
}

func (move *Movement) GetMaxVelocity() vector.Vec3 {
	return move.maxVelocity
}
