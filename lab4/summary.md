# Lab 4 Summary: Session Persistence and Offline Buffering

## Key Takeaways

### 1. Persistent Sessions (`CleanSession = false`)
- When `CleanSession` is set to `false`, the broker maintains the client's state even after they disconnect.
- State includes the client's subscriptions and missed QoS > 0 messages.
- **Client ID Requirement:** The `ClientID` must be **stable** (identical between sessions).

### 2. Offline Message Queueing
- **Mechanism:** If a persistent client is offline, the broker queues incoming messages for its subscribed topics (provided QoS > 0).
- **Delivery:** Upon reconnection with the same `ClientID` and `CleanSession=false`, the broker flushes the queue.

## Execution Example

Running `run.ps1` follows a timeline:

1. **Subscriber connects & disconnects:** Establishes a session on the broker with a subscription to `lab4/alerts`.
2. **Publisher sends data:** 
   ```text
   Publishing 5 messages to 'lab4/alerts' (Subscriber is currently OFFLINE)
   ```
3. **Subscriber reconnects:**
   ```text
   Subscriber: Connected. (SessionPresent: true)
   Recv: [lab4/alerts] Alert #1 (QUEUED)
   Recv: [lab4/alerts] Alert #2 (QUEUED)
   ...
   ```
   *Note: The `SessionPresent: true` flag in the CONNACK response confirms the broker remembered the client.*

## Conclusion
Persistent sessions are vital for unreliable network conditions (e.g., mobile or IoT sensors) where data loss during transient disconnections must be avoided.
