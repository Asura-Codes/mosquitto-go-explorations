# Lab 13 Execution Script

Write-Host "Starting Mosquitto Broker (TLS Enabled)..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

Write-Host "Building Secure Subscriber..."
go build -o secure_sub/subscriber.exe ./secure_sub/main.go

Write-Host "Building Secure Publisher..."
go build -o secure_pub/publisher.exe ./secure_pub/main.go

Write-Host "Starting Secure Subscriber..."
$subJob = Start-Job -ScriptBlock {
    cd $using:PWD
    ./secure_sub/subscriber.exe
}

Start-Sleep -Seconds 2

Write-Host "Starting Secure Publisher..."
./secure_pub/publisher.exe

Start-Sleep -Seconds 2

Write-Host "`nSubscriber Output:"
Receive-Job -Job $subJob
Stop-Job -Job $subJob
Remove-Job -Job $subJob

Write-Host "`nCleaning up..."
docker compose down
rm secure_sub/subscriber.exe
rm secure_pub/publisher.exe
