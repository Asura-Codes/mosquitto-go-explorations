package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("lab5-lwt-client")

	// --- LWT CONFIGURATION ---
	// If this client disconnects ungracefully, the broker will publish this message.
	willTopic := "lab5/status/lwt-client"
	willPayload := "OFFLINE (Crashed/Unexpected)"
	opts.SetWill(willTopic, willPayload, 1, true)
	// --------------------------

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	fmt.Println("LWT Client connected. Publishing 'ONLINE'...")
	client.Publish(willTopic, 1, true, "ONLINE")

	fmt.Println("Simulating a crash in 3 seconds (exiting without Disconnect())...")
	time.Sleep(3 * time.Second)
	
	// We exit WITHOUT calling client.Disconnect(). 
	// The broker will detect the socket closure and trigger the LWT.
	os.Exit(1)
}
