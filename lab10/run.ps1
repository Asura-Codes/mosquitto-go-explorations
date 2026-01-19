$ErrorActionPreference = "Stop"

Write-Host "Starting Mosquitto Broker for Lab 10..."
docker compose -f docker-compose.yml up -d

Write-Host "Waiting for broker to stabilize..."
Start-Sleep -Seconds 3

# Ensure go.sum is updated
if (-not (Test-Path "go.sum")) {
    Write-Host "Tidying Go modules..."
    go mod tidy
}

Write-Host "Building Applications..."
go build -o publisher.exe ./publisher/main.go
go build -o monitor.exe ./monitor/main.go
go build -o deleter.exe ./deleter/main.go

Write-Host "`n--- STEP 1: Setting Initial State (Publisher) ---"
./publisher.exe

Write-Host "`n--- STEP 2: Starting Monitor (Catches up with Retained state) ---"
$monJob = Start-Job -ScriptBlock {
    Set-Location $using:PWD
    ./monitor.exe
}
Start-Sleep -Seconds 3

Write-Host "`n--- Monitor Output (Initial Catch-up):"
Receive-Job -Job $monJob -Keep

Write-Host "`n--- STEP 3: Updating State (Publisher) ---"
# We can run publisher again to update states
$env:MQTT_BROKER = "localhost:1883"
./publisher.exe
Start-Sleep -Seconds 2

Write-Host "`n--- Monitor Output (Live Update):"
Receive-Job -Job $monJob -Keep | Select-Object -Last 3

Write-Host "`n--- STEP 4: Deleting a Retained Message ---"
./deleter.exe "settings/mode"
Start-Sleep -Seconds 2

Write-Host "`n--- Monitor Output (After Deletion):"
# The monitor will receive an empty payload for the deleted topic
Receive-Job -Job $monJob -Keep | Select-Object -Last 1

Write-Host "`nLab Complete."

# Cleanup
Stop-Job $monJob
Remove-Job $monJob
Remove-Item *.exe -ErrorAction SilentlyContinue

Write-Host "Stopping Docker containers..."
docker compose -f docker-compose.yml down
