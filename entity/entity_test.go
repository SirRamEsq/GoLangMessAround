package entity_test

import (
	"lengine/entity"
	"testing"
)

func CompareEqualityEID(expected entity.EID, actual entity.EID, t *testing.T) {
	if expected != actual {
		errorString := "Expected: " + expected.String() + " - Actual: " + actual.String()
		t.Error(errorString)
	}
}
func CompareInequalityEID(expected entity.EID, actual entity.EID, t *testing.T) {
	if expected == actual {
		errorString := "Expected NOT: " + expected.String() + " - Actual: " + actual.String()
		t.Error(errorString)
	}
}

type testPrefab struct {
	entity.Entity
}

func TestNewEID(t *testing.T) {
	eid1 := entity.New()
	prefab := testPrefab{}
	prefab.SetEID(eid1)
	CompareInequalityEID(0, prefab.EID(), t)
	CompareEqualityEID(eid1, prefab.EID(), t)

	eid2 := entity.New()
	prefab.SetEID(eid2)
	CompareInequalityEID(eid2, prefab.EID(), t)
}
