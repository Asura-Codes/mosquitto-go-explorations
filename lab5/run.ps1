# Lab 5 Execution Script

Write-Host "Starting Mosquitto Broker for Lab 5..."
docker compose up -d

Write-Host "Waiting for broker to start..."
Start-Sleep -Seconds 3

Write-Host "Building components..."
go build -o monitor/monitor.exe ./monitor/main.go
go build -o lwt_client/lwt_client.exe ./lwt_client/main.go
go build -o responder/responder.exe ./responder/main.go
go build -o requester/requester.exe ./requester/main.go

Write-Host "`n--- STEP 1: Starting Monitor ---"
$monitorJob = Start-Job -ScriptBlock { cd $using:PWD; ./monitor/monitor.exe }
Start-Sleep -Seconds 2

Write-Host "`n--- STEP 2: Last Will and Testament (LWT) Demo ---"
Write-Host "LWT Client will connect, then 'crash'..."
./lwt_client/lwt_client.exe
Start-Sleep -Seconds 2
Write-Host "LWT Client has crashed. Checking Monitor for LWT message..."

Write-Host "`n--- STEP 3: Request-Reply Demo ---"
$responderJob = Start-Job -ScriptBlock { cd $using:PWD; ./responder/responder.exe }
Start-Sleep -Seconds 2
./requester/requester.exe
Start-Sleep -Seconds 1

Write-Host "`n--- Final Monitor Output ---"
Receive-Job -Job $monitorJob
Stop-Job -Job $monitorJob
Remove-Job -Job $monitorJob

Write-Host "`nCleaning up..."
Stop-Job -Job $responderJob
Remove-Job -Job $responderJob
docker compose down
Get-ChildItem -Recurse -Include *.exe | Remove-Item
