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
	if addr := os.Getenv("MQTT_BROKER"); addr != "" {
		brokerAddr = addr
	}

	conn, err := net.Dial("tcp", brokerAddr)
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}

	client := paho.NewClient(paho.ClientConfig{
		Conn: conn,
	})

	cp := &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		ClientID:   "state_publisher",
	}

	ca, err := client.Connect(context.Background(), cp)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	if ca.ReasonCode != 0 {
		log.Fatalf("Failed to connect: %v", ca.ReasonCode)
	}

	states := map[string]string{
		"settings/mode":       "automatic",
		"settings/threshold":  "25.5",
		"status/last_update":  time.Now().Format(time.RFC3339),
	}

	fmt.Println("--- Publishing Retained States ---")
	for topic, payload := range states {
		_, err := client.Publish(context.Background(), &paho.Publish{
			Topic:   topic,
			QoS:     1,
			Retain:  true, // CRITICAL: This makes it a stateful message
			Payload: []byte(payload),
		})
		if err != nil {
			log.Printf("Failed to publish to %s: %v", topic, err)
		} else {
			fmt.Printf("Published [Retained]: %s = %s\n", topic, payload)
		}
	}

	client.Disconnect(&paho.Disconnect{ReasonCode: 0})
	fmt.Println("Publisher finished and disconnected.")
}
