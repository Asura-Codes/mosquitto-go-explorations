# Lab 1 Execution Script

Write-Host "Starting Mosquitto Broker..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

Write-Host "Checking Broker Logs..."
docker logs lab1-mosquitto

Write-Host "`nBroker is running. You can verify it by looking at the logs above for 'mosquitto version ... starting'."
