package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("lab5-responder")
	
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Pattern: Request-Reply
	// We listen for requests on "lab5/request/+". The last part is the Correlation ID.
	client.Subscribe("lab5/request/+", 1, func(c mqtt.Client, msg mqtt.Message) {
		parts := strings.Split(msg.Topic(), "/")
		correlationID := parts[len(parts)-1]
		
		fmt.Printf("[RESPONDER] Received request [%s]: %s\n", correlationID, string(msg.Payload()))
		
		// In MQTT v3, we manually construct the response topic.
		responseTopic := "lab5/response/" + correlationID
		reply := "ACK: " + string(msg.Payload())
		
		fmt.Printf("[RESPONDER] Sending reply to %s\n", responseTopic)
		c.Publish(responseTopic, 1, false, reply)
	})

	fmt.Println("Responder active. Waiting for requests...")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
