# Lab 1 Summary: Setting Up the Environment

## Key Takeaways
- **Mosquitto Configuration:** Configured basic listeners for MQTT (1883) and WebSockets (9001).
- **Docker Integration:** Running Mosquitto via Docker for quick setup.
- **Verification:** Using `docker logs` to check socket status.

## Execution Example

Running `run.ps1` starts the container and displays the logs:

```text
Starting Mosquitto Broker...
[+] Running 1/1
 ✔ Container lab1-mosquitto  Started
Checking Broker Logs...

167302...: mosquitto version 2.0.18 starting
167302...: Config loaded from /mosquitto/config/mosquitto.conf.
167302...: Opening ipv4 listen socket on port 1883.
167302...: Opening ipv4 listen socket on port 9001.
```

- **Persistence:** Enabling `persistence true` allows saving state across restarts.
- **Anonymous Access:** Using `allow_anonymous true` for simple initial testing.
