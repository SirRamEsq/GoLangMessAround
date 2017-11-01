package event

import (
	"lengine/entity"
)

type Type int

const (
	KEY_UP Type = iota
	KEY_DOWN
	MISC
)

//Event interface encapsulates a basic event
type Event interface {
	GetSender() entity.EID
	GetType() Type
	GetMessage() string
}

func String(e Event) string {
	returnValue := "Event: \n" +
		"  Sender  | " + e.GetSender().String() + "\n" +
		//"  Type  | " + e.Type().String() + "\n" +
		"  Message | " + e.GetMessage() + "\n"

	return returnValue
}

//BasicEvent is a concrete type of a basic event
type BasicEvent struct {
	Sender  entity.EID
	Message string
	T       Type
}

func (e *BasicEvent) GetType() Type {
	return e.T
}

func (m *BasicEvent) GetMessage() string {
	return m.Message
}

func (sr *BasicEvent) GetSender() entity.EID {
	return sr.Sender
}
