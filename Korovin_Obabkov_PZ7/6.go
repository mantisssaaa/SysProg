package main

import (
	"fmt"
	"sync"
)

type EventBus struct {
	subs map[string][]func(interface{})
	mu   sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subs: make(map[string][]func(interface{})),
	}
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subs[event] = append(eb.subs[event], handler)
}

func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	if handlers, ok := eb.subs[event]; ok {
		for _, handler := range handlers {
			handler(data)
		}
	}
}

func main() {
	bus := NewEventBus()
	bus.Subscribe("news", func(data interface{}) {
		news := data.(string)
		fmt.Printf("Новость: %s\n", news)
	})
	bus.Subscribe("weather", func(data interface{}) {
		temp := data.(int)
		fmt.Printf("Температура: %d°C\n", temp)
	})
	bus.Subscribe("order", func(data interface{}) {
		order := data.(map[string]string)
		fmt.Printf("Прибыл заказ #%s для %s\n", order["id"], order["customer"])
	})

	bus.Publish("news", "Важное объявление!")
	bus.Publish("weather", 25)
	bus.Publish("order", map[string]string{"id": "1", "customer": "Дмитрий"})
}
