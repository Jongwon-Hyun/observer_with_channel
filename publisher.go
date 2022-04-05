package pub_sub

type Publisher interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Publish(message string)
}

type Emergency struct {
	observers []Observer
}

func NewEmergency() *Emergency {
	return &Emergency{}
}

func (e *Emergency) Subscribe(observer Observer) {
	e.observers = append(e.observers, observer)
	go observer.ListenOnChannel()
}

func (e *Emergency) Unsubscribe(observer Observer) {
	for i, o := range e.observers {
		if o.GetName() == observer.GetName() {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
			return
		}
	}
}

func (e *Emergency) Publish(message string) {
	for _, o := range e.observers {
		o.Receiver(message)
	}
}
