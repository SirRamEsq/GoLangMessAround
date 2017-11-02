package input_test

import (
	"lengine/input"
	"lengine/testing/comparison"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	keyFileName = "input-test.json"
)

func TestCanLoadKeysFromFile(t *testing.T) {
	keyMap, err := input.GetKeyMapFromFile(keyFileName)
	if err != nil {
		t.Error(err.Error())
	}

	keyCode := sdl.Keycode(100)
	keyName := keyMap[keyCode]

	comparison.CompareEqualityString(keyName, "right", t)
}
