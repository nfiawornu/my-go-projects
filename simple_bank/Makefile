createdb:
	docker exec -it postgres createdb --username=postgres_nobel --owner=postgres_nobel simple_bank
dropdb:
	docker exec -it postgres dropdb --username=postgres_nobel --owner=postgres_nobel simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/nfiawornu/my-go-projects/simple_bank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test server mock