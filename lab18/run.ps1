$lab_name = "lab18"
Set-Location $PSScriptRoot

Write-Host "Cleaning up previous runs..."
docker compose down -v
if (Test-Path "client/client.exe") { rm client/client.exe }

Write-Host "Starting Lab 18 environment..."
docker compose up -d --build

Write-Host "Waiting for services to start..."
Start-Sleep -Seconds 5

Write-Host "Building Go Client..."
cd client
go build -o client.exe .
cd ..

Write-Host "Testing Authentication with 'alice' (Correct credentials)..."
# Alice has RW access to sensors/alice
$aliceJob = Start-Job -ScriptBlock {
    Set-Location $using:PSScriptRoot
    ./client/client.exe -user alice -pass password123 -action pub -topic sensors/alice
}
Start-Sleep -Seconds 2

Write-Host "Testing Authentication with 'bob' (Correct credentials)..."
# Bob has RW access to sensors/bob
$bobJob = Start-Job -ScriptBlock {
    Set-Location $using:PSScriptRoot
    ./client/client.exe -user bob -pass secret456 -action sub -topic sensors/bob
}
Start-Sleep -Seconds 2

Write-Host "Testing ACL: Alice attempts to publish to 'sensors/bob' (Forbidden)..."
$aliceAclJob = Start-Job -ScriptBlock {
    Set-Location $using:PSScriptRoot
    # Alice tries to publish to Bob's topic
    ./client/client.exe -user alice -pass password123 -action pub -topic sensors/bob
}
Start-Sleep -Seconds 5

Write-Host "Testing Authentication with 'mallory' (Unknown user)..."
./client/client.exe -user mallory -pass evil -action sub -topic sensors/mallory

Write-Host "`nAlice Output (Publisher - sensors/alice):"
Receive-Job -Job $aliceJob
Stop-Job -Job $aliceJob
Remove-Job -Job $aliceJob

Write-Host "`nAlice ACL Test Output (sensors/bob):"
Receive-Job -Job $aliceAclJob
Stop-Job -Job $aliceAclJob
Remove-Job -Job $aliceAclJob

Write-Host "`nBob Output (Subscriber):"
Receive-Job -Job $bobJob
Stop-Job -Job $bobJob
Remove-Job -Job $bobJob

Write-Host "`nCheck docker logs for auth-service to see the verification requests:"
docker logs auth_service_lab18

Write-Host "`nCleaning up..."
docker compose down -v
rm client/client.exe

