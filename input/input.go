package input

import (
	"encoding/json"
	"io/ioutil"
	"lengine/entity"
	"lengine/event"
	"lengine/event/dispatcher"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	DefaultInputFile = "keys.json"
)

//KeyMapping maps an sdl keycode to a string
type KeyMapping map[sdl.Keycode]string

type keyEvent struct {
	KeyName string
	EType   event.Type
}

type Input struct {
	mouseX     int
	mouseY     int
	mouseWheel float64

	//Map key name to an array of listening entities
	keyListeners map[string][]entity.EID
	keyMapping   KeyMapping

	simulatedInput []keyEvent
}

func (in *Input) sendEvent(keyName string, t event.Type) {
	keyEvent := event.BasicEvent{Sender: entity.EID_SUB_INPUT, Message: keyName, T: t}
	dispatcher.Broadcast(&keyEvent)
}

func (in *Input) keyPress(keyName string) {
	in.sendEvent(keyName, event.KEY_DOWN)
}

func (in *Input) keyRelease(keyName string) {
	in.sendEvent(keyName, event.KEY_UP)
}

func (in *Input) SimulateKeyPress(keyName string) {
	kEvent := keyEvent{KeyName: keyName, EType: event.KEY_DOWN}
	//appending will change slice address pointed to
	in.simulatedInput = append(in.simulatedInput, kEvent)
}

func (in *Input) SimulateKeyRelease(keyName string) {
	kEvent := keyEvent{KeyName: keyName, EType: event.KEY_UP}
	in.simulatedInput = append(in.simulatedInput, kEvent)
}

func (in *Input) keyCodeToString(key sdl.Keycode) string {
	return in.keyMapping[key]
}

func (in *Input) AddMapping(key int, keyName string) {
	keyCode := sdl.Keycode(key)
	in.keyMapping[keyCode] = keyName
}

func (in *Input) RemoveMapping(key int) {
	keyCode := sdl.Keycode(key)
	delete(in.keyMapping, keyCode)
}

//Processes simulated and actual input
func (in *Input) Update() {
	// Reset mousewheel every frame
	in.mouseWheel = 0.0

	/*
		if (remapKey != "") {
			WriteMapSetKeyToNextInput(remapKey);
			ReadKeyIniFile();
			remapKey = "";
		}
	*/

	// Process simulated Input first
	for _, value := range in.simulatedInput {
		if value.EType == event.KEY_DOWN {
			in.keyRelease(value.KeyName)
		} else {
			in.keyPress(value.KeyName)
		}
	}

	in.simulatedInput = make([]keyEvent, 0)

	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch t := e.(type) {
		case *sdl.KeyDownEvent:
			in.keyPress(in.keyCodeToString(t.Keysym.Sym))
		case *sdl.KeyUpEvent:
			in.keyRelease(in.keyCodeToString(t.Keysym.Sym))
			//case *sdl.MouseWheelEvent:
			//in.mouseWheel = e.wheel.y
		}
	}
}

func GetKeyMapFromFile(fileName string) (KeyMapping, error) {
	mapping := KeyMapping{}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return mapping, err
	}

	mappingTemp := make(map[string]string)
	err = json.Unmarshal(file, &mappingTemp)
	if err != nil {
		return mapping, err
	}

	for k, v := range mappingTemp {
		key, _ := strconv.ParseInt(k, 10, 32)
		mapping[sdl.Keycode(key)] = v
	}

	return mapping, err
}
