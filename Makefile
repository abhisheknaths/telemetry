NAME := $(shell basename $(CURDIR))
.PHONY: build
build: 
	docker compose build

.PHONY: serve
serve: build
	@echo "Starting ${NAME}"
	docker compose up -d

.PHONY: down
down:
	docker compose down