# Lab 12 Summary: Advanced Topologies - Aggregation

In this final lab, we combined multiple advanced concepts to build a scalable **Edge-to-Cloud Aggregation** architecture.

## Architecture

1.  **Edge Brokers (Edge 1 & Edge 2):**
    -   Simulate local gateways at different physical locations.
    -   Accept local messages on ports `1884` and `1885`.
    -   Bridge messages to the central Hub, remapping them to the `aggregated/` namespace.
    -   *Example:* `edge1/data` -> `aggregated/edge1/data`.

2.  **Hub Broker (Cloud):**
    -   Acts as the central integration point (Port `1883`).
    -   Receives bridged traffic from all edges.

3.  **Hub Workers (Shared Subscription):**
    -   Connected to the Hub.
    -   Used a **Shared Subscription** (`$share/processors/aggregated/#`) to distribute the workload.
    -   This allows the system to scale horizontally; adding more workers automatically spreads the processing load without changing the publisher or broker configuration.

## Key Takeaways

-   **Bridging for Aggregation:** We can funnel data from thousands of edge locations into a central cloud broker using simple bridge configuration rules.
-   **Namespace Separation:** Remapping topics during bridging (e.g., adding `edge1/` prefix) is crucial to prevent collisions and identify the source of data in the central hub.
-   **Load Balancing:** MQTT v5 Shared Subscriptions provide a native way to parallelize message processing, ensuring high throughput for the aggregated data stream.

## Execution Example

When running `run.ps1`:

1.  **Publishers** send data to their local Edge brokers.
2.  **Edge Brokers** forward this data to the Hub.
3.  **Hub Workers** (A and B) take turns processing the incoming stream:

```text
--- Publishing to Edge 1 ---
Sent: Data from pub-edge1: 1
Sent: Data from pub-edge1: 2
...

--- Hub Worker A Output ---
[Worker A] Subscribed to shared topic '$share/processors/aggregated/#'
[Worker A] Processing: aggregated/edge1/data | Payload: Data from pub-edge1: 1
[Worker A] Processing: aggregated/edge2/data | Payload: Data from pub-edge2: 2

--- Hub Worker B Output ---
[Worker B] Subscribed to shared topic '$share/processors/aggregated/#'
[Worker B] Processing: aggregated/edge1/data | Payload: Data from pub-edge1: 2
[Worker B] Processing: aggregated/edge2/data | Payload: Data from pub-edge2: 1
```

*Note: The messages are distributed between Worker A and Worker B, demonstrating effective load balancing of the aggregated traffic.*
