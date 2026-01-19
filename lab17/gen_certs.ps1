$certDir = "certs"

# Ensure clean slate
if (Test-Path "$certDir") { Remove-Item "$certDir" -Recurse -Force }
New-Item -ItemType Directory -Force -Path $certDir | Out-Null

Write-Host "Generating Certificates in $certDir..."

# Check for OpenSSL
if (-not (Get-Command openssl -ErrorAction SilentlyContinue)) {
    Write-Host "OpenSSL not found on host. Using Docker..."
    # We will use a docker container to generate certs if openssl is missing
    # This maps the current directory's cert folder to /certs in container
    docker run --rm -v "${PWD}/${certDir}:/certs" --entrypoint /bin/sh alpine/openssl -c "
        apk add --no-cache openssl
        cd /certs
        
        # 1. CA
        openssl req -new -x509 -days 365 -extensions v3_ca -nodes -keyout ca.key -out ca.crt -subj '/CN=MosquittoCA'
        
        # 2. Server
        openssl req -new -nodes -keyout server.key -out server.csr -subj '/CN=localhost'
        echo 'subjectAltName=DNS:localhost,IP:127.0.0.1' > server.ext
        openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -extfile server.ext
        rm server.ext

        # 3. Client (optional, but good to have)
        openssl req -new -nodes -keyout client.key -out client.csr -subj '/CN=client1'
        openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365
        
        chmod 644 *.crt *.key
    "
} else {
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
}

Write-Host "Certificates generated successfully."
