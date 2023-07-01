postgres: 


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
