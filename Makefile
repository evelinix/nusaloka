# Makefile untuk menjalankan dan rebuild container Docker Compose

# Lokasi untuk service
ACCOUNT_SERVICE_DIR=cmd/account
TRIP_SERVICE_DIR=cmd/trip
FINANCE_SERVICE_DIR=cmd/finance
MEDIA_SERVICE_DIR=cmd/media
REALTIME_SERVICE_DIR=cmd/realtime
GATEWAY_SERVICE_DIR=cmd/gateway

# Start all services with Docker Compose (without rebuild)
run:
	@echo "Starting all services with Docker Compose..."
	docker-compose up -d  # Jalankan di background

run-db:
	@echo "Starting Postgres Database with Docker Compose..."
	docker-compose up -d db

run-redis:
	@echo "Starting Redis with Docker Compose..."
	docker-compose up -d redis

# Jalankan hanya Account Service
run-account:
	@echo "Starting Account Service with Docker Compose..."
	docker-compose up -d account-service

# Jalankan hanya Trip Service
run-trip:
	@echo "Starting Trip Service with Docker Compose..."
	docker-compose up -d trip-service

# Jalankan hanya Finance Service
run-finance:
	@echo "Starting Finance Service with Docker Compose..."
	docker-compose up -d finance-service

# Jalankan hanya Media Service
run-media:
	@echo "Starting Media Service with Docker Compose..."
	docker-compose up -d media-service

# Jalankan hanya Realtime Service
run-realtime:
	@echo "Starting Realtime Service with Docker Compose..."
	docker-compose up -d realtime-service

# Jalankan hanya API Gateway
run-gateway:
	@echo "Starting API Gateway with Docker Compose..."
	docker-compose up -d gateway

# Rebuild dan jalankan semua service
rebuild:
	@echo "Rebuilding and restarting all containers..."
	docker-compose down --volumes --remove-orphans  # Hapus container dan volumes yang lama
	docker-compose up --build -d  # Build ulang container dan jalankan di background

.PHONY: stop # Hentikan semua service yang berjalan
stop:
	@echo "Stopping all containers..."
	docker-compose down  # Hentikan dan hapus container

.PHONY: build # Build ulang service tanpa menjalankan container
build:
	@echo "Building all services without running them..."
	docker-compose build  # Build ulang container tanpa menjalankan

.PHONY: install-air # Install air (optional)
install-air:
	@echo "Installing air..."
	go install github.com/cosmtrek/air@latest

.PHONY: clean # Clean up containers, networks, and volumes
clean:
	@echo "Cleaning all containers..."
	docker-compose down --volumes --rmi all

.PHONY: help # Menampilkan help tentang semua target yang tersedia
help:
	@printf "Makefile commands:\n"
	@printf "  run           - Start all services with Docker Compose\n"
	@printf "  run-account   - Start only Account Service with Docker Compose\n"
	@printf "  run-trip      - Start only Trip Service with Docker Compose\n"
	@printf "  run-finance   - Start only Finance Service with Docker Compose\n"
	@printf "  run-media     - Start only Media Service with Docker Compose\n"
	@printf "  run-realtime  - Start only Realtime Service with Docker Compose\n"
	@printf "  run-gateway   - Start only API Gateway with Docker Compose\n"
	@printf "  rebuild       - Rebuild and restart all services\n"
	@printf "  stop          - Stop all running services\n"
	@printf "  build         - Build all services without running\n"
	@printf "  clean         - Clean all services \033[31mDo not use if not needed\033[0m\n"
	@printf "  install-air   - Install the 'air' live-reload tool\n"
	@printf "  help          - Display this help message\n"


.PHONY: dev
dev:
	@echo "🚀 Starting $(SERVICE) service with Air"
	@air -root . -tmp_dir "tmp" \
	-build.cmd "go build -o ./tmp/$(SERVICE)-service ./cmd/$(SERVICE)/main.go" \
	-build.bin "./tmp/$(SERVICE)-service" \
	-log.time true \
	-misc.clean_on_exit true \
	-screen.clear_on_rebuild true \
	-build.exclude_dir "docker" 

jwt-key:
	@echo "Start Creating Private KEY"
	@openssl ecparam -name secp521r1 -genkey -noout -out ./keys/es512-private.pem
	@echo "Start Generating Public KEY"
	@ openssl ec -in ./keys/es512-private.pem -pubout -out ./keys/es512-public.pem