build:
	@go build -o bin/supabase-testProject1 main.go

test:
	@go test -v ./...

run: build
	@./bin/supabase-testProject1

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
