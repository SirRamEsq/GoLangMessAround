package input

type Input struct {
	keyPressed  map[string]bool
	keyReleased map[string]bool
	keyHeld     map[string]bool
}

func (input *Input) Press(key string) {
	input.keyPressed[key] = true
	input.keyHeld[key] = true
}

func (input *Input) Release(key string) {
	input.keyReleased[key] = true
	input.keyHeld[key] = false
}

func (input *Input) Update(key string) {
	input.keyPressed[key] = false
	input.keyReleased[key] = false
}
