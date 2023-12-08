package producer

import "fmt"

type Producer struct {
	counter int
	pipe    map[string]string
}

func NewProducer(p *map[string]string) *Producer {
	return &Producer{
		pipe:    *p,
		counter: 0,
	}
}

func (pr *Producer) Produce() {
	id := fmt.Sprintf("%d", pr.counter)
	message := fmt.Sprintf("produced text with id %d", pr.counter)
	pr.pipe[id] = message
	pr.counter = pr.counter + 1
}
