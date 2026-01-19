# Lab 16 Summary: High Availability - Server-Side Replication (Star Topology)

In this lab, we implemented a robust **High Availability** solution using **Server-Side Replication** and **Load Balancing**.

## Architecture: Star Topology (Hub & Spoke)
To avoid routing loops inherent in a Full Mesh, we implemented a **Star Topology**:
*   **Hub (mosquitto1):** The central node. It is "Passive," meaning it accepts incoming bridge connections but initiates none.
*   **Leaves (mosquitto2, mosquitto3):** These nodes initiate a bridge connection to the Hub.

**Data Flow:**
*   **Leaf -> Hub:** `mosquitto2` sends messages to `mosquitto1`.
*   **Hub -> Leaf:** `mosquitto1` replicates messages to `mosquitto3`.
*   **Result:** A message published to ANY broker is replicated to ALL brokers exactly once.

## Key Configuration Patterns

### 1. Loop Prevention (Star Topology)
By strictly defining the direction of the bridge (Leaf -> Hub), we prevent circular routing loops (A->B->C->A) that occur in a naive full mesh.

### 2. Reliable Replication (QoS 1)
We configured the bridges to use **QoS 1** (`topic # both 1`). This ensures that messages are acknowledged across the bridge, preventing data loss if the network glitches.

### 3. Session Stickiness (HAProxy)
We configured HAProxy with **Source IP Stickiness**:
```haproxy
stick-table type ip size 200k expire 30m
stick on src
```
**Why?** Mosquitto does not share session state (subscriptions/offline queue).
*   **Without Stickiness:** A client reconnecting might hit a different broker (Round Robin) that doesn't know about its session, causing message loss (the "Disconnected Session" problem).
*   **With Stickiness:** The client always reconnects to the *same* broker (e.g., Broker A), preserving its `CleanSession=false` queue and ensuring it receives all missed messages.

## Testing the Resilience
1.  **Start Subscriber:** Connects to Broker A (Pinned).
2.  **Stop Subscriber.**
3.  **Publish Messages:** Sent to Broker B (Load Balanced) -> Replicated to Broker A -> Queued for Subscriber.
4.  **Restart Subscriber:** Reconnects to Broker A -> **Receives all queued messages.**
