package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Main function serves as the entry point for the Publisher application.
// In a production app, this logic would likely be encapsulated in a 'Publisher' struct/service
// to adhere to Single Responsibility Principle (SRP).
func main() {
	// 1. Define Broker Configuration
	// We connect to localhost on the standard MQTT port (1883).
	broker := "tcp://localhost:1883"
	
	// Create ClientOptions struct.
	// This uses the Builder pattern to configure the client.
	opts := mqtt.NewClientOptions().AddBroker(broker)
	
	// Set a ClientID. 
	// Important: ClientIDs should be unique per client. If two clients connect with the same ID,
	// the broker will disconnect the older one.
	opts.SetClientID("go-publisher")

	// 2. Create and Start the Client
	client := mqtt.NewClient(opts)
	
	// Connect to the broker.
	// Connect() is asynchronous and returns a Token. We call Wait() to block until completion.
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 3. Define Messages
	// We define a slice of anonymous structs to represent our data.
	// This simulates reading from different sensors.
	messages := []struct {
		topic   string
		payload string
	}{
		{"sensor/temp/livingroom", "22.5"},       // Matches sensor/#
		{"sensor/temp/kitchen", "25.0"},          // Matches sensor/#
		{"sensor/humidity/livingroom", "45.2"},   // Matches sensor/#
		{"other/topic", "not a sensor"},          // Does NOT match sensor/#
	}

	// 4. Publish Messages
	for _, m := range messages {
		fmt.Printf("PUBLISHING: Topic: %s, Message: %s\n", m.topic, m.payload)
		
		// Publish() sends the message to the broker.
		// QoS 1 (At Least Once) is used here to ensure delivery.
		// Retained = false: New subscribers won't see this message immediately upon connecting; 
		// they only see new messages sent after they subscribe.
		token := client.Publish(m.topic, 1, false, m.payload)
		
		// Wait for the publish to be acknowledged by the broker.
		token.Wait()
		
		// Simulate a delay between sensor readings.
		time.Sleep(500 * time.Millisecond)
	}

	// 5. Clean Disconnect
	// Allow 250ms for pending work to complete before closing the connection.
	client.Disconnect(250)
}
