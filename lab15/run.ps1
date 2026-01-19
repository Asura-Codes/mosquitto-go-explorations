$ErrorActionPreference = "Stop"

Write-Host "Starting Lab 15: High Availability - Load Balancing..." -ForegroundColor Cyan

# 1. Start Docker environment
Write-Host "Starting HAProxy and Brokers..."
$ErrorActionPreference = "Continue"
docker-compose down --remove-orphans 2>$null
$ErrorActionPreference = "Stop"
docker-compose up -d
if ($LASTEXITCODE -ne 0) {
    Write-Error "Failed to start Docker containers."
}

Write-Host "Waiting for services to initialize..."
Start-Sleep -Seconds 5

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
Write-Host "3. View HAProxy Stats: http://localhost:8404/monitor"
