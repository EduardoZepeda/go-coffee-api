## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## docs/generate: Generate api docs using swaggo
.PHONY: docs/generate
docs/generate:
	@echo 'Generating API documentation using swaggo'
	cd api/handler && swag init --parseDependency --d ./,../../handlers -o ../../docs

## run/dev: Run development server
.PHONY: run/dev
run/dev:
	@echo 'Running server in development mode'
	vercel dev

## migrate/new name=$1: create a new database migration
.PHONY: migrate/new
migrate/new:
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## migrate/up: Run the appropiate up migrations and set the database in its final state
.PHONY: migrate/up
migrate/up:
	@echo 'Running migrations'
	migrate -path=./migrations -database ${GO_COFFEE_API_DSN} up

## migrate/down: Run the appropiate up migrations and set the database in its intial state
.PHONY: migrate/down
migrate/down:
	@echo 'Running migrations'
	migrate -path=./migrations -database ${GO_COFFEE_API_DSN} down