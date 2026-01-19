package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// BrokerConfig holds the details for a connection
type BrokerConfig struct {
	Name string
	URL  string
}

func main() {
	brokers := []BrokerConfig{
		{"Primary", "tcp://localhost:1883"},
		{"Secondary", "tcp://localhost:1884"},
		{"Tertiary", "tcp://localhost:1885"},
	}

	clients := make([]mqtt.Client, len(brokers))

	// 1. Connect to ALL brokers simultaneously (Fan-out Publisher)
	for i, b := range brokers {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(b.URL)
		opts.SetClientID(fmt.Sprintf("ha-publisher-%s", b.Name))
		opts.SetCleanSession(true)
		opts.SetAutoReconnect(true)
		opts.SetConnectRetry(true)
		opts.SetConnectRetryInterval(time.Second * 2)

		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Printf("[%s] Connection failed: %v", b.Name, token.Error())
		} else {
			log.Printf("[%s] Connected", b.Name)
		}
		clients[i] = client
	}

	// Publishing Loop
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	go func() {
		counter := 0
		for range ticker.C {
			msg := fmt.Sprintf("Critical Data #%d", counter)
			topic := "critical/metrics"
			
			log.Printf(">> Publishing: %s", msg)

			// Publish to ALL connected brokers
			var wg sync.WaitGroup
			for i, client := range clients {
				if client.IsConnected() {
					wg.Add(1)
					go func(c mqtt.Client, name string) {
						defer wg.Done()
						token := c.Publish(topic, 1, false, msg)
						token.Wait()
						if token.Error() != nil {
							log.Printf("   [%s] Publish Failed: %v", name, token.Error())
						}
					}(client, brokers[i].Name)
				} else {
					log.Printf("   [%s] Disconnected - skipping", brokers[i].Name)
				}
			}
			wg.Wait()
			counter++
		}
	}()

	// Wait for termination
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	for _, c := range clients {
		c.Disconnect(250)
	}
	fmt.Println("Publisher disconnected")
}