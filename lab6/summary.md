# Lab 6 Summary: Authentication and Authorization

## Key Takeaways

### 1. Authentication (Who are you?)
- **Mechanism:** Mosquitto uses a password file (generated via `mosquitto_passwd`) to store hashed credentials (username + password).
- **Configuration:**
    - `allow_anonymous false`: Forces all clients to provide credentials.
    - `password_file /path/to/file`: Specifies the database of users.
- **Client Side:** Paho MQTT client uses `SetUsername()` and `SetPassword()` to send these credentials during the handshake.

### 2. Authorization (What can you do?)
- **Mechanism:** Access Control Lists (ACLs) define permissions for users.
- **Rules:**
    - `topic read`: User can only subscribe.
    - `topic write`: User can only publish.
    - `topic readwrite`: User can do both.
    - `user <username>`: Starts a block of rules for a specific user.

## Execution Example

When running `run.ps1`, the ACLs are tested by having a `sensor` client try to publish to two topics:

1. **Admin (Subscribed to `#`):**
   ```text
   Admin: Connected as 'admin'
   Admin: Subscribing to #
   Admin: Received [sensors/temp] - 22.5
   ```

2. **Sensor (Trying to publish):**
   ```text
   Sensor: Connected as 'sensor_1'
   Sensor: Publishing to 'sensors/temp' (ALLOWED)
   Sensor: Publishing to 'admin/secret' (DENIED BY ACL)
   ```

*Note: The Admin NEVER receives the message on `admin/secret` because the broker enforces the ACL and drops the unauthorized publication.*

## Security Best Practice
Always disable `allow_anonymous` in production. Use ACLs to enforce the **Principle of Least Privilege**—users should only have access to the topics they absolutely need.
