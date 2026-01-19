package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
)

func main() {
	// Parse CLI flags
	brokerAddr := flag.String("broker", "localhost:1884", "Broker address (e.g., localhost:1884)")
	topicPrefix := flag.String("prefix", "edge1", "Topic prefix (e.g., edge1)")
	clientID := flag.String("id", "pub-edge1", "Client ID")
	flag.Parse()

	ctx := context.Background()

	// Channel to signal when connected
	connected := make(chan struct{})

	cfg := autopaho.ClientConfig{
		ServerUrls: []*url.URL{
			{Scheme: "tcp", Host: *brokerAddr},
		},
		KeepAlive: 20,
		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
			fmt.Printf("Connected to %s as %s\n", *brokerAddr, *clientID)
			// Non-blocking send to avoid deadlocks if called multiple times (reconnects)
			select {
			case connected <- struct{}{}:
			default:
			}
		},
		ClientConfig: paho.ClientConfig{
			ClientID: *clientID,
		},
	}

	cm, err := autopaho.NewConnection(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for connection
	fmt.Println("Waiting for connection...")
	<-connected

	fmt.Printf("Publishing 5 messages to '%s/data'\n", *topicPrefix)

	for i := 1; i <= 5; i++ {
		payload := fmt.Sprintf("Data from %s: %d", *clientID, i)
		topic := fmt.Sprintf("%s/data", *topicPrefix)

		_, err := cm.Publish(ctx, &paho.Publish{
			Topic:   topic,
			QoS:     1,
			Payload: []byte(payload),
		})

		if err != nil {
			log.Printf("Error publishing: %s\n", err)
		} else {
			fmt.Printf("Sent: %s\n", payload)
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Done. Disconnecting...")
	cm.Disconnect(ctx)
}
