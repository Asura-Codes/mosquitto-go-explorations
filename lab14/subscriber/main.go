package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// MQTT Client Options
	opts := mqtt.NewClientOptions()

	// High Availability: Add ALL brokers to the failover list
	// The client will connect to ONE of these. If it fails, it tries the next.
	opts.AddBroker("tcp://localhost:1883") // Primary
	opts.AddBroker("tcp://localhost:1884") // Secondary
	opts.AddBroker("tcp://localhost:1885") // Tertiary

	opts.SetClientID("ha-subscriber")
	opts.SetCleanSession(false) 
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(time.Second * 2)
	opts.SetKeepAlive(time.Second * 5)

	// Connection Callbacks
	opts.OnConnect = func(client mqtt.Client) {
		log.Println("Subscriber CONNECTED.")
		
		// Re-subscribe heavily relies on the fact that the Publisher is sending 
		// to ALL brokers. So wherever we land, the data will be there.
		topic := "critical/metrics"
		token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
			// We might get duplicates if we somehow flap between brokers rapidly, 
			// but for failover, we just want to ensure we get 'at least once'.
			fmt.Printf("<< RECEIVED: %s from %s\n", msg.Payload(), msg.Topic())
		})
		token.Wait()
		if token.Error() != nil {
			log.Printf("Error subscribing: %v", token.Error())
		} else {
			log.Printf("Subscribed to %s", topic)
		}
	}

	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Printf("Subscriber Connection LOST: %v", err)
	}

	opts.OnReconnecting = func(client mqtt.Client, opts *mqtt.ClientOptions) {
		log.Println("Subscriber RECONNECTING (Searching for active broker)...")
	}

	// Create Client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting: %v", token.Error())
	}

	// Wait for termination
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-
sigChan

	client.Disconnect(250)
	fmt.Println("Subscriber disconnected")
}