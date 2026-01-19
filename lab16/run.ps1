$ErrorActionPreference = "Stop"

Write-Host "Starting Lab 16: High Availability - Server-Side Replication (Mesh Bridging)..." -ForegroundColor Cyan

# 1. Start Docker environment
Write-Host "Starting HAProxy and Mesh Brokers..."
$ErrorActionPreference = "Continue"
docker-compose down --remove-orphans 2>$null
$ErrorActionPreference = "Stop"
docker-compose up -d
if ($LASTEXITCODE -ne 0) {
    Write-Error "Failed to start Docker containers."
}

Write-Host "Waiting for mesh synchronization..."
Start-Sleep -Seconds 10

# 2. Build Go Applications
Write-Host "Building Publisher..."
cd publisher
go build -o publisher.exe .
if ($LASTEXITCODE -ne 0) {
    Write-Error "Failed to build Publisher."
}
cd ..

Write-Host "Building Subscriber..."
cd subscriber
go build -o subscriber.exe .
if ($LASTEXITCODE -ne 0) {
    Write-Error "Failed to build Subscriber."
}
cd ..

Write-Host "`nEnvironment is ready!" -ForegroundColor Green
Write-Host "Run the applications in separate terminals:"
Write-Host "1. .\subscriber\subscriber.exe"
Write-Host "2. .\publisher\publisher.exe"
Write-Host "Now, unlike Lab 15, the subscriber SHOULD receive messages even if routed to a different broker!"
