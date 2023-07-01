package main

import (
	"database/sql"
	"log"

	"github.com/Hedi-zouaoui/go_project/api"
	db "github.com/Hedi-zouaoui/go_project/db/sqlc"
	util "github.com/Hedi-zouaoui/go_project/utils"
	_ "github.com/lib/pq"
)



func main() {
	config , err := util.LoadConfig(".")
	if err != nil {
log.Fatal("cannot log config " , err )
	}
	conn, err := sql.Open(config.DBDriver , config.DBSource)
	if err != nil {
		log.Fatal("cannot connect", err)
	}
	store := db.NewStore(conn)
	server , err := api.NewServer(config , store)
	if  err!=nil{
		log.Fatal("cannot create new server ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server " , err)
	}

}
