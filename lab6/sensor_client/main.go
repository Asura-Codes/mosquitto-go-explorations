package main

import (
	"fmt"
	"log"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Sensor Credentials
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("lab6-sensor").
		SetUsername("sensor").
		SetPassword("sensor123")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Sensor Connection Failed: %v", token.Error())
	}
	fmt.Println("Sensor connected successfully.")

	// 1. Allowed Action
	topicAllowed := "sensors/temp"
	fmt.Printf("[SENSOR] Publishing to Allowed Topic: %s\n", topicAllowed)
	token := client.Publish(topicAllowed, 1, false, "25.5C")
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("ERROR: %v\n", token.Error())
	} else {
		fmt.Println("SUCCESS: Message sent.")
	}

	// 2. Denied Action
	// Note: Mosquitto often accepts the publish packet but silently drops it 
	// or disconnects the client if configured, depending on protocol version and settings.
	// With MQTT v3.1.1, the client might not get an immediate error token unless the connection is dropped.
	topicDenied := "admin/secret"
	fmt.Printf("[SENSOR] Publishing to Denied Topic: %s\n", topicDenied)
	token2 := client.Publish(topicDenied, 1, false, "I shouldn't be here")
	token2.Wait()
	if token2.Error() != nil {
		fmt.Printf("ERROR: %v\n", token2.Error())
	} else {
		fmt.Println("INFO: Message sent (Broker may silently drop it due to ACL).")
	}

	time.Sleep(1 * time.Second)
	client.Disconnect(250)
}
