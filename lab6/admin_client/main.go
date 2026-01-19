package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Admin Credentials
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("lab6-admin").
		SetUsername("admin").
		SetPassword("admin123") // In prod, load from env var!

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("[ADMIN] RECEIVED: %s | Payload: %s\n", msg.Topic(), string(msg.Payload()))
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Admin Connection Failed: %v", token.Error())
	}
	fmt.Println("Admin connected successfully.")

	// Admin can subscribe to everything
	if token := client.Subscribe("#", 1, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Admin Subscribe Failed: %v", token.Error())
	}
	fmt.Println("Admin subscribed to '#'")

	// Wait
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}

