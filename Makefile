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