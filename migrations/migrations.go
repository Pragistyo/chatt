package main

import(
	"net/http"
	"log"
	"reflect"
	"context"

	db "github.com/Pragistyo/chatt/db"
	
)

func main() {
	


}

func createUser(){
	conn := db.Connect()
	defer conn.Close()

	var query string = `CREATE TABLE IF NOT EXISTS UserChat(
		user_id SERIAL PRIMARY KEY,
		email VARCHAR(50) UNIQUE,
		name VARCHAR(50)
	) `

	tableUser, err := conn.Exec(context.Background(), query)

	if err!=nil {
		panic(err)
	}

	log.Println("==========")
	log.Println(tableUser)
}