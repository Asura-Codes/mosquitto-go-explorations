package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/eclipse/paho.golang/paho"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: deleter <topic>")
		os.Exit(1)
	}
	topic := os.Args[1]

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
		ClientID:   "state_deleter",
	}

	_, err = client.Connect(context.Background(), cp)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	fmt.Printf("--- Nulling Retained Message on Topic: %s ---\n", topic)
	
	// CRITICAL: Publishing an empty payload with Retain=true deletes the retained message on the broker.
	_, err = client.Publish(context.Background(), &paho.Publish{
		Topic:   topic,
		QoS:     1,
		Retain:  true, 
		Payload: []byte{}, // Empty payload
	})
	if err != nil {
		log.Fatalf("Failed to delete retained message: %v", err)
	}

	fmt.Println("Retained message deleted.")
	client.Disconnect(&paho.Disconnect{ReasonCode: 0})
}
