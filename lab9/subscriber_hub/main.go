package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/eclipse/paho.golang/paho"
)

func main() {
	// Connect to Hub Broker (1883)
	conn, err := net.Dial("tcp", "localhost:1883")
	if err != nil {
		log.Fatalf("Failed to connect to Hub: %v", err)
	}

	router := paho.NewSingleHandlerRouter(func(m *paho.Publish) {
		fmt.Printf("[Hub Recv] Topic: %s | Payload: %s\n", m.Topic, string(m.Payload))
	})

	client := paho.NewClient(paho.ClientConfig{
		Conn:   conn,
		Router: router,
	})

	_, err = client.Connect(context.Background(), &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		ClientID:   "subscriber-hub",
	})
	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}
	fmt.Println("Connected to Hub Broker (1883)")

	// We subscribe to 'hub_sensor/#' because the bridge maps 'leaf_sensor/' -> 'hub_sensor/'
	topic := "hub_sensor/#"
	fmt.Printf("Subscribing to %s\n", topic)
	
	_, err = client.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: []paho.SubscribeOptions{
			{Topic: topic, QoS: 1},
		},
	})
	if err != nil {
		log.Fatalf("Subscribe error: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	fmt.Println("Subscriber shutting down...")
}
