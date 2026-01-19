#!/bin/sh
# init_auth.sh - Generates Mosquitto password file safely

PASSWD_FILE="/work/passwd"

# Create file if it doesn't exist
touch "$PASSWD_FILE"

# Set secure permissions (suppresses warnings)
chmod 0700 "$PASSWD_FILE"

# Add users (using batch mode -b)
# Note: mosquitto_passwd -b overwrites or appends correctly
mosquitto_passwd -b "$PASSWD_FILE" admin admin123
mosquitto_passwd -b "$PASSWD_FILE" sensor sensor123

echo "Password file generated successfully."
