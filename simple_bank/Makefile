createdb:
	docker exec -it postgres createdb --username=postgres_nobel --owner=postgres_nobel simple_bank
dropdb:
	docker exec -it postgres dropdb --username=postgres_nobel --owner=postgres_nobel simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/nfiawornu/my-go-projects/simple_bank/db/sqlc Store

proto:
	rm -rf pb/*.go
	protoc -I="C:/Users/NOBEL FIAWORNU/go/protoc-32.1-win64/include"  -I=./proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
evans:
	evans --reflection --proto "$(pwd)/proto/*.proto"  --port 9092 -r repl

.PHONY: createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 proto evans