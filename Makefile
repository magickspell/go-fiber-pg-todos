# Makefile

# Команда для запуска Docker Compose
start:
	docker compose up -d

stop:
	docker compose down

reset:
	docker compose down
	docker rmi todo-go-pg-fiber-todo-go:latest