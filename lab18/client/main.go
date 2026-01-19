package main

import (
	"crypto/tls"

	"crypto/x509"

	"flag"

	"fmt"

	"log"

	"os"

	"os/signal"

	"syscall"

	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (

	broker   = "tcp://localhost:1883"

	topicPtr = flag.String("topic", "test/topic", "MQTT Topic")

	clientID = "go-client-lab18"

)



func main() {

	username := flag.String("user", "", "MQTT Username")

	password := flag.String("pass", "", "MQTT Password")

	action := flag.String("action", "sub", "Action: sub, pub, or both")

	caFile := flag.String("ca", "", "Path to CA certificate (optional)")

	flag.Parse()



	if *username == "" {

		fmt.Println("Warning: No username provided. Connection might fail if auth is required.")

	}



	opts := mqtt.NewClientOptions()

	opts.AddBroker(broker)

	

	// TLS Configuration (Optional)

	if *caFile != "" {

		certpool := x509.NewCertPool()

		ca, err := os.ReadFile(*caFile)

		if err != nil {

			log.Printf("Warning: Failed to read CA file '%s': %v. Proceeding without TLS config.", *caFile, err)

		} else {

			certpool.AppendCertsFromPEM(ca)

			tlsConfig := &tls.Config{

				RootCAs: certpool,

				ServerName: "localhost",

			}

			opts.SetTLSConfig(tlsConfig)

			opts.AddBroker("ssl://localhost:8883")

		}

	}



	// Ensure unique ClientID even if same user is used twice

	opts.SetClientID(fmt.Sprintf("%s-%s-%d", clientID, *username, time.Now().UnixNano()))

	opts.SetUsername(*username)

	opts.SetPassword(*password)

	

	// Default Handler

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {

		fmt.Printf("[%s] Received on %s: %s\n", *username, msg.Topic(), string(msg.Payload()))

	})



	opts.SetOnConnectHandler(func(client mqtt.Client) {

		fmt.Printf("[%s] Connected to %s\n", *username, broker)

	})



	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {

		fmt.Printf("[%s] Connection Lost: %v\n", *username, err)

	})



	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {

		log.Fatalf("[%s] Failed to connect: %v", *username, token.Error())

	}

	defer client.Disconnect(250)



	if *action == "sub" || *action == "both" {

		if token := client.Subscribe(*topicPtr, 1, nil); token.Wait() && token.Error() != nil {

			fmt.Printf("[%s] Subscribe Error: %v\n", *username, token.Error())

		} else {

			fmt.Printf("[%s] Subscribed to %s\n", *username, *topicPtr)

		}

	}



	if *action == "pub" || *action == "both" {

		go func() {

			for {

				msg := fmt.Sprintf("Hello from %s at %s", *username, time.Now().Format(time.RFC3339))

				token := client.Publish(*topicPtr, 1, false, msg)

				token.Wait()

				if token.Error() != nil {

					fmt.Printf("[%s] Publish Error: %v\n", *username, token.Error())

				} else {

					fmt.Printf("[%s] Published to %s: %s\n", *username, *topicPtr, msg)

				}

				time.Sleep(2 * time.Second)

			}

		}()

	}

	// Keep running until interrupt
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	fmt.Println("Exiting...")
}
