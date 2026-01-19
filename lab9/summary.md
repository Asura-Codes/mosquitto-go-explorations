# Lab 9: Mosquitto Topologies - Bridging

## Summary
In this lab, we explored how to connect two Mosquitto brokers using a **Bridge**. This is a fundamental pattern for Edge-to-Cloud architectures, where edge devices publish to a local "Leaf" broker, which then forwards selected messages to a central "Hub" broker.

## Key Concepts Learned

### 1. Bridge Configuration
We configured the "Leaf" broker to initiate a connection to the "Hub" broker using the `connection` directive in `mosquitto.conf`.
```conf
connection edge-to-cloud
address mosquitto-hub:1883
```

### 2. Topic Mapping & Remapping
We used the `topic` directive to control which messages cross the bridge and how they are renamed.
```conf
# Syntax: topic <pattern> <direction> <qos> <local-prefix> <remote-prefix>
topic # out 1 leaf_sensor/ hub_sensor/
```
*   **Pattern `#`**: Matches all topics (relative to the prefix).
*   **Direction `out`**: Messages flow from Leaf -> Hub.
*   **Remapping**: Messages published to `leaf_sensor/foo` locally are republished as `hub_sensor/foo` remotely.

## Execution Example

When running `run.ps1`, the output demonstrates the bridge forwarding and remapping:

1. **Publisher (Leaf):**
   ```text
   Connected to Leaf Broker (1884)
   Publishing to leaf_sensor/temperature: 25.5
   ```

2. **Subscriber (Hub):**
   ```text
   Connected to Hub Broker (1883) Subscribing to hub_sensor/#
   [Hub Recv] Topic: hub_sensor/temperature | Payload: 25.5
   ```
   *Note: Notice how the prefix changed from `leaf_sensor/` to `hub_sensor/` automatically due to the bridge configuration.*

## Architecture
*   **Leaf Broker (Edge):** buffer messages locally if the connection is lost (Quality of Service 1/2 is preserved across the bridge).
*   **Hub Broker (Cloud):** Acts as the central aggregator.

## Lab Outcome
*   The **Publisher** connected to the Leaf broker and sent a message to `leaf_sensor/temperature`.
*   The **Bridge** forwarded this message.
*   The **Subscriber** connected to the Hub broker and successfully received the message on `hub_sensor/temperature`.
