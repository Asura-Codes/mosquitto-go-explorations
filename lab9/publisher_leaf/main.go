package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/eclipse/paho.golang/paho"
)

func main() {
	// Connect to Leaf Broker (1884)
	conn, err := net.Dial("tcp", "localhost:1884")
	if err != nil {
		log.Fatalf("Failed to connect to Leaf: %v", err)
	}

	client := paho.NewClient(paho.ClientConfig{
		Conn: conn,
	})

	_, err = client.Connect(context.Background(), &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		ClientID:   "publisher-leaf",
	})
	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}
	fmt.Println("Connected to Leaf Broker (1884)")

	topic := "leaf_sensor/temperature"
	payload := []byte("25.5")

	fmt.Printf("Publishing to %s: %s\n", topic, payload)
	_, err = client.Publish(context.Background(), &paho.Publish{
		Topic:   topic,
		QoS:     1,
		Payload: payload,
	})
	if err != nil {
		log.Fatalf("Publish error: %v", err)
	}

	// Give it a moment to ensure network flush before exit
	time.Sleep(1 * time.Second)
	client.Disconnect(&paho.Disconnect{})
	fmt.Println("Published and Disconnected")
}
