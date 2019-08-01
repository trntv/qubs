package broker

import (
	"math"
	"strconv"
	"testing"
)

func TestBroker_Dispatch(t *testing.T) {
	broker := Broker{
		hubs: make(map[string]*Hub),
	}

	n := 1024 * 1024
	for i := 0; i < n; i++ {
		broker.Dispatch(&Message{
			hub:     strconv.Itoa(i % 16),
			queue:   strconv.Itoa(i % 32),
			tags:    []string{strconv.Itoa(i % 64)},
			Payload: []byte{},
		})
	}

	ai := broker.hubs["15"].queues["31"].index
	ei := uint64(math.Pow(2, 15))

	if ai != ei {
		t.Error("wrong number of Messages")
	}
}
