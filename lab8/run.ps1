# Lab 8 Execution Script

# 1. Start Broker
Write-Host "Starting Mosquitto Broker for Lab 8..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

# 2. Build App
Write-Host "Building monitor..."
go build -o monitor/monitor.exe ./monitor/main.go

# 3. Run Demo
Write-Host "`n--- Starting Monitor ---"
$monitorJob = Start-Job -ScriptBlock { cd $using:PWD; ./monitor/monitor.exe }
Start-Sleep -Seconds 5

Write-Host "`n--- Generating Activity ---"
# We'll use a simple docker run to publish a message and trigger stats
docker run --rm eclipse-mosquitto:latest mosquitto_pub -h host.docker.internal -t "test/activity" -m "hello"
Start-Sleep -Seconds 2

Write-Host "`n--- Monitor Output ---"
Receive-Job -Job $monitorJob
Stop-Job -Job $monitorJob
Remove-Job -Job $monitorJob

Write-Host "`n--- Broker Logs (Showing management info) ---"
docker logs lab8-mosquitto

Write-Host "`nCleaning up..."
docker compose down
rm monitor/monitor.exe
