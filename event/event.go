package event

import (
	"lengine/entity"
)

type Type int

const (
	KeyUp Type = iota
	KeyDown
)

//Event interface encapsulates a basic event
type Event interface {
	Sender() entity.EID
	Reciever() entity.EID
	Type() Type
	Message() string
}

func String(e Event) string {
	returnValue := "Event: \n" +
		"  Sender  | " + e.Sender().String() + "\n" +
		"  Reciever| " + e.Reciever().String() + "\n" +
		"  Message | " + e.Message() + "\n"

	return returnValue
}

//SendRecieve implements basic Sender/Reciever functionality for Events
type SendRecieve struct {
	reciever entity.EID
	sender   entity.EID
}

func (sr *SendRecieve) Sender() entity.EID {
	return sr.sender
}

func (sr *SendRecieve) Reciever() entity.EID {
	return sr.reciever
}

//Messager implements Message() functionality for events
type Messager struct {
	message string
}

func (m *Messager) Messager() string {
	return m.message
}

//BasicEvent is a concrete type of a basic event
type BasicEvent struct {
	SendRecieve
	Messager
	T Type
}

func (e *BasicEvent) Type() Type {
	return e.T
}
