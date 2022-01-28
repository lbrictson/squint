build:
	@echo "Building binary"
	@go build -o build/squint cmd/main.go
	@echo "Binary available in build/"

test:
	@echo "Running tests"
	@echo "Ensuring compose is down"
	@docker-compose down
	@echo "Starting postgres test database"
	@docker run --name squint_test_pg -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres docker.io/postgres:14
	@go test ./... -cover
	@echo "Cleaning up test database"
	@docker rm --force squint_test_pg

clean:
	@echo "Cleaning tmp files"
	@rm ./build/*
	@echo "Done cleaning"

migrate:
	@echo "Enter migration name:"; \
 	read MIGRATENAME; \
	migrate create -ext sql -dir sql/migrations -seq $$MIGRATENAME;

run:
	@docker-compose up -d --build

down:
	@docker-compose down

watch:
	@docker-compose up -d --build
	@air -c .air.toml
