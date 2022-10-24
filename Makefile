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
