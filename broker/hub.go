package broker

type Hub struct {
	name   string
	size   uint64
	queues map[string]*queue
	broker *Broker
}

func (h *Hub) getQueue(name string) *queue {
	queue, ok := h.queues[name]
	if !ok {
		queue = NewQueue()
		h.queues[name] = queue
	}

	return queue
}
