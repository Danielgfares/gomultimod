package consumer

import (
	"fmt"

	"dgf.io/inter/comm/channel"
)

type Consumer struct {
	counter int
	channel channel.IChannel
}

func NewConsumer(c channel.IChannel) *Consumer {
	return &Consumer{
		channel: c,
		counter: 0,
	}
}

func (co *Consumer) Consume() {
	// Read value by id
	id := fmt.Sprintf("%d", co.counter)
	value, found := co.channel.Read(id)

	if found {
		fmt.Println(value)
		co.counter = co.counter + 1
	}
}
