$certDir = "certs"

# Ensure clean slate
if (Test-Path "$certDir\*.key") { Remove-Item "$certDir\*.key" }
if (Test-Path "$certDir\*.crt") { Remove-Item "$certDir\*.crt" }
if (Test-Path "$certDir\*.csr") { Remove-Item "$certDir\*.csr" }
if (Test-Path "$certDir\*.srl") { Remove-Item "$certDir\*.srl" }

Write-Host "Generating Certificates in $certDir..."

# 1. Generate CA Key and Certificate
Write-Host "1. Generating CA..."
openssl req -new -x509 -days 365 -extensions v3_ca -nodes -keyout "$certDir\ca.key" -out "$certDir\ca.crt" -subj "/CN=MosquittoCA"

# 2. Generate Server Key and CSR
Write-Host "2. Generating Server Key/CSR..."
openssl req -new -nodes -keyout "$certDir\server.key" -out "$certDir\server.csr" -subj "/CN=localhost"

# Create extension file for SANs
$extContent = "subjectAltName=DNS:localhost,IP:127.0.0.1"
Set-Content -Path "$certDir\server.ext" -Value $extContent

# 3. Sign Server Certificate with CA
Write-Host "3. Signing Server Certificate..."
openssl x509 -req -in "$certDir\server.csr" -CA "$certDir\ca.crt" -CAkey "$certDir\ca.key" -CAcreateserial -out "$certDir\server.crt" -days 365 -extfile "$certDir\server.ext"

# Remove extension file
Remove-Item "$certDir\server.ext"

# 4. Generate Client Key and CSR
Write-Host "4. Generating Client Key/CSR..."
openssl req -new -nodes -keyout "$certDir\client.key" -out "$certDir\client.csr" -subj "/CN=client1"

# 5. Sign Client Certificate with CA
Write-Host "5. Signing Client Certificate..."
openssl x509 -req -in "$certDir\client.csr" -CA "$certDir\ca.crt" -CAkey "$certDir\ca.key" -CAcreateserial -out "$certDir\client.crt" -days 365

Write-Host "Certificates generated successfully."

# Set permissions (important for Mosquitto in Docker sometimes, though less critical on Windows mounts unless mapped to Linux user)
# We will handle permissions via Docker if needed.
