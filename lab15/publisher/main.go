package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// 1. Create Client Options
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883") // Connect to HAProxy
	opts.SetClientID("go_publisher")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	// 2. Create Client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Publisher Connected to HAProxy (localhost:1883)")

	// 3. Publish Messages
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	count := 0
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			select {
			case <-ticker.C:
				msg := fmt.Sprintf("Message %d from Publisher", count)
				token := client.Publish("test/topic", 0, false, msg)
				token.Wait()
				fmt.Printf("Published: %s\n", msg)
				count++
			case <-c:
				return
			}
		}
	}()

	<-c
	client.Disconnect(250)
	fmt.Println("\nDisconnected")
}
