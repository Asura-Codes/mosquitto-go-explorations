# Lab 13: Security - TLS/SSL Encryption Summary

## Key Concepts
*   **Transport Layer Security (TLS):** Encrypting the connection between client and broker to prevent eavesdropping and tampering.
*   **Certificate Authority (CA):** A trusted entity that issues digital certificates. In this lab, we acted as our own CA.
*   **Mutual TLS (mTLS):** A security mode where *both* the server (broker) and the client authenticate each other using certificates.
*   **PEM Format:** The standard format for storing certificates and keys.

## Implementation Details
1.  **Certificate Generation:** We generated a self-signed CA, and then used it to sign a Server Certificate (for Mosquitto) and a Client Certificate (for our Go apps).
2.  **Mosquitto Configuration:**
    *   Enabled a listener on port `8883` (standard MQTT over TLS port).
    *   configured `cafile`, `certfile`, and `keyfile` to point to the generated assets.
    *   Set `require_certificate true` to enforce mTLS.
    *   Set `use_identity_as_username true` (optional but useful) to use the CN from the client cert as the MQTT username.
3.  **Go Client (`crypto/tls`):**
    *   Loaded the CA cert into a `x509.CertPool` to verify the broker.
    *   Loaded the client's public/private key pair using `tls.LoadX509KeyPair`.
    *   Configured `tls.Config` with these assets.
    *   Used the `tls://` scheme in the broker URI.

## Testing
*   We verified that the publisher could connect securely and send messages.
*   We verified that the subscriber could connect securely and receive those messages.
*   Standard (non-TLS) connections or connections without the correct client cert would be rejected by the broker.
