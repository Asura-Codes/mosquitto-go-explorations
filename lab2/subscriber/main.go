package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// messageHandler is a callback function that implements the mqtt.MessageHandler type.
// It complies with the Single Responsibility Principle by doing one thing: handling incoming messages.
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("RECEIVED: Topic: %s, Message: %s\n", msg.Topic(), msg.Payload())
}

func main() {
	// 1. Define Broker Configuration
	broker := "tcp://localhost:1883"
	opts := mqtt.NewClientOptions().AddBroker(broker)
	
	// Unique ClientID for the subscriber.
	opts.SetClientID("go-subscriber")
	
	// Set the default handler. This is called if a message arrives on a topic
	// that doesn't have a specific subscription handler.
	opts.SetDefaultPublishHandler(messageHandler)

	// 2. Create and Connect Client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 3. Subscribe to Topics
	// We subscribe to "sensor/#". 
	// The '#' is a multi-level wildcard, meaning it matches "sensor/temp/kitchen", 
	// "sensor/humidity/livingroom", and any other sub-topics.
	topic := "sensor/#"
	token := client.Subscribe(topic, 1, nil) // nil means use the default handler set above
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// 4. Handle Graceful Shutdown
	// We want to keep the application running to receive messages until the user stops it.
	// We listen for OS interrupt signals (Ctrl+C).
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	
	// Block until a signal is received
	<-sig

	// 5. Clean Up
	fmt.Println("\nDisconnecting...")
	client.Disconnect(250)
}

