postgres:
	docker run --name postgres13 5432:5342 -e POSTGRES_USER=postgres12 -e POSTGRES_PASSWORD=secret -d postgres:14.3-alpine

createdb:
	docker exec -it postgres13 createdb --username=postgres12 --owner=postgres12 lostandfound

dropdb:
	docker exec -it postgres13 dropdb --username=postgres12 lostandfound

migrateup:
	migrate -path db/migration  -database "postgresql://postgres12:secret@localhost:5342/lostandfound?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration  -database "postgresql://postgres12:secret@localhost:5342/lostandfound?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server