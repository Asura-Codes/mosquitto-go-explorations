# Lab 3 Summary: Reliability and Quality of Service

## Key Takeaways

### 1. Retained Messages
- **Concept:** A published message with the `retained` flag set to `true` is stored by the broker.
- **Behavior:** New subscribers to that topic immediately receive the last retained message.

### 2. Quality of Service (QoS) Levels
We implemented and observed three QoS levels:
- **QoS 0 (At Most Once):** "Fire and forget".
- **QoS 1 (At least Once):** Guarantees delivery via PUBACK.
- **QoS 2 (Exactly Once):** Guarantees exactly one delivery (4-way handshake).

## Execution Example

Running `run.ps1` demonstrates the behavior:

1. **Publisher (Retained Mode):**
   ```text
   Publishing retained message to 'lab3/config'
   ```
2. **Subscriber Starts (Later):**
   ```text
   Subscriber: Connected.
   Recv: [RETAINED] Topic: lab3/config | Payload: {"units": "celsius"}
   ```
   *Note: Even though it started later, it received the last known configuration.*

3. **Publisher (QoS Mode):**
   ```text
   Publishing [QoS 0] temp: 20
   Publishing [QoS 1] temp: 21
   Publishing [QoS 2] temp: 22
   ```

4. **Subscriber (Live Output):**
   ```text
   Recv: [QoS 0] Topic: lab3/sensor | Payload: 20
   Recv: [QoS 1] Topic: lab3/sensor | Payload: 21
   Recv: [QoS 2] Topic: lab3/sensor | Payload: 22
   ```

### 3. Clean Code Implementation
- **Publisher:** Separated logic into `publishRetainedMessage` and `publishQoSSequence` functions (SRP).
- **Subscriber:** Used a dedicated `messageHandler` to separate processing from connection setup.
