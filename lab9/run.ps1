$ErrorActionPreference = "Stop"

Write-Host "Starting Mosquitto Brokers (Hub and Leaf) for Lab 9..."
docker compose -f docker-compose.yml up -d --build

Write-Host "Waiting for brokers to stabilize..."
Start-Sleep -Seconds 5

# Ensure go.sum exists
if (-not (Test-Path "go.sum")) {
    Write-Host "Tidying Go modules..."
    go mod tidy
}

Write-Host "Building Subscriber Hub..."
go build -o subscriber_hub.exe ./subscriber_hub/main.go

Write-Host "Building Publisher Leaf..."
go build -o publisher_leaf.exe ./publisher_leaf/main.go

Write-Host "Starting Subscriber on Hub (Background)..."
$subJob = Start-Job -ScriptBlock {
    Set-Location $using:PWD
    ./subscriber_hub.exe
}

Write-Host "Waiting for subscriber to connect..."
Start-Sleep -Seconds 3

Write-Host "--- Publishing to Leaf Broker ---"
./publisher_leaf.exe

Write-Host "--- Reading Subscriber Output ---"
# Loop briefly to show output
for ($i=0; $i -lt 5; $i++) {
    $out = Receive-Job -Job $subJob -Keep
    if ($out) {
        Write-Host $out
        # Clear the output buffer so we don't reprint
        Receive-Job -Job $subJob | Out-Null
    }
    Start-Sleep -Seconds 1
}

Write-Host "`nTest Complete."
Write-Host "Stop the subscriber job? (y/n)"
# In a real automated script we might just kill it, but for a lab interactive feel:
Stop-Job $subJob
Remove-Job $subJob

Write-Host "Cleaning up executables..."
Remove-Item -Path .\subscriber_hub.exe -ErrorAction SilentlyContinue
Remove-Item -Path .\publisher_leaf.exe -ErrorAction SilentlyContinue

Write-Host "Stopping Docker containers..."
docker compose -f docker-compose.yml down
