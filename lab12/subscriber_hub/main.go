package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	syscall "syscall"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get Worker ID from env or default
	workerID := os.Getenv("WORKER_ID")
	if workerID == "" {
		workerID = "unknown"
	}

	cfg := autopaho.ClientConfig{
		ServerUrls: []*url.URL{
			{Scheme: "tcp", Host: "localhost:1883"}, // Connect to Hub
		},
		KeepAlive: 20,
		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
			fmt.Printf("[Worker %s] Connected to Hub.\n", workerID)

			// Shared Subscription: $share/<group>/<topic>
			// Group: processors
			// Topic: aggregated/#
			_, err := cm.Subscribe(context.Background(), &paho.Subscribe{
				Subscriptions: []paho.SubscribeOptions{
					{Topic: "$share/processors/aggregated/#", QoS: 1},
				},
			})
			if err != nil {
				fmt.Printf("Error subscribing: %s\n", err)
			} else {
				fmt.Printf("[Worker %s] Subscribed to shared topic '$share/processors/aggregated/#'\n", workerID)
			}
		},
		ClientConfig: paho.ClientConfig{
			ClientID: "hub_worker_" + workerID,
			Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
				fmt.Printf("[Worker %s] Processing: %s | Payload: %s\n", workerID, m.Topic, string(m.Payload))
			}),
		},
	}

	cm, err := autopaho.NewConnection(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for interrupt signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	fmt.Println("Disconnecting...")
	cm.Disconnect(context.Background())
}
