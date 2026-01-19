package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// 1. Load the Certificate Authority (CA) certificate
	// This is used to verify the broker's identity.
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("Failed to append CA cert to pool")
	}

	// 2. Load the Client Certificate and Key
	// This is used for Mutual TLS (mTLS) so the broker can verify our identity.
	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatalf("Failed to load client keypair: %v", err)
	}

	// 3. Configure TLS
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
		InsecureSkipVerify: false,
		// Note: Since we are running locally and the certs likely use "localhost" or similar, 
		// validation should pass if generated correctly.
	}

	// 4. Configure MQTT Client
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tls://localhost:8883") // Note 'tls://' scheme and port 8883
	opts.SetClientID("secure-publisher")
	opts.SetTLSConfig(tlsConfig)
	
	// Define a simple callback for connection lost
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect: %v", token.Error())
	}
	fmt.Println("Secure Publisher Connected!")

	// 5. Publish Messages
	topic := "secure/updates"
	for i := 1; i <= 5; i++ {
		text := fmt.Sprintf("Secure Message #%d", i)
		token := client.Publish(topic, 1, false, text)
		token.Wait()
		fmt.Printf("Published: %s\n", text)
		time.Sleep(1 * time.Second)
	}

	client.Disconnect(250)
	fmt.Println("Publisher Disconnected")
}
