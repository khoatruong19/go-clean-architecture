up:
	docker-compose up -d

down:
	docker-compose up -d

createdb:
	docker exec -it stuhub-be-db-1 createdb --username=postgres --owner=postgres stuHub

dropdb:
	docker exec -it stuhub-be-db-1 dropdb --username=postgres stuHub

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/stuHub?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/stuHub?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: down createdb dropdb migrateup migratedown sqlc