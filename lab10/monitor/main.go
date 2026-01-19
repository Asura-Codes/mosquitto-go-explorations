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
	brokerAddr := "localhost:1883"
	if addr := os.Getenv("MQTT_BROKER"); addr != "" {
		brokerAddr = addr
	}

	conn, err := net.Dial("tcp", brokerAddr)
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}

	// Channel to receive messages
	msgChan := make(chan *paho.Publish)

	client := paho.NewClient(paho.ClientConfig{
		Conn: conn,
		OnPublishReceived: []func(paho.PublishReceived) (bool, error){
			func(pr paho.PublishReceived) (bool, error) {
				msgChan <- pr.Packet
				return true, nil
			},
		},
	})

	cp := &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		ClientID:   "state_monitor",
	}

	_, err = client.Connect(context.Background(), cp)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	fmt.Println("--- Monitor Connected. Subscribing to '#' to catch all states ---")
	
	// Subscribe to wildcards to see all current retained states
	_, err = client.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: []paho.SubscribeOptions{
			{Topic: "settings/#", QoS: 1},
			{Topic: "status/#", QoS: 1},
		},
	})
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case msg := <-msgChan:
			retainedStr := ""
			if msg.Retain {
				retainedStr = "[RETAINED] "
			}
			fmt.Printf("Recv: %s%s = %s\n", retainedStr, msg.Topic, string(msg.Payload))
		case <-sigChan:
			fmt.Println("\nMonitor shutting down...")
			client.Disconnect(&paho.Disconnect{ReasonCode: 0})
			return
		}
	}
}
