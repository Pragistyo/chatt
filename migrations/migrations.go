package main

import(
	// "net/http"
	"log"
	// "reflect"
	"context"

	db "github.com/Pragistyo/chatt/db"
	"github.com/joho/godotenv"

	
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	log.Println("start migrations ====== ")
	createUser()
	createChatRoom()
	createChatRoom_UserChat()
	createMessage()
	return
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
	return
}

func createChatRoom(){
	conn := db.Connect()
	defer conn.Close()

	var query string = `CREATE TABLE IF NOT EXISTS ChatRoom(
		chat_room_id SERIAL PRIMARY KEY,
		name VARCHAR(50) UNIQUE
	) `

	tableChatRoom, err := conn.Exec(context.Background(), query)

	if err!=nil {
		panic(err)
	}

	log.Println("==========")
	log.Println(tableChatRoom)
	return
}

func createChatRoom_UserChat(){
	conn := db.Connect()
	defer conn.Close()

	var query string = ` CREATE TABLE IF NOT EXISTS ChatRoom_UserChat(
	user_id integer REFERENCES UserChat (user_id),
	chat_room_id integer REFERENCES ChatRoom (chat_room_id)
	)`
	tableChatRoom, err := conn.Exec(context.Background(), query)

	if err!=nil {
		panic(err)
	}

	log.Println("==========")
	log.Println(tableChatRoom)
	return

}

func createMessage(){
	conn := db.Connect()
	defer conn.Close()

	var query string = ` CREATE TABLE IF NOT EXISTS Message(
	id SERIAL PRIMARY KEY,
	message  VARCHAR,
	sent_time TIMESTAMP NOT NULL,
	read_time TIMESTAMP,
	user_id integer REFERENCES UserChat (user_id),
	chat_room_id integer REFERENCES ChatRoom (chat_room_id)
	)`
	tableChatRoom, err := conn.Exec(context.Background(), query)

	if err!=nil {
		panic(err)
	}

	log.Println("==========")
	log.Println(tableChatRoom)
	return
}