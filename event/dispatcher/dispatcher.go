package dispatcher

import (
	"lengine/entity"
	"lengine/event"
)

type EventHandler interface {
	HandleEvent(event.Event)
}

type Listeners map[entity.EID][]EventHandler

type EventDispatcher struct {
	handlers map[event.Type]Listeners
	//all registered event handling entities
	registered map[entity.EID]EventHandler
}

func (e *EventDispatcher) Init() {
	e.registered = make(map[entity.EID]EventHandler)
	e.handlers = make(map[event.Type]Listeners)
}

//Allows an entitiy to be sent events directly
func (e *EventDispatcher) Register(id entity.EID, handler EventHandler) {
	e.registered[id] = handler
}

/*Will register the passed EventHandler to recieve events of
the specified type from the specified eid. Can Pass EID_ALL to listen to
all events of a specified type
*/
func (e *EventDispatcher) Listen(t event.Type, id entity.EID, handler EventHandler) {
	typeHandler, ok := e.handlers[t]
	if !ok {
		e.handlers[t] = make(Listeners)
	}
	listener, ok := typeHandler[id]
	if !ok {
		e.handlers[t][id] = make([]EventHandler, 0)
	}
	e.handlers[t][id] = append(listener, handler)
}

//Will broadcast to all listening entities
func (e *EventDispatcher) Broadcast(thisEvent event.Event) {
	sender := thisEvent.GetSender()
	listeners := e.handlers[thisEvent.GetType()]

	for _, listener := range listeners[sender] {
		listener.HandleEvent(thisEvent)
	}
	for _, listener := range listeners[entity.EID_ALL] {
		listener.HandleEvent(thisEvent)
	}
}

//Will send this event to the specified eid, regardless of whether it is listening
func (e *EventDispatcher) Send(thisEvent event.Event, reciever entity.EID) {
	handler := e.registered[reciever]
	handler.HandleEvent(thisEvent)
}
