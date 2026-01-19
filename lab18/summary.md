# Lab 18 Summary: External Authentication Integration

In this lab, we explored how to offload MQTT authentication and authorization to an external system using a plugin.

## Key Concepts

1.  **Authentication Plugins:** Mosquitto supports external plugins (written in C or Go) that can intercept connection and ACL requests.
2.  **`mosquitto-go-auth`:** A powerful plugin that allows using various backends like HTTP (Webhooks), MySQL, PostgreSQL, Redis, etc.
3.  **Delegated Auth (Webhooks):** The broker acts as a proxy, forwarding credentials to an external HTTP service (`auth-service`) and enforcing the result.
4.  **Separation of Concerns:** The broker handles the MQTT protocol, while the `auth-service` handles user management and business logic.

## Lab Architecture

*   **Mosquitto Broker:** Configured with `go-auth` plugin and `http` backend.
*   **Auth Service (Go):** A simple HTTP server implementing `/user` and `/acl` endpoints.
    *   Validates users against a mock database.
    *   Enforces topic-based ACLs (e.g., Alice can only read/write to `sensors/alice`).

## Results

*   **Alice** connects successfully with `password123`.
*   **Bob** connects successfully with `secret456`.
*   **Mallory** (unknown user) is rejected immediately by the `auth-service`.
*   The `auth-service` logs show every verification attempt made by the broker in real-time.
