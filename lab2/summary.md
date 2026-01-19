# Lab 2 Summary: Go Client - Basic MQTT Messaging

## Key Takeaways
- **Paho MQTT Client:** Used the `paho.mqtt.golang` library.
- **Connection Options:** Configured `ClientOptions` with a unique Client ID.
- **Publishing:** Sent messages to topics like `sensor/temp`.
- **Subscribing:** Registered a callback function to handle incoming messages.
- **Topic Wildcards:** Used `sensor/#` to match any sub-topic.

## Execution Example

When running `run.ps1`, the publisher sends data that the subscriber catches via a wildcard:

1. **Publisher:**
   ```text
   Connected to Broker.
   Publishing to sensor/temp: 23.5
   Publishing to sensor/humidity: 60
   ```

2. **Subscriber:**
   ```text
   Connected to Broker. Subscribing to sensor/#
   Received message on topic: sensor/temp | Payload: 23.5
   Received message on topic: sensor/humidity | Payload: 60
   ```

## Project Structure
- Organized into separate `publisher` and `subscriber` directories.
- Built as standalone Windows executables (`.exe`).
- Automation via `run.ps1` for building and executing.
