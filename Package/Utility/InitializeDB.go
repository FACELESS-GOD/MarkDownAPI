package Utility

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DatabaseInstace *sql.DB

func InitialiseDatabaseConnection() {
	var ConnectionString string = "root:Admin@123@tcp/MarkDownAPIV2?charset=UTF8&parseTime=True&loc=Local"
	DB, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		panic(err)
	}
	DatabaseInstace = DB
}

func TerminateDatabaseConnection() {
	DatabaseInstace.Close()
}
