ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: default run build docs clean \
        atlas-inspect atlas-diff atlas-apply atlas-status

default: run

debug-env:
	@echo "DB_URL=$(DB_URL)"

run:
	@swag init -g cmd/main.go
	@go run cmd/main.go

build:
	@go build -o $(APP_NAME) cmd/main.go

docs:
	@swag init -g cmd/main.go

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs

# Atlas (Migrations)
atlas-inspect:
	@atlas schema inspect --env $(ATLAS_ENV) --url "env://src"

atlas-diff:
	@atlas migrate diff --env $(ATLAS_ENV)

atlas-apply:
	@atlas migrate apply --url "$(DB_URL)"

atlas-status:
	@atlas migrate status --url "$(DB_URL)"

