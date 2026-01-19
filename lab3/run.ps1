# Lab 3 Execution Script

Write-Host "Starting Mosquitto Broker for Lab 3..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

Write-Host "Building Subscriber..."
go build -o subscriber/subscriber.exe ./subscriber/main.go

Write-Host "Building Publisher..."
go build -o publisher/publisher.exe ./publisher/main.go

Write-Host "`n--- STEP 1: Sending Retained Message (Subscriber is OFFLINE) ---"
./publisher/publisher.exe -mode retained
Start-Sleep -Seconds 1

Write-Host "`n--- STEP 2: Starting Subscriber (Should receive Retained msg immediately) ---"
$subJob = Start-Job -ScriptBlock {
    cd $using:PWD
    ./subscriber/subscriber.exe
}
# Give subscriber time to connect and process the retained message
Start-Sleep -Seconds 2

Write-Host "`n--- STEP 3: Sending Live QoS Messages ---"
./publisher/publisher.exe -mode qos

# Allow time for messages to be processed
Start-Sleep -Seconds 2

Write-Host "`n--- Subscriber Output ---"
Receive-Job -Job $subJob
Stop-Job -Job $subJob
Remove-Job -Job $subJob

Write-Host "`nCleaning up..."
docker compose down
rm subscriber/subscriber.exe
rm publisher/publisher.exe
