postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root todo-db

dropdb:
    docker exec -it postgres12 dropdb todo-db

migrateup:
    migrate -path ./schema -database "postgresql://root:secret@localhost:5432/todo-db?sslmode=disable" up

migratedown:
    migrate -path ./schema -database "postgresql://root:secret@localhost:5432/todo-db?sslmode=disable" down

.PHONY: postgres createdb dropdb migrateup migratedown