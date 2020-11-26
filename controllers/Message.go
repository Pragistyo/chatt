package controllers


import (
	"log"
	"net/http"
	db "github.com/Pragistyo/chatt/db"
)

func PostMessage(w http.ResponseWriter,r *http.Request){
	log.Println("======= post message")
	conn := db.Connect()
	defer conn.Close()
}

func GetMessagesChatRoom(w http.ResponseWriter,r *http.Request){

}

func UpdateMessagesRead(w http.ResponseWriter,r *http.Request){

}