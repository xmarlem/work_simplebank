.PHONY: db migrate.up migrate.down

DB_PASSWORD=test
DB_HOST=localhost
DB_PORT=5433
DB=simple_bank
DB_USER=root

POSTGRESQL_URL="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB)?sslmode=disable"

db.run:
	docker run --name postgres -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:12-alpine
	sleep 1
	docker exec -it postgres createdb --username=$(DB_USER) --port=5432 --owner=$(DB_USER) $(DB)
#	psql -h localhost -U postgres -w -c "create database simple_bank;" 

db.create:
	docker exec -it postgres createdb --username=$(DB_USER) --port=5432 --owner=$(DB_USER) $(DB)
db.drop:
	docker exec -it postgres dropdb --username=$(DB_USER) $(DB)

db.exec:
	docker exec -it postgres psql -U root $$(DB)

db.stop:
	@docker stop postgres

db.start:
	@docker start postgres

db.remove:
	@docker kill postgres && docker rm postgres


migrate.init:
	migrate create -ext sql -dir db/migration -seq init_schema

migrate.up:
	migrate -database=$(POSTGRESQL_URL) -path=db/migration -verbose up
migrate.down:
	migrate -database=$(POSTGRESQL_URL) -path=db/migration -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...