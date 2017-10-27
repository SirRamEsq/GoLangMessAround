package entity

import "strconv"

/*
  We have almost endless EIDs
  To put into perspective,
  If we generate 200 entities per frame @ 60fps using a uint64 for EIDS then
  we can keep going for 48,745,201 Years
  ((2^64) / (60 * 200)) / 60 / 60 / 24 / 365

  with a uint32 we get 4 days
  ((2^32) / (60 * 200)) / 60 / 60 / 24 = 4.14
*/

//EID is an Entity ID; a unique identifier for entities
type EID uint64

//initialize to 1, 0 refers to no entitiy
var currentEID EID = 1

//New returns an unused EID
func New() EID {
	returnValue := currentEID
	currentEID++
	return returnValue
}

func (eid EID) String() string {
	return "EID: " + strconv.FormatUint(uint64(eid), 10)
}

type IEntity interface {
	EID() EID
}

type Entity struct {
	eid EID
}

func (e *Entity) EID() EID {
	return e.eid
}

/*
Prefabs can be implemented like so

type prefab_1 struct{
	Entity
	ComponentPosition
	ComponentCollision
}

func NewPrefab_1 struct{
	prefab := prefab_1{}
	...
	ini code...
	...

	return prefab
}

Upon creations, prefabs will be registered with systems that will examine the prefab
via reflection to determine if it can be added
*/
