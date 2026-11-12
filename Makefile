include .env
export

export PROJECT_ROOT = $(shell pwd)

COMPOSE ?= docker compose

.PHONY: run docker-ps db-up db-down db-cleanup db-exec env-port-forward env-port-close migrate-create migrate-action migrate-up migrate-down

run:
	@go run cmd/api/main.go

docker-ps:
	@$(COMPOSE) ps -a

db-up:
	@$(COMPOSE) up -d postgresql

db-down:
	@$(COMPOSE) down postgresql

db-cleanup:
	@$(COMPOSE) down -v postgresql

db-exec:
	@$(COMPOSE) exec -it postgresql psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

env-port-forward:
	@$(COMPOSE) up -d port_forwarder

env-port-close:
	@$(COMPOSE) down port_forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "'seq' parameter is required"; \
		exit 1; \
	fi; \
	$(COMPOSE) run --rm migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "'action' parameter is required"; \
		exit 1; \
	fi; \
	$(COMPOSE) run --rm migrate \
		-path migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@beatbase_postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down