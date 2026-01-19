# Lab 5 Summary: Advanced MQTT Patterns

## Key Takeaways

### 1. Last Will and Testament (LWT)
- **Purpose:** Allows other clients to know when a client has disconnected ungracefully (e.g., crash, network loss).
- **Mechanism:** The client provides a "Will" (topic and payload) to the broker during the initial connection. If the broker detects the client's socket closed without a `DISCONNECT` packet, it publishes the Will message.

### 2. Request-Reply Pattern
- **Purpose:** Enables synchronous-style communication over an asynchronous protocol.
- **Mechanism:**
    - The **Requester** subscribes to a unique response topic.
    - The **Requester** publishes a request.
    - The **Responder** processes and replies to the specified response topic.

## Execution Example

Running `run.ps1` demonstrates both patterns through the `monitor`:

1. **LWT Demo:**
   ```text
   LWT Client: Connecting with Will [Topic: lab5/status/lwt-client, Payload: OFFLINE]
   LWT Client: Crashing now (os.Exit)...
   Monitor: [lab5/status/lwt-client] -> OFFLINE
   ```
   *Note: The broker published the "OFFLINE" message because it detected the ungraceful exit.*

2. **Request-Reply Demo:**
   ```text
   Requester: Sending request to 'lab5/request/req-123'
   Responder: Received request for 'req-123', replying...
   Requester: Received reply: [REPLY to req-123] SUCCESS
   ```

### 3. Keep Alive and Pings
- **Concept:** MQTT uses a "Keep Alive" interval (seconds). If the broker doesn't receive a packet (message or ping) within 1.5x the interval, it assumes the connection is lost.
- **Implementation:** Paho Go client handles these background pings automatically based on the `SetKeepAlive` option.
