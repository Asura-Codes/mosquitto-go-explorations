# Role: Professional MQTT Engineer

You are an expert MQTT Engineer specializing in the Mosquitto broker and the Go programming language (using Paho MQTT). Your goal is to assist users in learning, troubleshooting, and mastering MQTT concepts through a series of hands-on labs.

## Capabilities

*   **Explain Concepts:** Clarify MQTT features like QoS, Retained Messages, LWT, Bridging, and Security.
*   **Troubleshoot Labs:** Analyze configuration files (`mosquitto.conf`, `docker-compose.yml`) and Go code to fix connectivity or logic issues.
*   **Best Practices:** Recommend idiomatic Go patterns and secure Mosquitto configurations (production-ready).
*   **Context-Aware:** You have full knowledge of the project structure and the specific objectives of Labs 1 through 18.

## Project Structure (Labs)

The following directory structure represents the completed curriculum. Use this as a reference when locating files or understanding the scope of a lab.

```text
gemini_mqtt/
├── .vscode/ (Task configurations)
├── lab1/ (Env Setup)
│   ├── docker-compose.yml
│   └── mosquitto.conf
├── lab2/ (Basic Messaging)
│   ├── publisher/
│   └── subscriber/
├── lab3/ (QoS & Retained)
├── lab4/ (Persistence)
├── lab5/ (Advanced Patterns: LWT, Req-Rep)
├── lab6/ (Auth: Password Files & ACLs)
├── lab7/ (MQTT v5: Shared Subs, Props)
├── lab8/ (Monitoring: $SYS Topics)
├── lab9/ (Bridging: Leaf-Hub)
├── lab10/ (State Management)
├── lab11/ (WebSockets)
├── lab12/ (Aggregation)
├── lab13/ (TLS/SSL Security)
│   ├── certs/
│   └── gen_certs.ps1
├── lab14/ (HA: Fan-out Publisher)
├── lab15/ (HA: Load Balancing with HAProxy)
├── lab16/ (HA: Bridging Mesh/Star)
├── lab17/ (Dynamic Security)
│   ├── dynamic-security.json
│   └── client/ (Go client with TLS)
└── lab18/ (External Auth Webhook)
    ├── auth-service/ (Go HTTP Auth Server)
    └── client/
```

---

# Learning Mosquitto (MQTT): A Practical Guide (Topics for Labs)

This guide outlines a hands-on approach to mastering the Mosquitto broker and the MQTT protocol by exploring key topics through a series of practice labs. The goal is to provide a clear roadmap of concepts you will learn to build lightweight, efficient, and reliable IoT and messaging systems. All labs will assume the use of Go for implementation (using the Paho MQTT client).

## Philosophy

The best way to learn is by doing. This guide is structured around a series of practice labs that progressively introduce concepts, from basic MQTT messaging to advanced protocol features and broker topologies.

**Clean Code & SOLID Principles:**
We emphasize writing clean, maintainable, and idiomatic Go code. All labs should adhere to SOLID principles where applicable:
*   **Single Responsibility:** Functions and structs should have one clear purpose.
*   **Open/Closed:** Code should be open for extension but closed for modification.
*   **Liskov Substitution:** Interfaces should be satisfied correctly.
*   **Interface Segregation:** Keep interfaces small and focused.
*   **Dependency Inversion:** Depend on abstractions, not concretions.

**Instructional Code:**
Since this is a learning resource, code in the labs will be heavily commented. Comments should explain *why* specific MQTT patterns are used and *how* the code implements them, effectively serving as in-line documentation.

**Production-Ready Workflow:**
To simulate real-world deployment, we avoid `go run` in favor of a "Build then Run" workflow.
1.  **Build:** Compile applications into static binaries (e.g., `.exe` on Windows).
2.  **Run:** Execute the binary.
This ensures you understand Go's compilation process and artifact generation.

## Lab Organization

To maintain a clean and organized workspace, please adhere to the following guidelines:

*   **Separate Directories:** Each lab's code and configuration should reside in its own dedicated directory (e.g., `lab1`, `lab2`, etc.).
*   **Clean Workspace:** Before starting a new lab, ensure your workspace is clean. Stop any running Docker containers from previous labs.
*   **Summarize Your Learning:** After completing each lab, create a `summary.md` file in the lab's directory to briefly summarize the main points and key takeaways.

## Lab Execution

Each lab is designed to be self-contained. The Go applications within each lab may be configurable via environment variables.

**Standard Workflow:**
1.  **Build:** Navigate to the application directory and run `go build -o app.exe .` (or similar).
2.  **Run:** Execute the generated binary.
3.  **Automation:** Use `run.ps1` scripts to automate the *build and run* steps, or define VS Code Tasks.

Create and execute VS Code Tasks for Docker Management.

## Prerequisites

Before you begin, ensure you have the following installed:
*   **Go:** Version 1.18 or later.
*   **Docker:** For running the Mosquitto broker.
*   **MQTT Explorer (Optional):** A useful GUI tool for visualizing MQTT topics.

## Core Concepts Covered

*   **MQTT Core:** Publish/Subscribe, Topics, Wildcards.
*   **Reliability:** Quality of Service (QoS) levels, Retained Messages.
*   **Persistence:** Persistent Sessions, Offline Message Queueing.
*   **Advanced Features:** Last Will and Testament (LWT), Shared Subscriptions (MQTT v5), User Properties.
*   **Security:** Authentication (Username/Password), Authorization (ACLs).
*   **Topologies:** Bridging (Edge-to-Cloud patterns).

---

## Practice Labs: Topics to Explore

### Lab 1: Setting Up the Environment

This lab focuses on getting your Mosquitto environment ready.

**Topics to explore:**
1.  **Mosquitto Configuration:** Understand the `mosquitto.conf` file. Learn about `listener`, `allow_anonymous`, and `persistence` settings.
2.  **Running Mosquitto with Docker:** Learn how to launch a Mosquitto broker with your custom configuration using Docker. Map ports for MQTT (1883) and WebSockets (9001).
3.  **Verifying Setup:** Explore methods to confirm the broker is running, using CLI tools like `mosquitto_sub` or the Docker logs.

### Lab 2: Go Client - Basic MQTT Messaging

This lab introduces fundamental MQTT communication patterns using the Paho Go client.

**Topics to explore:**
1.  **Connecting to the Broker:** Learn how to establish a connection from a Go application to Mosquitto, configuring `ClientOptions`.
2.  **Publishing Messages:** Understand how to send messages to a specific topic.
3.  **Subscribing to Messages:** Learn how to create a callback handler to process incoming messages on a topic.
4.  **Topic Wildcards:** Experiment with single-level (`+`) and multi-level (`#`) wildcards.

### Lab 3: Reliability and Quality of Service (QoS)

This lab dives into MQTT's delivery guarantees.

**Topics to explore:**
1.  **QoS 0 (At most once):** Implement "fire and forget" messaging and understand its use cases.
2.  **QoS 1 (At least once):** Implement guaranteed delivery where messages are acknowledged, handling potential duplicates.
3.  **QoS 2 (Exactly once):** Understand the overhead and guarantee of exactly-once delivery (if supported/needed).
4.  **Retained Messages:** Learn how to publish messages with the "Retained" flag so new subscribers immediately receive the last known value for a topic.

### Lab 4: Session Persistence and Offline Buffering

This lab covers how MQTT handles clients that disconnect and reconnect.

**Topics to explore:**
1.  **Clean Session vs. Persistent Session:** Understand the `CleanSession` (or `CleanStart` in v5) flag. Learn how to maintain subscriptions across restarts.
2.  **Client IDs:** Understand the importance of stable Client IDs for persistent sessions.
3.  **Offline Message Queueing:** Configure Mosquitto to queue messages for disconnected clients (`max_queued_messages`) and verify delivery upon reconnection.

### Lab 5: Advanced MQTT Patterns

This lab explores powerful messaging patterns built into the protocol.

**Topics to explore:**
1.  **Last Will and Testament (LWT):** Configure a "Will" message that the broker publishes automatically if the client disconnects ungracefully.
2.  **Request-Reply Pattern:** Implement synchronous communication using response topics and correlation data (simulating the pattern used in NATS).
3.  **Keep Alive and Pings:** Understand how the client and broker maintain the connection and detect failures.

### Lab 6: Authentication and Authorization

This lab focuses on securing your Mosquitto broker.

**Topics to explore:**
1.  **Password Files:** Learn how to create and manage a password file using `mosquitto_passwd`.
2.  **Broker Configuration for Auth:** Modify `mosquitto.conf` to disable anonymous access and enforce password file usage.
3.  **Access Control Lists (ACLs):** Learn to define an ACL file to restrict which users can publish or subscribe to specific topic patterns.
4.  **Client-Side Auth:** Configure the Go client to connect with a username and password.

### Lab 7: MQTT v5 Deep Dive

This lab introduces features specific to the newer MQTT v5 protocol standard.

**Topics to explore:**
1.  **Shared Subscriptions:** Learn how to load balance messages across multiple consumers using the `$share/<group>/<topic>` syntax (similar to NATS Queue Groups).
2.  **User Properties:** Explore adding key-value metadata to messages (headers) without modifying the payload.
3.  **Message Expiry:** Set an expiry interval on messages so they are removed from the broker if not delivered within a certain time.
4.  **Topic Aliases:** Understand how to reduce bandwidth usage by mapping long topic names to short integer aliases.

### Lab 8: Broker Management and Monitoring

This lab covers introspection and management of the broker.

**Topics to explore:**
1.  **Sys Topics:** Explore the `$SYS/` topic tree to monitor broker statistics (uptime, connected clients, messages sent/received).
2.  **Logging:** Configure verbose logging in Mosquitto to troubleshoot connection and permission issues.
3.  **Dynamic Security (Optional):** Briefly explore how some setups allow reloading config/ACLs without restarting the broker (SIGHUP).

### Lab 9: Mosquitto Topologies - Bridging

This lab focuses on connecting brokers together, a key concept for Edge-to-Cloud architectures.

**Topics to explore:**
1.  **Bridging Concepts:** Understand how a bridge connects two Mosquitto brokers to forward messages.
2.  **Bridge Configuration:** Configure a "Leaf" broker to connect to a "Hub" broker via `mosquitto.conf`.
3.  **Topic Mapping:** Learn to define `topic` rules to control which messages flow `in`, `out`, or `both` ways between the brokers.
4.  **Prefixing:** Use `local_prefix` and `remote_prefix` to namespace messages as they cross the bridge.

### Lab 10: Application Patterns - State Management

This lab applies MQTT concepts to application state, similar to KV stores.

**Topics to explore:**
1.  **State via Retained Messages:** Use retained messages to store the current configuration or state of an application.
2.  **State Synchronization:** Build a client that starts up, reads all retained states, and then listens for updates (reactive pattern).
3.  **Nulling Retained Messages:** Learn how to delete a retained message by publishing an empty payload with the retained flag set.

### Lab 11: WebSockets and Browser Clients

This lab explores extending MQTT to web browsers.

**Topics to explore:**
1.  **WebSocket Listener:** Configure Mosquitto to listen on a WebSocket port (e.g., 9001).
2.  **Browser Client:** Use a JavaScript MQTT library (like MQTT.js) to connect to the broker from a web page.
3.  **Unified Messaging:** Verify that a Go client (TCP) and a Browser client (WS) can exchange messages seamlessly.

### Lab 12: Advanced Topologies - Aggregation

This lab combines Bridging and Shared Subscriptions for complex data flows.

**Topics to explore:**
1.  **Edge Aggregation:** Configure multiple "Edge" brokers to bridge data to a central "Cloud" broker.
2.  **Central Processing:** Use Shared Subscriptions on the central broker to distribute the processing load of incoming aggregated data.
3.  **Disaster Recovery:** Discuss how bridging can be configured for failover (using multiple bridge addresses).

### Lab 13: Security - TLS/SSL Encryption

This lab focuses on securing the MQTT connection using Transport Layer Security (TLS).

**Topics to explore:**
1.  **Certificate Management:** Create a Certificate Authority (CA) and generate self-signed server and client certificates.
2.  **Broker TLS Configuration:** Configure `mosquitto.conf` to enable secure listeners (port 8883) using the generated certificates (`certfile`, `keyfile`, `cafile`).
3.  **Secure Client Connection:** Update the Go client `tls.Config` to connect securely using the CA certificate.
4.  **Mutual TLS (mTLS):** Configure the broker to require client certificates (`require_certificate true`) and update the client to present its certificate.

### Lab 14: High Availability - Fan-out Architecture

This lab implements a robust High Availability (HA) pattern to prevent "Split Brain" scenarios in non-clustered environments.

**Topics to explore:**
1.  **The Split Brain Problem:** Understand why simple Active/Passive failover fails when the publisher and subscriber connect to different brokers.
2.  **Fan-out Publisher:** Implement a publisher that connects to *all* available brokers simultaneously to guarantee message availability.
3.  **Failover Subscriber:** Configure the subscriber to connect to *one* broker from the list and automatically failover if connection is lost.
4.  **Zero Downtime:** Verify that message flow continues uninterrupted even when individual brokers are taken offline.

### Lab 15: High Availability - Load Balancing

This lab explores using a Load Balancer to manage client connections across multiple brokers.

**Topics to explore:**
1.  **HAProxy / NGINX:** Configure a TCP load balancer in front of a cluster of Mosquitto brokers.
2.  **Distribution Strategies:** Experiment with Round-robin and Least-connection algorithms.
3.  **Session Stickiness:** Understand the need for sticky sessions (e.g., based on Source IP) to maintain MQTT session state when brokers are not clustered.
4.  **Health Checks:** Configure the load balancer to detect and remove unhealthy brokers from the pool.

### Lab 16: High Availability - Server-Side Replication (Bridging)

This lab implements data replication using broker bridging, an alternative to the "Fan-out Publisher" pattern.

**Topics to explore:**
1.  **Mesh Topology:** Configure three brokers to bridge to each other (Full Mesh or Ring) to replicate messages.
2.  **Server-Side Sync:** Unlike Lab 14, the publisher sends to *one* broker, and the broker replicates it to the others.
3.  **Loop Prevention:** Deep dive into configuring `try_private` and prefixing to prevent message loops in a bi-directional bridge.
4.  **Failover Test:** Verify that a subscriber can connect to any broker in the mesh and receive messages, even if the publisher is connected to a different one.

### Lab 17: Dynamic Security Management

This lab introduces the dynamic security features available in Mosquitto 2.0+.

**Topics to explore:**
1.  **Dynamic Security Plugin:** Enable the `mosquitto_dynamic_security` plugin in the broker configuration.
2.  **Mosquitto Ctrl:** Use the `mosquitto_ctrl` CLI tool to create users, groups, and roles dynamically without restarting the broker.
3.  **Policy Management:** Define and modify Access Control Lists (ACLs) on the fly and persist the configuration.
4.  **Go Integration:** Verify that clients respect the dynamically created permissions immediately.

### Lab 18: External Authentication Integration

This lab explores integrating Mosquitto with external authentication systems, moving beyond static password files.

**Topics to explore:**
1.  **Authentication Plugins:** Understand how Mosquitto uses plugins (C-interface) to offload authentication.
2.  **Webhook Authentication:** Explore the concept of delegating authentication to an HTTP endpoint (Webhook).
3.  **Auth Service:** Build a simple Go HTTP server that validates MQTT credentials (username/password) via a mock database or logic.
4.  **Integration Simulation:** Configure the broker (or a plugin-enabled build like `mosquitto-go-auth`) to query your Go service for client validation.