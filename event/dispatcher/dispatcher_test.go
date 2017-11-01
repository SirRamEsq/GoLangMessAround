package dispatcher

import (
	"lengine/entity"
	"lengine/entity/manager"
	"lengine/event"
	"lengine/testing/comparison"
	"testing"
)

type prefab struct {
	entity.Entity
	LastKeyDown string
	LastKeyUp   string
	LastMessage string
	LastEID     entity.EID
}

func (p *prefab) HandleEvent(e event.Event) {
	p.LastMessage = e.GetMessage()
	p.LastEID = e.GetSender()
	switch e.GetType() {
	case event.KEY_DOWN:
		p.LastKeyDown = e.GetMessage()
	case event.KEY_UP:
		p.LastKeyUp = e.GetMessage()
	}
}

func TestSendRecieve(t *testing.T) {
	dis := EventDispatcher{}
	dis.Init()
	eMan := manager.EntityManager{}

	ent1 := prefab{}
	ent1.SetEID(eMan.NewEID())

	ent2 := prefab{}
	ent2.SetEID(eMan.NewEID())

	dis.Register(ent1.EID(), &ent1)
	dis.Register(ent2.EID(), &ent2)

	inputEID := entity.EID_SUB_INPUT
	dis.Listen(event.KEY_DOWN, inputEID, &ent1)
	dis.Listen(event.KEY_UP, inputEID, &ent1)

	keyUp := event.BasicEvent{T: event.KEY_UP, Message: "Up", Sender: inputEID}
	keyDown := event.BasicEvent{T: event.KEY_DOWN, Message: "Up", Sender: inputEID}
	weird := event.BasicEvent{T: event.MISC, Message: "Other", Sender: ent1.EID()}

	dis.Broadcast(&keyUp)
	dis.Broadcast(&keyDown)
	dis.Send(&weird, ent2.EID())

	comparison.CompareEqualityString(ent1.LastKeyDown, keyDown.Message, t)
	comparison.CompareEqualityString(ent1.LastKeyUp, keyUp.Message, t)
	comparison.CompareEqualityString(ent1.LastMessage, keyDown.Message, t)
	comparison.CompareEqualityEID(ent1.LastEID, keyDown.Sender, t)

	comparison.CompareEqualityString(ent2.LastMessage, weird.Message, t)
	comparison.CompareEqualityEID(ent2.LastEID, weird.Sender, t)
	comparison.CompareInequalityString(ent2.LastKeyDown, keyDown.Message, t)
	comparison.CompareInequalityString(ent2.LastKeyUp, keyUp.Message, t)
}
