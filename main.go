package main

import (
	"sync"
	"time"

	"dgf.io/inter/comm/consumer"
	"dgf.io/inter/comm/pipe"
	"dgf.io/inter/comm/producer"
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
			time.Sleep(time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			co.Consume()
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
