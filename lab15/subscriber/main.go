package main

import (
	"fmt"
	os "os"
	signal "os/signal"
	"syscall"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func main() {
	// 1. Create Client Options
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883") // Connect to HAProxy
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	// 2. Create Client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Subscriber Connected to HAProxy (localhost:1883)")

	// 3. Subscribe
	if token := client.Subscribe("test/topic", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("Subscribed to 'test/topic'")

	// 4. Wait for Signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	client.Disconnect(250)
	fmt.Println("\nDisconnected")
}
