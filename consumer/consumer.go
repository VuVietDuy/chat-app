package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
)

// the topic and broker address are initialized as constants
var (
	topicUser1     string
	topicUser2     string
	broker1Address = "localhost:9092"
)

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topicUser2,
	})

	for {
		var msg string
		fmt.Scanln(&msg)
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(msg),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		i++
	}
}

func consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topicUser1,
		GroupID: "my-group",
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println(string(msg.Value))
	}
}

func main() {
	// create a new context

	ctx := context.Background()

	var username1, username2 string
	fmt.Println("Username1: ")
	fmt.Scanln(&username1)
	fmt.Println("Username2")
	fmt.Scanln(&username2)
	topicUser1 = username1 + "-topic"
	topicUser2 = username2 + "-topic"

	go produce(ctx)
	consume(ctx)
}
