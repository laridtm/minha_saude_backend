.PHONY: env env-stop start

COMPOSE_FILE = build/docker-compose.yml
MONGO=mongodb://localhost:27017

env:
	docker compose -f $(COMPOSE_FILE) up -d

env-stop:
	docker compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

start:
	MONGO_URL=$(MONGO) \
	go run cmd/minha_saude/main.go

build:
	go build cmd/minha_saude/main.go