package producer

import (
	"fmt"

	"dgf.io/pipe"
)

type Producer struct {
	counter int
	pipe    pipe.IPipe
}

func NewProducer(p pipe.IPipe) *Producer {
	return &Producer{
		pipe:    p,
		counter: 0,
	}
}

func (pr *Producer) Produce() {
	// define id and value
	id := fmt.Sprintf("%d", pr.counter)
	value := fmt.Sprintf("produced text with id %d", pr.counter)
	_ = pr.pipe.Write(id, value)
	pr.counter = pr.counter + 1
}
