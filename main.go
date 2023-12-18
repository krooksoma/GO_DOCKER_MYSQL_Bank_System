package main

import (
	"bank_system/db"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	
	
	newUser, _ := db.NewAccount("Telomero", 1233, "EUR")
	db.InsertAccount(newUser)
	db.FetchAccountById(9)

}
