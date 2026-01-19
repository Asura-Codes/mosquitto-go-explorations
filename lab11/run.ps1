# Lab 11 Execution Script

# 1. Start Broker
Write-Host "Starting Mosquitto Broker with WebSockets support..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

# 2. Build Go Publisher
Write-Host "Building Go Publisher..."
go build -o publisher/pub.exe ./publisher/main.go

# 3. Instructions
Write-Host "`n--- WEB CLIENT SETUP ---"
Write-Host "Please open the following file in your browser:"
Write-Host "$PWD\index.html"
Write-Host "`nOnce the browser is open and connected, press ENTER to start the Go Publisher."
Read-Host

# 4. Run Go Publisher
Write-Host "Starting Go Publisher (TCP:1883) -> Browser (WS:9001)"
./publisher/pub.exe

# 5. Cleanup
Write-Host "`nCleaning up..."
docker compose down
rm publisher/pub.exe