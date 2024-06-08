package signal

import (
	"log"
	"sync"
)

type topic struct {
	Listeners []chan struct{}
	mu        *sync.Mutex
}

type signal struct {
	Listeners *sync.Map
	Topicsize int
}

func Newtopic() Signal {
	return &signal{
		Listeners: new(sync.Map),
	}
}

func (c *signal) Subscribe(id int) (<-chan struct{}, func(), error) {
	tpicinf0, _ := c.Listeners.LoadOrStore(id, &topic{mu: new(sync.Mutex)})
	t := tpicinf0.(*topic)
	t.mu.Lock()
	defer t.mu.Unlock()
	ch := make(chan struct{}, 1)
	t.Listeners = append(t.Listeners, ch)
	return ch, func() {
		log.Println("INSIDE CHANNEL CLEARING FUNCTION")
		t.mu.Lock()
		defer t.mu.Unlock()
		for i := 0; i < len(t.Listeners); i++ {
			if t.Listeners[i] == ch {
				t.Listeners = append(t.Listeners[:i], t.Listeners[i+1:]...)
			}
		}
	}, nil
}

func (c *signal) Publish(id int) error {
	log.Println("INSIDE PUBLISH")
	log.Println(id)
	tpicinfo, ok := c.Listeners.Load(id)
	if !ok {
		log.Println("returning from publish")
		return nil
	}
	t := tpicinfo.(*topic)
	l := len(t.Listeners)
	if l == 0 {
		return err
	}
	for _, value := range t.Listeners {
		log.Println("channel inside listeners = ", value)
		value <- struct{}{}
	}
	return nil
}
