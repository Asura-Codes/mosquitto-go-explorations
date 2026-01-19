# Lab 15 Summary: High Availability - Load Balancing

In this lab, we explored using **HAProxy** as a TCP load balancer in front of multiple Mosquitto brokers.

## Architecture

*   **Load Balancer:** HAProxy listening on port `1883`.
*   **Brokers:** Three independent Mosquitto instances (backend servers).
*   **Clients:** Connect to the HAProxy port, unaware of the specific backend broker.

## Key Observations

### Round Robin Distribution
By default, HAProxy uses `roundrobin`. 
1.  **Connection 1 (Subscriber):** Routed to Broker A.
2.  **Connection 2 (Publisher):** Routed to Broker B.

**Result:** The Subscriber **DOES NOT** receive messages because Broker A and Broker B are not bridged and do not share state. This demonstrates the limitations of simple load balancing with stateful protocols like MQTT without clustering or bridging.

### Session Stickiness (Optional)
By enabling sticky sessions (e.g., `stick on src`), we ensure that a specific client IP always connects to the same backend broker (as long as it is up). This helps maintain session state (subscriptions) but does not solve the data sharing problem if different clients (Pub vs Sub) connect from different IPs.

## HAProxy Stats
We accessed the stats dashboard at `http://localhost:8404/monitor` to verify connection distribution and server health.
