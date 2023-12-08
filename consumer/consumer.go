package consumer

import (
	"fmt"

	"dgf.io/pipe"
)

type Consumer struct {
	counter int
	pipe    pipe.IPipe
}

func NewConsumer(p pipe.IPipe) *Consumer {
	return &Consumer{
		pipe:    p,
		counter: 0,
	}
}

func (co *Consumer) Consume() {
	id := fmt.Sprintf("%d", co.counter)
	value, found := co.pipe.Read(id)

	if found {
		fmt.Println(value)
		co.counter = co.counter + 1
	}
}
