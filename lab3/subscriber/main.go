package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// messageHandler is a dedicated function for processing incoming messages.
// It adheres to SRP by separating processing logic from main setup.
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("--> RECEIVED:\n")
	fmt.Printf("    Topic:    %s\n", msg.Topic())
	fmt.Printf("    Payload:  %s\n", string(msg.Payload()))
	fmt.Printf("    QoS:      %d\n", msg.Qos())
	fmt.Printf("    Retained: %v\n", msg.Retained())
	fmt.Println("--------------------------------------------------")
}

func main() {
	// 1. Setup Client Options
	broker := "tcp://localhost:1883"
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID("lab3-subscriber")
	opts.SetDefaultPublishHandler(messageHandler)

	// 2. Connect
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Connection failed: ", token.Error())
	}
	defer client.Disconnect(250)

	// 3. Subscribe
	// We subscribe to 'lab3/#' to catch both 'lab3/config' (retained) and 'lab3/data' (qos).
	// We request QoS 2, but the actual granted QoS depends on the broker and publisher.
	topic := "lab3/#"
	if token := client.Subscribe(topic, 2, nil); token.Wait() && token.Error() != nil {
		log.Fatal("Subscription failed: ", token.Error())
	}
	fmt.Printf("Subscribed to topic: %s (Waiting for messages...)\n", topic)
	fmt.Println("--------------------------------------------------")

	// 4. Keep Alive
	waitForSignal()
}

// waitForSignal blocks until an interrupt signal is received.
func waitForSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	fmt.Println("\nShutdown signal received. Exiting...")
}
