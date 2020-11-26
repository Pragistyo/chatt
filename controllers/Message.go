package controllers

import (
	"log"
	"net/http"
	"time"
	"encoding/json"

	db "github.com/Pragistyo/chatt/db"
	"github.com/Pragistyo/chatt/models"
)

type Response map[string]interface{}

func PostMessage(w http.ResponseWriter,r *http.Request){
	conn := db.Connect()
	defer conn.Close()

	err := r.ParseMultipartForm(64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte  ("error parse input"))
		panic(err)
		return
	}

	var message string = r.FormValue("message")
	var chat_room_name string = r.FormValue("chat_room_name")
	var user_post_id string = r.FormValue("user_post_id")
	var sentTime time.Time = time.Now()
	var readTime time.Time = nil

	var msg models.Message

	sqlStatement := `
	INSERT INTO Users (message, sent_time, read_time, user_id, chat_room_name)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	err = conn.QueryRow(context.Background(), sqlStatement, message, sentTime, readTime, user_post_id, chat_room_name).Scan(&id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	   log.Println(err)
	   return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	respString, _ := json.Marshal( 
		Response{ "Message": "success post message", "Status": 201, "id_message": id } )
	w.Write([]byte  (respString))
	return

}

func GetMessagesChatRoom(w http.ResponseWriter,r *http.Request){

}

func UpdateMessagesRead(w http.ResponseWriter,r *http.Request){

}