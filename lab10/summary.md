# Lab 10 Summary: State Management with Retained Messages

In this lab, we explored how MQTT can be used for state management, similar to a Key-Value (KV) store, using the **Retained Messages** feature.

## Key Takeaways

1.  **Retained Flag:**
    *   When a message is published with the `Retain` flag set to `true`, the broker stores the message for that topic.
    *   Any *new* subscriber to that topic will immediately receive the last retained message.
    *   This is perfect for configuration settings, status updates, or "last known good" values.

2.  **State Synchronization:**
    *   The `monitor` application demonstrated that even though it started *after* the initial states were published, it received the current state immediately upon subscription.
    *   Subsequent updates are received as normal (non-retained) messages.

3.  **Deleting State (Nulling):**
    *   To delete a retained message from the broker, you must publish a message with an **empty payload** and the `Retain` flag set to `true`.
    *   The `deleter` application demonstrated this by clearing the `settings/mode` topic.
    *   Active subscribers receive a notification with the empty payload, signaling that the state has been removed.

## Execution Example

When running `run.ps1`, you will see a sequence that demonstrates "Time Travel" (Retained Messages):

1. **Setting Initial State:** The publisher sends data and disconnects.
2. **Monitor Connects:**
   ```text
   Recv: [RETAINED] settings/mode = automatic
   Recv: [RETAINED] settings/threshold = 25.5
   ```
   *Note: The monitor receives these even though it started LATER.*
3. **Live Update:**
   ```text
   Recv: settings/mode = automatic (Live update)
   ```
4. **Deletion (Nulling):**
   ```text
   Recv: settings/mode = 
   ```
   *Note: Receiving an empty payload on a topic indicates the state was deleted.*

## Workflow Adherence
    *   This lab strictly followed the **Build then Run** workflow, compiling Go source into `.exe` binaries before execution.
    *   VS Code tasks were added for manual Docker management.
