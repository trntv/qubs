package broker

import (
	"sync"
	"testing"
)

func TestEnqDeq(t *testing.T) {
	var wg sync.WaitGroup
	q := NewQueue()

	enq := func() {
		for i := 0; i < 128; i++ {
			q.Enqueue(NewMessage([]byte(""), false))
		}
		wg.Done()
	}

	wg.Add(4)
	for i := 0; i < 4; i++ {
		go enq()
	}

	wg.Wait()

	var i uint64
	for {

		msg := q.Dequeue()

		if msg == nil {
			if i == 0 {
				t.Error("empty queue")
			}
			break
		}

		if msg.delivery.index != i {
			t.Error("broken seq")
		}

		i++
	}

	if i != 512 {
		t.Errorf("msg count mismatch %d != 512", i)
	}
}
