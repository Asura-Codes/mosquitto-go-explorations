package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/eclipse/paho.golang/paho"
)

func main() {
	brokerAddr := "localhost:1883"
	if envAddr := os.Getenv("MQTT_BROKER"); envAddr != "" {
		brokerAddr = envAddr
	}

	conn, err := net.Dial("tcp", brokerAddr)
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}

	client := paho.NewClient(paho.ClientConfig{
		Conn: conn,
	})

	// Connect to the broker with MQTT v5
	_, err = client.Connect(context.Background(), &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		ClientID:   "lab7-publisher",
	})
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	fmt.Println("Publisher connected using MQTT v5")

	topic := "v5/demo"

	for i := 1; i <= 10; i++ {
		payload := fmt.Sprintf("Message #%d", i)

		// Create a Publish packet with v5 features
		pub := &paho.Publish{
			Topic:   topic,
			QoS:     1,
			Payload: []byte(payload),
			Properties: &paho.PublishProperties{
				// 1. User Properties (Metadata/Headers)
				User: paho.UserProperties{
					{Key: "app-version", Value: "1.0.0"},
					{Key: "sequence", Value: fmt.Sprintf("%d", i)},
				},
				// 2. Message Expiry Interval (seconds)
				// If a subscriber isn't there and the message is queued, it expires.
				MessageExpiry: uint32Ptr(60),
			},
		}

		_, err := client.Publish(context.Background(), pub)
		if err != nil {
			log.Printf("Error publishing: %v", err)
		} else {
			fmt.Printf("[PUB] Sent: %s\n", payload)
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Publisher finished.")
}

func uint32Ptr(i uint32) *uint32 {
	return &i
}
