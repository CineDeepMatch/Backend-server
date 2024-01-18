DB_URL=postgresql://dat:secret@localhost:5432/cine-deep-match?sslmode=disable

postgres:
	docker run --name postgres-cdm --network cdm-network -p 5432:5432 -e POSTGRES_USER=dat -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres-cdm createdb --username=dat --owner=dat cine-deep-match

dropdb:
	docker exec -it postgres-cdm dropdb --username=dat cine-deep-match

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	rm -f db/sqlc/*.sql.go
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
dockerserver:
	docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://dat:secret@postgres:5432/cine-deep-match?sslmode=disable" simplebank:latest
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/CineDeepMatch/Backend-server/db/sqlc Store
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
proto: 
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=cdm \
	proto/*.proto
evans:
	evans --host localhost --port 9090 -r repl

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock migrateup1 migratedown1 new_migration dockerserver proto evans

	