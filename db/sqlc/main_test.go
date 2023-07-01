package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	util "github.com/Hedi-zouaoui/go_project/utils"
	_ "github.com/lib/pq"
) 

var testQueries *Queries
var testDB *sql.DB
func TestMain(m *testing.M){
	var err error
	config , err := util.LoadConfig("../..")
	testDB , err = sql.Open(config.DBDriver , config.DBSource)
	if err!= nil{
		log.Fatal("cannot connect" , err)
	}
testQueries = New(testDB)
os.Exit(m.Run()) 
}