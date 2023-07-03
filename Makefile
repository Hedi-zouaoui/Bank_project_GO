postgres: 

createMigrate: 
	migrate create -ext sql -dir ./db/migration -seq <name of the migrate>
migratUp: 
   migrate -path ./db/migration -database "postgres://postgres:postgres@localhost:5435/simple_bank?sslmode=disable" up 
migrateDown : 
	migrate -path ./db/migration -database "postgres://postgres:postgres@localhost:5435/simple_bank?sslmode=disable" down


migratUp1: 
   migrate -path ./db/migration -database "postgres://postgres:postgres@localhost:5435/simple_bank?sslmode=disable" up1 
migrateDown1 : 
	migrate -path ./db/migration -database "postgres://postgres:postgres@localhost:5435/simple_bank?sslmode=disable" down1

sqlc: 
	docker run --rm -v "C:\Users\Lenovo\Desktop\go_project:/src" -w /src kjconroy/sqlc generate 

server : 
 	go run main.go

mock : 
	mockgen --build_flags=--mod=mod  -destination db/mock/store.go github.com/Hedi-zouaoui/go_project/db/sqlc store
proto : 
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative proto/*.proto 

evans : 
	evans -r -p 9091 



	