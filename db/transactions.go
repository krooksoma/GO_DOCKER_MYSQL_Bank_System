package db

import (
	"errors"
	"fmt"
	"log"
)

type Account struct{
	id int32 
	owner string
	balance int32
	currency string
	created_at string
}

var db = DbConnection()

func FetchAll() {

	results, err := db.Query(`SELECT * FROM account`)

	if err != nil{
		log.Fatal(err)
	}

	for results.Next(){
		var account Account

		err = results.Scan(&account.id, &account.owner, &account.balance, &account.currency, &account.created_at)
		if err != nil {
			panic(err.Error())
		}


		fmt.Printf(" account ID: %v account Owner: %v account balance: $%v %v.", account.id, account.owner, account.balance, account.currency)
		fmt.Println("")
	}
	
}

func InsertAccount(new *Account){
	
	_, err := db.Query(`INSERT INTO account( owner, balance, currency ) VALUE (?, ?, ?)`, new.owner, new.balance, new.currency)

	if err != nil{
		log.Fatal(err)
	}

	log.Printf("New Record inserted into account DB: \n Owner: %v \n Balance: %v, Currency: %v", new.owner, new.balance, new.currency)


} 

func NewAccount(ownerV string, balanceV int32, currencyV string) (*Account, error){
	if ownerV == "" || currencyV == ""{
		return nil, errors.New("Given fields are empty")
	}

	return &Account{
		owner: ownerV,
		balance: balanceV,
		currency: currencyV,
	}, nil

}

func FetchAccountById(idNumber int32) {
	result, err := db.Query(`SELECT * FROM account WHERE id = ?`, idNumber);

	if err != nil{
		log.Fatal(err)
	}

	for result.Next(){
		var account Account

		err = result.Scan(&account.id, &account.owner, &account.balance, &account.currency, &account.created_at)
		if err != nil {
			panic(err.Error())
		}


		fmt.Printf(" account ID: %v account Owner: %v account balance: $%v %v.", account.id, account.owner, account.balance, account.currency)
		fmt.Println("")
	}
}