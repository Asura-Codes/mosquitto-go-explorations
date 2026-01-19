package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("lab5-requester")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer client.Disconnect(250)

	correlationID := "req-123"
	requestTopic := "lab5/request/" + correlationID
	responseTopic := "lab5/response/" + correlationID

	// 1. Subscribe to the response topic FIRST
	replyChan := make(chan string)
	client.Subscribe(responseTopic, 1, func(c mqtt.Client, msg mqtt.Message) {
		replyChan <- string(msg.Payload())
	})

	// 2. Publish the request
	fmt.Printf("[REQUESTER] Sending request to %s...\n", requestTopic)
	client.Publish(requestTopic, 1, false, "Hello Responder!")

	// 3. Wait for reply with timeout
	select {
	case reply := <-replyChan:
		fmt.Printf("[REQUESTER] Received Reply: %s\n", reply)
	case <-time.After(5 * time.Second):
		fmt.Println("[REQUESTER] Timeout waiting for reply.")
	}
}

