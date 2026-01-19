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
	workerID := os.Getenv("WORKER_ID")
	if workerID == "" {
		workerID = "default"
	}

	brokerAddr := "localhost:1883"
	if envAddr := os.Getenv("MQTT_BROKER"); envAddr != "" {
		brokerAddr = envAddr
	}

	conn, err := net.Dial("tcp", brokerAddr)
	if err != nil {
		log.Fatalf("Worker %s: Failed to connect: %v", workerID, err)
	}

	// Setup Router for handling incoming messages
	router := paho.NewSingleHandlerRouter(func(m *paho.Publish) {
		fmt.Printf("[WORKER %s] Received: %s\n", workerID, string(m.Payload))
		if m.Properties != nil && len(m.Properties.User) > 0 {
			fmt.Printf("           Properties: %v\n", m.Properties.User)
		}
	})

	client := paho.NewClient(paho.ClientConfig{
		Conn:   conn,
		Router: router,
	})

	// Connect to the broker with MQTT v5
	_, err = client.Connect(context.Background(), &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		ClientID:   "lab7-worker-" + workerID,
	})
	if err != nil {
		log.Fatalf("Worker %s: Connect error: %v", workerID, err)
	}

	// 3. Shared Subscription
	// Format: $share/<ShareName>/<TopicFilter>
	// This tells the broker to distribute messages among clients in the same group.
	sharedTopic := "$share/processing-group/v5/demo"

	_, err = client.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: []paho.SubscribeOptions{
			{Topic: sharedTopic, QoS: 1},
		},
	})
	if err != nil {
		log.Fatalf("Worker %s: Subscribe error: %v", workerID, err)
	}

	fmt.Printf("Worker %s connected and subscribed to shared topic\n", workerID)

	// Wait for interruption
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	fmt.Printf("Worker %s shutting down...\n", workerID)
}
