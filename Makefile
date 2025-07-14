createdb:
	docker exec -it postgres createdb --username=postgres_nobel --owner=postgres_nobel simple_bank
dropdb:
	docker exec -it postgres dropdb --username=postgres_nobel --owner=postgres_nobel simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable" -verbose down



.PHONY: createdb dropdb migrateup migratedown