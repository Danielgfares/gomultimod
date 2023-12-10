package main

import (
	"sync"
	"time"

	"dgf.io/inter/comm/channel"
	"dgf.io/inter/comm/consumer"
	"dgf.io/inter/comm/producer"
)

func main() {

	channel := channel.NewPipe()
	pr := producer.NewProducer(channel)
	co := consumer.NewConsumer(channel)

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
