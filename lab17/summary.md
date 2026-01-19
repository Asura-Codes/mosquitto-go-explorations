# Lab 17 Summary: Dynamic Security Management

In this lab, we explored the **Dynamic Security Plugin** (`mosquitto_dynamic_security`), which allows managing authentication and authorization at runtime without restarting the broker.

## Key Concepts

1.  **Dynamic Security Plugin:** Replaces static password and ACL files. Stores configuration in a JSON file (e.g., `dynamic-security.json`).
2.  **`mosquitto_ctrl`:** A CLI tool (MQTT client) used to send administrative commands to the broker via the `$CONTROL` topic API.
3.  **RBAC (Role-Based Access Control):**
    *   **Clients:** Users/Devices.
    *   **Roles:** Sets of ACL permissions (publish/subscribe to topics).
    *   **Groups:** Collections of clients that share roles.

## Lab Steps

1.  **Initialization:**
    *   Loaded the plugin in `mosquitto.conf`.
    *   Initialized the security store using `docker run ... mosquitto_ctrl dynsec init`. This created the initial `admin` user in `dynamic-security.json`.

2.  **User Management:**
    *   Created a role `sensorRole` with `allow` permissions for `publishClientSend` and `subscribeLiteral` on `test/topic`.
    *   Created a client `sensor1` with a password.
    *   Assigned `sensorRole` to `sensor1`.
    *   All these steps were done via `docker exec ... mosquitto_ctrl` while the broker was running.

3.  **Go Integration:**
    *   Verified that the Go client can connect and communicate using the `sensor1` credentials.
    *   Confirmed that permissions are enforced by the Dynamic Security plugin.

