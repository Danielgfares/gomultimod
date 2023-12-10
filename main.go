package main

import (
	"sync"
	"time"

	"dgf.io/intercomm/consumer"
	"dgf.io/intercomm/pipe"
	"dgf.io/intercomm/producer"
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
