$lab_name = "lab17"
Set-Location $PSScriptRoot

Write-Host "Cleaning up previous runs..."
docker compose down -v
if (Test-Path "dynamic-security.json") { rm dynamic-security.json }
if (Test-Path "client/client.exe") { rm client/client.exe }

Write-Host "Generating Certificates..."
./gen_certs.ps1

Write-Host "Initializing Dynamic Security Store..."
# We use a temporary container to initialize the JSON file with an admin user
docker run --rm -v "${PSScriptRoot}:/target" eclipse-mosquitto:2.0 mosquitto_ctrl dynsec init /target/dynamic-security.json admin secret

Write-Host "Starting Mosquitto Broker..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

Write-Host "Configuring Dynamic Security via mosquitto_ctrl..."
# Define common arguments for secure connection
$ctrlArgs = @("--cafile", "/mosquitto/certs/ca.crt", "-h", "localhost", "-p", "8883")

# Create a role for sensors
docker exec mosquitto_lab17 mosquitto_ctrl $ctrlArgs -u admin -P secret dynsec createRole sensorRole
# Allow the role to publish and subscribe to 'test/topic'
docker exec mosquitto_lab17 mosquitto_ctrl $ctrlArgs -u admin -P secret dynsec addRoleACL sensorRole publishClientSend test/topic allow
docker exec mosquitto_lab17 mosquitto_ctrl $ctrlArgs -u admin -P secret dynsec addRoleACL sensorRole subscribeLiteral test/topic allow
# Create a client user
docker exec mosquitto_lab17 mosquitto_ctrl $ctrlArgs -u admin -P secret dynsec createClient sensor1 -p password123
# Assign the role to the client
docker exec mosquitto_lab17 mosquitto_ctrl $ctrlArgs -u admin -P secret dynsec addClientRole sensor1 sensorRole

Write-Host "Building Go Client..."
cd client
go build -o client.exe .
cd ..

Write-Host "Starting Subscriber (sensor1)..."
$subJob = Start-Job -ScriptBlock {
    Set-Location $using:PSScriptRoot
    ./client/client.exe -user sensor1 -pass password123 -action sub
}

Start-Sleep -Seconds 2

Write-Host "Running Publisher (sensor1)..."
$pubJob = Start-Job -ScriptBlock {
    Set-Location $using:PSScriptRoot
    ./client/client.exe -user sensor1 -pass password123 -action pub
}

Write-Host "Allowing messages to flow for 5 seconds..."
Start-Sleep -Seconds 5

Write-Host "`nSubscriber Output:"
Receive-Job -Job $subJob
Stop-Job -Job $subJob
Remove-Job -Job $subJob

Write-Host "Stopping Publisher..."
Stop-Job -Job $pubJob
Remove-Job -Job $pubJob

Write-Host "`nCleaning up..."
docker compose down -v
rm dynamic-security.json
rm client/client.exe
