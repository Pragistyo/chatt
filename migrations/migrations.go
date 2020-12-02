package main

import(
	// "net/http"
	"log"
	// "reflect"
	"context"

	db "chatt/db"
	"github.com/joho/godotenv"

	
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	log.Println("start migrations ====== ")
	// dropTable()
	createUser()
	createChatRoom()
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
		chat_room_name VARCHAR(120) UNIQUE PRIMARY KEY,
		user_id_1 integer REFERENCES UserChat (user_id),
		user_id_2 integer REFERENCES UserChat (user_id)
	) `

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
	chat_room_name VARCHAR(120) REFERENCES ChatRoom (chat_room_name)
	)`
	tableChatRoom, err := conn.Exec(context.Background(), query)

	if err!=nil {
		panic(err)
	}

	log.Println("==========")
	log.Println(tableChatRoom)
	return
}

func dropTable(){
	conn := db.Connect()
	defer conn.Close()

	var query string = `
	DROP TABLE Message, ChatRoom, UserChat;
	`

	dropTable, err := conn.Exec(context.Background(), query)

	if err!=nil {
		log.Println("====== error")
		log.Println(err)
		panic(err)
	}

	log.Println("==========")
	log.Println(dropTable)
	return

}

func alterChatRoom(){
	conn := db.Connect()
	defer conn.Close()

	var query string = `
	ALTER TABLE chatroom RENAME COLUMN user_id to user_id_1;
	ALTER TABLE chatroom ADD COLUMN user_id_2 integer REFERENCES UserChat (user_id)
	`
	alterTable, err := conn.Exec(context.Background(), query)

	if err!=nil {
		log.Println("====== error")
		log.Println(err)
		panic(err)
	}

	log.Println("==========")
	log.Println(alterTable)
	return
}