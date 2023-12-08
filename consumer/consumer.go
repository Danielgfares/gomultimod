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
	message := co.pipe[id]

	fmt.Println(message)

	co.counter = co.counter + 1
}
