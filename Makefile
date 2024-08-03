postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root tododb

dropdb:
    docker exec -it postgres12 dropdb tododb

migrateup:
    migrate -path ./schema -database "postgresql://root:secret@localhost:5432/tododb?sslmode=disable" up

migratedown:
    migrate -path ./schema -database "postgresql://root:secret@localhost:5432/tododb?sslmode=disable" down

.PHONY: postgres createdb dropdb migrateup migratedown