postgres:
	docker run --name postgres-bank -e POSTGRES_PASSWORD=handsome2022 -p 2345:5432 -d postgres:12-alpine
createdb:
	docker exec -it postgres-bank createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres-bank dropdb --username=postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:handsome2022@localhost:2345/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:handsome2022@localhost:2345/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
