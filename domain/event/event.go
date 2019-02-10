package event

type EventBus interface {
	Subscribe(eventName string, subscriber Subscriber)
	Unsubscribe(eventName string, subscriber Subscriber)
	Publish(eventName string, domain interface{})
}

type eventBus struct {
	subscribers map[string][]Subscriber
}

type Message struct {
	EventName string
	Response  interface{}
}

type Subscriber chan *Message

func NewBus() EventBus {
	return &eventBus{
		subscribers: make(map[string][]Subscriber),
	}
}

func (eb *eventBus) Subscribe(e string, s Subscriber) {
	if eb.subscribers == nil {
		eb.subscribers = make(map[string][]Subscriber)
	}

	if _, ok := eb.subscribers[e]; ok {
		eb.subscribers[e] = append(eb.subscribers[e], s)
	} else {
		eb.subscribers[e] = make([]Subscriber, 0)
		eb.subscribers[e] = append(eb.subscribers[e], s)
	}
}

func (eb *eventBus) Unsubscribe(e string, s Subscriber) {
	if _, ok := eb.subscribers[e]; ok {
		for i := range eb.subscribers[e] {
			if eb.subscribers[e][i] == s {
				eb.subscribers[e] = append(eb.subscribers[e][:i], eb.subscribers[e][i+1:]...)
				break
			}
		}
	}
}

func (eb *eventBus) Publish(e string, response interface{}) {
	if _, ok := eb.subscribers[e]; ok {
		for _, s := range eb.subscribers[e] {
			go func(s Subscriber) {
				s <- &Message{
					EventName: e,
					Response:  response,
				}
			}(s)
		}
	}
}
