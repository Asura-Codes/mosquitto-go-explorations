# Lab 11 Summary: WebSockets and Browser Clients

In this lab, we extended MQTT communication to the web browser by configuring Mosquitto to support WebSockets.

## Key Takeaways

### 1. WebSocket Listener
- **Configuration:** In `mosquitto.conf`, we added a second `listener` on port `9001` with the `protocol websockets` directive.
- **Port Choice:** Port `1883` is the standard for TCP MQTT, while `9001` is the common convention for MQTT over WebSockets.

### 2. Browser Client (MQTT.js)
- **Library:** We used the popular `MQTT.js` library via CDN.
- **Connection String:** Browser clients connect using `ws://` or `wss://` (secure) protocols rather than `tcp://`.
- **Async Nature:** Just like the Go client, the browser client is event-driven, handling `connect` and `message` events via callbacks.

### 3. Cross-Protocol Communication
- **Verification:** We demonstrated that a backend application (Go) using standard TCP can communicate seamlessly with a frontend application (Browser) using WebSockets. 
- **Broker as a Bridge:** The Mosquitto broker handles the protocol translation between TCP and WebSockets internally, allowing topics to remain the same across all client types.

## Execution Example

When running `run.ps1`:

1. **Broker Starts:** Both ports `1883` and `9001` are opened.
2. **Browser Client:** Opens `index.html`, connects to `ws://localhost:9001`, and subscribes to `lab11/#`.
3. **Go Publisher:** Connects to `tcp://localhost:1883` and sends messages.
4. **Result:** The browser UI updates in real-time with the Go publisher's data:
   ```text
   [19:25:01] lab11/data: Sensor update #1 - Temperature: 20.10
   [19:25:04] lab11/data: Sensor update #2 - Temperature: 20.20
   ```

## Workflow Adherence
- Compiled the Go publisher to `.exe` as per project standards.
- Automated the environment setup via PowerShell.
