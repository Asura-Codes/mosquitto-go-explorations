package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_publisher_lab11")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer client.Disconnect(250)

	fmt.Println("Go Publisher connected to TCP:1883")
	fmt.Println("Publishing 20 messages to lab11/data...")

	for counter := 1; counter <= 20; counter++ {
		payload := fmt.Sprintf("Sensor update #%d - Temperature: %.2f", counter, 20.0+(float64(counter)*0.1))
		topic := "lab11/data"

		fmt.Printf("Publishing: %s\n", payload)
		client.Publish(topic, 0, false, payload)

		time.Sleep(1 * time.Second)
	}
	fmt.Println("Done.")
}

