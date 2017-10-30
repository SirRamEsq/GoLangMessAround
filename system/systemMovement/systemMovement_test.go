package systemMovement

import (
	"lengine/component/movement"
	"lengine/component/position"
	"lengine/coordinates/vector"
	"lengine/entity"
	"lengine/testing/comparison"
	"math"
	"testing"
)

type prefabMovePos struct {
	entity.Entity
	position.Position
	movement.Movement
}

func TestValidateEntity(t *testing.T) {
	prefab := prefabMovePos{}
	prefab.SetEID(entity.New())

	system := SystemMovement{}
	system.Entities = make(map[entity.EID]entity.IEntity)

	if !system.ValidateAddEntity(&prefab) {
		errorString := "System could not validate Entity!"
		t.Error(errorString)
	}
}

func TestUpdateMovement(t *testing.T) {
	prefab := prefabMovePos{}
	prefab.SetEID(entity.New())

	system := SystemMovement{}
	system.Entities = make(map[entity.EID]entity.IEntity)

	startPos := vector.Vec3{5, 5, 0}
	startVel := vector.Vec3{1, 2, 0}
	startAcc := vector.Vec3{.1, 1, 0}

	prefab.SetPosition(startPos)
	prefab.SetVelocity(startVel)
	prefab.SetAcceleration(startAcc)

	system.ValidateAddEntity(&prefab)

	system.Update()

	comparison.CompareEqualityFloat(startPos.X+startVel.X, prefab.GetPosition().X, t)
	comparison.CompareEqualityFloat(startPos.Y+startVel.Y, prefab.GetPosition().Y, t)
	comparison.CompareEqualityFloat(startPos.Z+startVel.Z, prefab.GetPosition().Z, t)

	comparison.CompareEqualityFloat(startAcc.X+startVel.X, prefab.GetVelocity().X, t)
	comparison.CompareEqualityFloat(startAcc.Y+startVel.Y, prefab.GetVelocity().Y, t)
	comparison.CompareEqualityFloat(startAcc.Z+startVel.Z, prefab.GetVelocity().Z, t)

	//Make sure speeds get clamped to their max velocity
	maxVelocity := vector.Vec3{11, 18, 17}
	prefab.SetMaxVelocity(maxVelocity)
	zeroPos := vector.Vec3{0, 0, 0}
	tooFast := vector.Vec3{17, -20, 100}

	prefab.SetPosition(zeroPos)
	prefab.SetVelocity(tooFast)
	prefab.SetAcceleration(zeroPos)

	system.Update()

	comparison.CompareEqualityFloat(math.Abs(maxVelocity.X), math.Abs(prefab.GetPosition().X), t)
	comparison.CompareEqualityFloat(math.Abs(maxVelocity.Y), math.Abs(prefab.GetPosition().Y), t)
	comparison.CompareEqualityFloat(math.Abs(maxVelocity.Z), math.Abs(prefab.GetPosition().Z), t)
}
