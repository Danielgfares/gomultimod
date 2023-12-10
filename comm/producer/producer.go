package producer

import (
	"fmt"

	"dgf.io/inter/comm/channel"
)

type Producer struct {
	counter int
	channel channel.IChannel
}

func NewProducer(c channel.IChannel) *Producer {
	return &Producer{
		channel: c,
		counter: 0,
	}
}

func (pr *Producer) Produce() {
	// Write id and value
	id := fmt.Sprintf("%d", pr.counter)
	value := fmt.Sprintf("produced text with id %d", pr.counter)

	_ = pr.channel.Write(id, value)
	pr.counter = pr.counter + 1
}
