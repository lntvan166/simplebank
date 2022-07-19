postgres:
	docker run --name postgres-bank -e POSTGRES_PASSWORD=handsome2022 -p 2345:5432 -d postgres:12-alpine
createdb:
	docker exec -it postgres-bank createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres-bank dropdb --username=postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:handsome2022@localhost:2345/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://postgres:handsome2022@localhost:2345/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://postgres:handsome2022@localhost:2345/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://postgres:handsome2022@localhost:2345/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server: 
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/lntvan166/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
