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
	if envAddr := os.Getenv("MQTT_BROKER"); envAddr != "" {
		brokerAddr = envAddr
	}

	conn, err := net.Dial("tcp", brokerAddr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Router to handle different $SYS topics
	router := paho.NewSingleHandlerRouter(func(m *paho.Publish) {
		fmt.Printf("[SYS] %-40s | %s\n", m.Topic, string(m.Payload))
	})

	client := paho.NewClient(paho.ClientConfig{
		Conn:   conn,
		Router: router,
	})

	_, err = client.Connect(context.Background(), &paho.Connect{
		KeepAlive:  30,
		CleanStart: true,
		ClientID:   "lab8-monitor",
	})
	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}
	fmt.Println("Monitor connected. Subscribing to $SYS topics...")

	// Subscribing to ALL $SYS topics using a wildcard
	// This reveals every available metric exposed by the broker.
	topics := []string{
		"$SYS/#",
	}

	var subs []paho.SubscribeOptions
	for _, t := range topics {
		subs = append(subs, paho.SubscribeOptions{Topic: t, QoS: 0})
	}

	_, err = client.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: subs,
	})
	if err != nil {
		log.Fatalf("Subscribe error: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	fmt.Println("Monitor shutting down...")
}
