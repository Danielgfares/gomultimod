package consumer

import (
	"fmt"
)

type Consumer struct {
	counter int
	pipe    map[string]string
}

func NewConsumer(p *map[string]string) *Consumer {
	return &Consumer{
		pipe:    *p,
		counter: 0,
	}
}

func (co *Consumer) Consume() {
	id := fmt.Sprintf("%d", co.counter)

	if message, ok := co.pipe[id]; ok {
		fmt.Println(message)
		delete(co.pipe, id)
		co.counter = co.counter + 1
	}
}
