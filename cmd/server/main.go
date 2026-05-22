package main

import (
	"fmt"
	"log"

	"github.com/arshpsps/semq/internal/broker"
)

func main() {
	b := broker.NewBroker()

	offset, err := b.Produce("orders", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("produced at offset:", offset)

	offset, err = b.Produce("orders", []byte("world"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("produced at offset:", offset)

	msgs, err := b.Consume("orders", 0, 10)
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range msgs {
		fmt.Printf("offset=%d timestamp=%d payload=%s\n",
			msg.Offset,
			msg.Timestamp,
			string(msg.Payload),
		)
	}
}
