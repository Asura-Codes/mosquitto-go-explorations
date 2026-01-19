# Lab 2 Execution Script

Write-Host "Starting Mosquitto Broker..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

Write-Host "Building Subscriber..."
go build -o subscriber/subscriber.exe ./subscriber/main.go

Write-Host "Building Publisher..."
go build -o publisher/publisher.exe ./publisher/main.go

Write-Host "Starting Subscriber (Topic: sensor/#)..."
$subJob = Start-Job -ScriptBlock {
    cd $using:PWD
    ./subscriber/subscriber.exe
}

Start-Sleep -Seconds 2

Write-Host "Starting Publisher..."
./publisher/publisher.exe

Start-Sleep -Seconds 2

Write-Host "`nSubscriber Output:"
Receive-Job -Job $subJob
Stop-Job -Job $subJob
Remove-Job -Job $subJob

Write-Host "`nCleaning up..."
docker compose down
rm subscriber/subscriber.exe
rm publisher/publisher.exe
