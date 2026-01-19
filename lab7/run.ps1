# Lab 7 Execution Script (MQTT v5)

# 1. Start Broker
Write-Host "Starting Mosquitto Broker for Lab 7..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

# 2. Build Apps
Write-Host "Building components..."
go build -o publisher/pub.exe ./publisher/main.go
go build -o subscriber/sub.exe ./subscriber/main.go

# 3. Run Demo
Write-Host "`n--- Starting Two Shared Workers (Load Balancing) ---"
# Worker A
$env:WORKER_ID = "A"
$workerA = Start-Job -ScriptBlock { cd $using:PWD; ./subscriber/sub.exe }
# Worker B
$env:WORKER_ID = "B"
$workerB = Start-Job -ScriptBlock { cd $using:PWD; ./subscriber/sub.exe }

Start-Sleep -Seconds 2

Write-Host "`n--- Starting Publisher (Sending 10 messages) ---"
./publisher/pub.exe

Write-Host "`nWaiting for workers to process..."
Start-Sleep -Seconds 2

Write-Host "`n--- Worker A Output ---"
Receive-Job -Job $workerA
Stop-Job -Job $workerA
Remove-Job -Job $workerA

Write-Host "`n--- Worker B Output ---"
Receive-Job -Job $workerB
Stop-Job -Job $workerB
Remove-Job -Job $workerB

Write-Host "`nCleaning up..."
docker compose down
rm publisher/pub.exe
rm subscriber/sub.exe
