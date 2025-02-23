package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)



func NewTestConfig() *sarama.Config{
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Retry.Backoff = 0
	config.Producer.Retry.Backoff = 0
	config.Version = sarama.MinVersion
	return config
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main(){
	go func() {
		http.HandleFunc("/", helloHandler)
	
		fmt.Println("Starting server at port 8070")
		if err := http.ListenAndServe(":8070", nil); err != nil {
			fmt.Println(err)
		}
		biliboba := [5]int{1,2,3,4,5}
		

	}()
	
	config := NewTestConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                                  sync.WaitGroup
		enqueued, successes, producerErrors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

	ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: "biliboba_topic", Value:  sarama.StringEncoder("testing 123")}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, producerErrors)


}