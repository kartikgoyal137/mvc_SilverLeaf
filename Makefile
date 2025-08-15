IMAGE_NAME=mvc-silverleaf
CONTAINER_NAME=silverleaf-app

build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .

run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file ./mvc_backend/.env --name $(CONTAINER_NAME) $(IMAGE_NAME)

stop:
	@echo "Stopping and removing container..."
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true
