# Lab 6 Execution Script

# 1. Generate Password File
Write-Host "Generating Password File (using Docker)..."
# We check if 'passwd' exists to avoid overwriting or docker overhead if repeated, 
# but for "repeatable" labs, we regenerate.
New-Item -ItemType File -Path .\passwd -Force | Out-Null

# Use a helper script mounted into the container to avoid shell/quoting issues
docker run --rm -v "${PWD}:/work" eclipse-mosquitto:latest sh /work/init_auth.sh

if (-not (Test-Path .\passwd)) {
    Write-Error "Failed to create password file."
    exit 1
}

# 2. Start Broker
Write-Host "Starting Mosquitto Broker for Lab 6..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

# 3. Build Apps
Write-Host "Building components..."
go build -o admin_client/admin.exe ./admin_client/main.go
go build -o sensor_client/sensor.exe ./sensor_client/main.go

# 4. Run Demo
Write-Host "`n--- STEP 1: Starting Admin (Has Full Access) ---"
$adminJob = Start-Job -ScriptBlock { cd $using:PWD; ./admin_client/admin.exe }
Start-Sleep -Seconds 2

Write-Host "`n--- STEP 2: Running Sensor (Restricted Access) ---"
./sensor_client/sensor.exe
Start-Sleep -Seconds 2

Write-Host "`n--- Admin Output (Should show only the ALLOWED message) ---"
Receive-Job -Job $adminJob
Stop-Job -Job $adminJob
Remove-Job -Job $adminJob

Write-Host "`nCleaning up..."
docker compose down
rm admin_client/admin.exe
rm sensor_client/sensor.exe
rm passwd
