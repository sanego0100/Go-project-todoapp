include .env
export


env-up:
	@docker compose up todoapp-pg -d
env-down:
	@docker compose down todoapp-pg
env-cleanup:
	@read -p "Clean up all volume? [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-pg && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "over"; \
	else \
		echo "operation cancelled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "add sequens. Example migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-pg-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"
	
migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "add action"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-pg-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-pg:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

env-forward:
	@docker compose up -d port-forwarder

todoapp-run:
	@go run cmd/todoapp/main.go