# Lab 8 Summary: Broker Management and Monitoring

In this lab, I explored how to monitor and manage the Mosquitto broker using `$SYS` topics and logging.

## Key Takeaways

1.  **$SYS Topics**:
    -   Mosquitto publishes internal status information to the `$SYS/` topic tree.
    -   Topics like `$SYS/broker/uptime`, `$SYS/broker/clients/connected`, and `$SYS/broker/messages/received` provide real-time insights into the broker's health and load.
    -   These topics are invaluable for building monitoring dashboards.

2.  **Broker Configuration for Monitoring**:
    -   `sys_interval`: Configured how often (in seconds) the broker updates the `$SYS` topics.
    -   `log_type`: Enabled various levels of logging (`information`, `notice`, `warning`, `error`) to see what the broker is doing internally.

3.  **Real-time Introspection**:
    -   Built a Go client that subscribes to `$SYS` topics and prints the broker's state.
    -   Observed how message counts and client connections update dynamically.

## Execution Example

Running `run.ps1` starts a monitor that listens to the internal broker state:

```text
--- Starting Monitor ---
Connected to Broker. Subscribing to $SYS/# ...
Recv: $SYS/broker/version = mosquitto version 2.0.18
Recv: $SYS/broker/uptime = 5 seconds
Recv: $SYS/broker/clients/total = 1
Recv: $SYS/broker/messages/received = 0

--- Generating Activity ---
(mosquitto_pub sends a message...)

Recv: $SYS/broker/messages/received = 1
Recv: $SYS/broker/publish/messages/received = 1
```

*Note: The `$SYS` hierarchy is unique to each broker and is the standard way to monitor MQTT infrastructure.*

4.  **Logging**:
    -   Verified that the broker logs connection events, subscriptions, and other management information to `stdout`, which can be viewed via `docker logs`.
