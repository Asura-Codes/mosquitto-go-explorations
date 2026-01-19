package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	broker := "tcp://localhost:1883"
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("lab4-publisher")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer client.Disconnect(250)

	topic := "lab4/alerts"
	for i := 1; i <= 5; i++ {
		payload := fmt.Sprintf("Alert #%d", i)
		fmt.Printf("Publishing: %s\n", payload)
		// Use QoS 1 to ensure the broker queues it for persistent subscribers.
		token := client.Publish(topic, 1, false, payload)
		token.Wait()
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("All messages published.")
}
