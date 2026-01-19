package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Config holds the application configuration.
type Config struct {
	Broker   string
	ClientID string
	Mode     string
}

func main() {
	// 1. Parse Command Line Arguments
	mode := flag.String("mode", "qos", "Mode of operation: 'retained' or 'qos'")
	flag.Parse()

	cfg := Config{
		Broker:   "tcp://localhost:1883",
		ClientID: "lab3-publisher",
		Mode:     *mode,
	}

	// 2. Initialize Client
	client := connectToBroker(cfg)
	defer client.Disconnect(250)

	// 3. Execute Mode
	switch cfg.Mode {
	case "retained":
		publishRetainedMessage(client)
	case "qos":
		publishQoSSequence(client)
	default:
		log.Fatalf("Unknown mode: %s. Use 'retained' or 'qos'.", cfg.Mode)
	}
}

// connectToBroker handles the connection logic.
func connectToBroker(cfg Config) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID + "-" + cfg.Mode) // Unique ID per mode to avoid conflicts

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Connection failed: ", token.Error())
	}
	return client
}

// publishRetainedMessage demonstrates the 'Retained' feature.
// Messages with retained=true are stored by the broker and delivered to new subscribers immediately.
func publishRetainedMessage(client mqtt.Client) {
	topic := "lab3/config"
	payload := "System Status: ONLINE (Retained)"
	qos := byte(1)
	retained := true

	fmt.Printf("[Retained] Publishing to %s: %s\n", topic, payload)
	token := client.Publish(topic, qos, retained, payload)
	token.Wait()
	fmt.Println("[Retained] Message sent.")
}

// publishQoSSequence demonstrates different Quality of Service levels.
func publishQoSSequence(client mqtt.Client) {
	topic := "lab3/data"
	
	// Define test cases
	messages := []struct {
		qos     byte
		payload string
		desc    string
	}{
		{0, "Temperature: 20C (QoS 0)", "Fire and Forget"},
		{1, "Temperature: 21C (QoS 1)", "At Least Once"},
		{2, "Temperature: 22C (QoS 2)", "Exactly Once"},
	}

	for _, msg := range messages {
		fmt.Printf("[QoS] Publishing '%s' with QoS %d (%s)\n", msg.payload, msg.qos, msg.desc)
		token := client.Publish(topic, msg.qos, false, msg.payload)
		token.Wait()
		time.Sleep(500 * time.Millisecond)
	}
}


