build:
	@echo "Building services with Docker Compose..."
	docker compose build

up:
	@echo "Starting services with Docker Compose..."
	docker compose up -d

down:
	@echo "Stopping and removing services..."
	docker compose down

logs:
	@echo "Tailing logs..."
	docker compose logs -f app