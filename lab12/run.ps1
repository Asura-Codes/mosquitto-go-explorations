# Lab 12 Execution Script

# 1. Start Brokers (Hub + 2 Edges)
Write-Host "Starting Hub and Edge Brokers..."
docker compose up -d

Write-Host "Waiting for brokers to stabilize..."
Start-Sleep -Seconds 5

# 2. Build Apps
Write-Host "Building Applications..."
go build -o subscriber_hub/sub.exe ./subscriber_hub/main.go
go build -o publisher_edge/pub.exe ./publisher_edge/main.go

# 3. Start Hub Workers (Shared Subscription)
Write-Host "`n--- Starting 2 Hub Workers (Shared Subscription) ---"
$env:WORKER_ID = "A"
$workerA = Start-Job -ScriptBlock { cd $using:PWD; ./subscriber_hub/sub.exe }

$env:WORKER_ID = "B"
$workerB = Start-Job -ScriptBlock { cd $using:PWD; ./subscriber_hub/sub.exe }

Start-Sleep -Seconds 2

# 4. Publish to Edge 1
Write-Host "`n--- Publishing to Edge 1 (Port 1884) ---"
./publisher_edge/pub.exe -broker="localhost:1884" -prefix="edge1" -id="pub-edge1"

# 5. Publish to Edge 2
Write-Host "`n--- Publishing to Edge 2 (Port 1885) ---"
./publisher_edge/pub.exe -broker="localhost:1885" -prefix="edge2" -id="pub-edge2"

Write-Host "`nWaiting for processing..."
Start-Sleep -Seconds 3

# 6. Retrieve Output
Write-Host "`n--- Hub Worker A Output ---"
Receive-Job -Job $workerA
Stop-Job -Job $workerA
Remove-Job -Job $workerA

Write-Host "`n--- Hub Worker B Output ---"
Receive-Job -Job $workerB
Stop-Job -Job $workerB
Remove-Job -Job $workerB

# 7. Cleanup
Write-Host "`nCleaning up..."
docker compose down
rm subscriber_hub/sub.exe
rm publisher_edge/pub.exe
