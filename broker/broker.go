package broker

type Broker struct {
	hubs map[string]*Hub
}

func NewBroker() *Broker {
	return &Broker{
		hubs: make(map[string]*Hub),
	}
}

func (b *Broker) GetQueue(hubName string, queueName string) *Queue {
	return b.getHub(hubName).getQueue(queueName)
}

func (b *Broker) Dispatch(msg *Message) {
	b.getHub(msg.hub).getQueue(msg.queue).dispatch(msg)
}

func (b *Broker) getHub(name string) *Hub {
	hub, ok := b.hubs[name]
	if !ok {
		hub = b.createHub(name)
	}

	return hub
}

func (b *Broker) createHub(name string) *Hub {
	hub := &Hub{
		name: name,
		broker: b,
		queues: make(map[string]*Queue, 0),
	}

	b.hubs[name] = hub

	return hub
}
