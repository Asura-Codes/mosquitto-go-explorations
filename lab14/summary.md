# Lab 14 Summary: High Availability (Fan-out Publisher)

In this lab, we implemented a robust High Availability (HA) pattern for MQTT without using broker clustering. We utilized a **Fan-out Publisher** and a **Failover Subscriber** with **Three Independent Brokers**.

## Architecture

*   **Brokers:** Three independent Mosquitto instances on ports `1883`, `1884`, and `1885`. They are *not* bridged and share no state.
*   **Publisher:** Connects to **ALL** available brokers simultaneously. Every message is published to every connected broker.
*   **Subscriber:** Connects to **ONE** broker from the list. If that broker fails, it automatically reconnects to the next available one.

## Why this solves the "Split Brain" issue
In the previous setup, if the Publisher was on Broker A and the Subscriber failed over to Broker B, they couldn't communicate.
With **Fan-out Publishing**:
1.  The Publisher sends "Message X" to Broker A, Broker B, and Broker C.
2.  The Subscriber connects to Broker A and receives "Message X".
3.  Broker A dies.
4.  Subscriber reconnects to Broker B.
5.  Publisher (still connected to B and C) sends "Message Y" to B and C.
6.  Subscriber receives "Message Y" from Broker B.

**Result:** Zero downtime (other than reconnection time) and guaranteed message availability on any active broker.

## Trade-offs
*   **Bandwidth:** The publisher uses N times the bandwidth (where N is the number of brokers).
*   **Duplicates:** If a subscriber somehow managed to connect to multiple brokers (not typical with standard clients), it would get duplicates.
*   **Complexity:** The publisher logic is more complex, managing multiple client instances.

## Testing
1.  Run `run.ps1`.
2.  Start Subscriber and Publisher.
3.  Stop the broker the subscriber is currently using.
4.  Observe seamless failover and continued message reception.
