SHELL=cmd.exe
APP_BINARY=shopService.exe

up:
	@echo "Starting Docker images On background..."
	docker-compose up -d
	@echo "Docker images started!"

run:
	@echo "Starting Docker images..."
	docker-compose up
	@echo "Docker images started!"

clean:
	cmd /c if exist ".\data\db-data" cmd /c rmdir /Q /S ".\data\db-data"
	@echo "Clean and Build"
	docker-compose down -v
	docker-compose build --no-cache
	docker-compose up --force-recreate
	@echo "Clean Build Done!"

down:
	@echo "Stopping docker compose..."
	docker-compose down -v
	@echo "Done!"

tidy:
	go mod tidy
	go mod vendor

build:
	@echo "Building Service..."
	chdir ..\shop-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${APP_BINARY} .
	@echo "Done!"

