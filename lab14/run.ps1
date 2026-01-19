# Lab 14: High Availability - Publisher Fan-out & Subscriber Failover

# 1. Start Mosquitto brokers (Primary, Secondary, Tertiary)
Write-Host "Starting 3 Mosquitto brokers..." -ForegroundColor Cyan
docker-compose up -d

# 2. Build the applications
Write-Host "Building Publisher..." -ForegroundColor Cyan
cd publisher
go build -o publisher.exe .
cd ..

Write-Host "Building Subscriber..." -ForegroundColor Cyan
cd subscriber
go build -o subscriber.exe .
cd ..

# 3. Instructions
Write-Host "`n--- Multi-Broker HA Test Instructions ---" -ForegroundColor Yellow
Write-Host "For the best experience, open TWO new terminal windows."
Write-Host "1. In Terminal A, run: ./subscriber/subscriber.exe"
Write-Host "2. In Terminal B, run: ./publisher/publisher.exe"
Write-Host "3. Observe: Publisher connects to ALL 3 brokers. Subscriber connects to ONE."
Write-Host "4. KILL the broker the subscriber is using (e.g., docker stop mosquitto-primary)."
Write-Host "5. Watch Subscriber. It detects failure and moves to Secondary/Tertiary."
Write-Host "6. CRITICAL: Message flow RESUMES immediately because Publisher is already sending there."
Write-Host "7. To reset: docker start mosquitto-primary"
