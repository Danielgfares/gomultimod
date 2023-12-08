package main

import (
	"sync"

	"dgf.io/consumer"
	"dgf.io/producer"
)

func main() {

	pipe := make(map[string]string)
	pr := producer.NewProducer(&pipe)
	co := consumer.NewConsumer(&pipe)

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
