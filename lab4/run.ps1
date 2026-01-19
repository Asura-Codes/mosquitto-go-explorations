# Lab 4 Execution Script

Write-Host "Starting Mosquitto Broker for Lab 4..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

Write-Host "Building Subscriber..."
go build -o subscriber/subscriber.exe ./subscriber/main.go

Write-Host "Building Publisher..."
go build -o publisher/publisher.exe ./publisher/main.go

Write-Host "`n--- STEP 1: Establish Persistent Session ---"
Write-Host "Starting Subscriber (CleanSession=false) briefly..."
$subJob = Start-Job -ScriptBlock {
    cd $using:PWD
    ./subscriber/subscriber.exe -clean=false
}
Start-Sleep -Seconds 3
Stop-Job -Job $subJob
Remove-Job -Job $subJob
Write-Host "Subscriber disconnected. Session should be persisted on broker."

Write-Host "`n--- STEP 2: Publish messages while Subscriber is OFFLINE ---"
./publisher/publisher.exe

Write-Host "`n--- STEP 3: Reconnect Subscriber and Receive Queued Messages ---"
Write-Host "Starting Subscriber (CleanSession=false) again..."
$subJob = Start-Job -ScriptBlock {
    cd $using:PWD
    ./subscriber/subscriber.exe -clean=false
}
Start-Sleep -Seconds 3

Write-Host "`nSubscriber Output (Should show missed 'Alert' messages):"
Receive-Job -Job $subJob
Stop-Job -Job $subJob
Remove-Job -Job $subJob

Write-Host "`nCleaning up..."
docker compose down
rm subscriber/subscriber.exe
rm publisher/publisher.exe
