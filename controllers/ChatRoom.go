package controllers


import (
	"log"
	"net/http"
	db "github.com/Pragistyo/chatt/db"
)

func CreateChatRoom(w http.ResponseWriter,r *http.Request){
	log.Println( "======= Create Chat Room ")
	conn := db.Connect()
	defer conn.Close()
}