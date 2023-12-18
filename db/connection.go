package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func DbConnection() *sql.DB{
	db, err := sql.Open("mysql", "root:defaria123@tcp(localhost:3306)/banking_system")

	if err != nil{
		panic(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db;
}

