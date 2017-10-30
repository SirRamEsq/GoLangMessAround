package position

import "lengine/coordinates/vector"

type Interface interface {
	SetPosition(vector.Vec3)
	GetPosition() vector.Vec3
}

type Position struct {
	position vector.Vec3
}

func (pos *Position) SetPosition(newPos vector.Vec3) {
	pos.position = newPos
}

func (pos *Position) GetPosition() vector.Vec3 {
	return pos.position
}
