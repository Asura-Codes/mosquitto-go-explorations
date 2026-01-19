# Lab 7 Summary: MQTT v5 Deep Dive

In this lab, I explored features introduced in MQTT v5 using the `paho.golang` client.

## Key Takeaways

1.  **Shared Subscriptions**:
    -   Used the `$share/<group>/<topic>` syntax.
    -   Demonstrated how the broker load-balances messages across multiple subscribers in the same group.
    -   This is ideal for scaling out message processing (similar to NATS Queue Groups).

2.  **User Properties**:
    -   Added custom metadata to messages using `UserProperty` (key-value pairs).
    -   This allows passing application-level headers without modifying the binary payload.

3.  **Message Expiry**:
    -   Set `MessageExpiryInterval` on published messages.
    -   The broker will discard the message if it cannot be delivered within the specified time, preventing stale data from accumulating in offline queues.

## Execution Example

Running `run.ps1` starts two shared workers (A and B) and a publisher that sends 10 messages.

**Publisher:**
```text
Publishing 10 messages to 'v5/test/topic' with UserProperties...
Done.
```

**Worker A:**
```text
[Worker A] Received message 1 | UserProperty [priority: high]
[Worker A] Received message 3 | UserProperty [priority: high]
...
```

**Worker B:**
```text
[Worker B] Received message 2 | UserProperty [priority: high]
[Worker B] Received message 4 | UserProperty [priority: high]
...
```

*Note: Each message is delivered to only ONE worker in the shared group, demonstrating round-robin load balancing.*

4.  **MQTT v5 Protocol**:
    -   The `paho.golang` library provides a cleaner API for v5 features compared to the older Paho client.
    -   Explicit `Connect`, `Publish`, and `Subscribe` packets allow setting v5-specific properties.
