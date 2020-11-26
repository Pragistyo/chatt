package controllers


import (
	"log"
	"net/http"
	db "github.com/Pragistyo/chatt/db"
)

func CardConversation(w http.ResponseWriter,r *http.Request){
	log.Println( "======= Card Conversation ")
	conn := db.Connect()
	defer conn.Close()
}