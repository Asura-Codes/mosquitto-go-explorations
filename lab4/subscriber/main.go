package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// messageHandler processes incoming messages.
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("--> RECEIVED: %s | Topic: %s | QoS: %d\n", string(msg.Payload()), msg.Topic(), msg.Qos())
}

func main() {
	// 1. Parse CLI Flags
	// CleanSession: If false, the broker stores subscriptions and queues messages while offline.
	cleanSession := flag.Bool("clean", true, "Set CleanSession flag")
	flag.Parse()

	// 2. Setup Client Options
	// CRITICAL: A stable ClientID is required for the broker to recognize a returning persistent session.
	clientID := "lab4-persistent-subscriber"
	broker := "tcp://localhost:1883"
	
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID).
		SetCleanSession(*cleanSession).
		SetDefaultPublishHandler(messageHandler)

	fmt.Printf("Connecting as ClientID: %s (CleanSession: %v)...\n", clientID, *cleanSession)

	// 3. Connect
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Connection failed: ", token.Error())
	}
	defer client.Disconnect(250)

	// 4. Subscribe
	// QoS 1 or 2 is usually required for effective offline queueing.
	topic := "lab4/alerts"
	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		log.Fatal("Subscription failed: ", token.Error())
	}
	fmt.Printf("Subscribed to %s. Waiting for messages...\n", topic)

	// 5. Keep Alive
	waitForSignal()
}

func waitForSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-
sig
	fmt.Println("\nExiting...")
}
