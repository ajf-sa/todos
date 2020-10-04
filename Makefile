createdb:
	docker exec -it db01 createdb --username=admin --owner=admin todos-api-1

dropdb:
	docker exec -it db01 dropdb todos-api-1 -U admin

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir server/db/migration -seq $$name

migrate:
	migrate -path server/db/migration -database "postgresql://admin:secret@localhost:5432/todos-api-1?sslmode=disable" -verbose up

rollback:
	migrate -path server/db/migration -database "postgresql://admin:secret@localhost:5432/todos-api-1?sslmode=disable" -verbose down

sqlc:
	sqlc generate

build:
	docker build . -t go-dock

run:
	docker run -v /home/north/projects/todos/public:/dist/public -p 3000:3000 go-dock 

.PHONY: createdb dropdb migrate rollback build run sqlc migration