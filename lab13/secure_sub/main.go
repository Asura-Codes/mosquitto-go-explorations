package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// 1. Load CA Cert
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("Failed to append CA cert to pool")
	}

	// 2. Load Client Cert/Key (Using the same client certs for simplicity, 
	// typically each client might have its own)
	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatalf("Failed to load client keypair: %v", err)
	}

	// 3. Configure TLS
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
		InsecureSkipVerify: false, 
	}

	// 4. Configure MQTT Client
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tls://localhost:8883")
	opts.SetClientID("secure-subscriber")
	opts.SetTLSConfig(tlsConfig)

	// Callback for received messages
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received: [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect: %v", token.Error())
	}
	fmt.Println("Secure Subscriber Connected!")

	// 5. Subscribe
	topic := "secure/#"
	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing: %v", token.Error())
	}
	fmt.Printf("Subscribed to %s\n", topic)

	// Keep running until interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect(250)
	fmt.Println("Subscriber Disconnected")
}
