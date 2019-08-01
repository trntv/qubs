package broker

type Hub struct {
	name string
	size uint64
	queues map[string]*Queue
	broker *Broker
}

func (h *Hub) getQueue(name string) *Queue {
	queue, ok := h.queues[name]
	if !ok {
		queue = NewQueue(name)
		h.queues[name] = queue
	}

	return queue
}
