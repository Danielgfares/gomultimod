package main

import (
	"sync"

	"dgf.io/consumer"
	"dgf.io/pipe"
	"dgf.io/producer"
)

func main() {

	pipe := pipe.NewPipe()
	pr := producer.NewProducer(pipe)
	co := consumer.NewConsumer(pipe)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			pr.Produce()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			co.Consume()
		}
	}()

	wg.Wait()
}
