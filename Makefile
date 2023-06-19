postgres: 


migratUp: 
   migrate -path ./db/migration -database "postgres://postgres:postgres@localhost:5435/simple_bank?sslmode=disable" up 
migrateDown : 
	migrate -path ./db/migration -database "postgres://postgres:postgres@localhost:5435/simple_bank?sslmode=disable" down

sqlc: 
	docker run --rm -v "C:\Users\Lenovo\Desktop\go_project:/src" -w /src kjconroy/sqlc generate 