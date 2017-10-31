package main

import "lengine/state"

func main() {
	firstState := state.PrimaryState{}
	firstState.Init()
	firstState.Update()
}
