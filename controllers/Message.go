package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	// "database/sql"
	// "reflect"
	"context"

	db "github.com/Pragistyo/chatt/db"
	"github.com/Pragistyo/chatt/models"
	"github.com/gorilla/mux"
)
type ResponseMessages struct {
	Message		string					`json:"message"` 
	Status		int32					`json:"status"` 
	Data		models.Message				`json:"Users"` 
}

func Huba(w http.ResponseWriter, r *http.Request){
	log.Println("Hello")
	return
}


func PostMessage(w http.ResponseWriter,r *http.Request){
	conn := db.Connect()
	defer conn.Close()

	err := r.ParseMultipartForm(64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte  ("error parse input"))
		return
	}

	var message string = r.FormValue("message")
	var chat_room_name string = r.FormValue("chat_room_name")
	var user_post_id string = r.FormValue("user_post_id")
	var sentTime time.Time = time.Now()

	var id int32

	sqlStatement := `
	INSERT INTO Message (message, sent_time, read_time, user_id, chat_room_name)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	err = conn.QueryRow(context.Background(), sqlStatement, message, sentTime, nil, user_post_id, chat_room_name).Scan(&id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	   log.Println("hey there is error : ======> ", err)
	   return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	respString, _ := json.Marshal( 
		Response{ "Message": "success post message", "Status": 201, "id_message": id } )
	w.Write([]byte  (respString))
	return

}

func GetMessagesChatRoom(w http.ResponseWriter,r *http.Request){ // array, plural
	log.Println("Hello")
	params := mux.Vars(r)
	chat_room_name := params["chat_room_name"]

	conn := db.Connect();
	defer conn.Close()

	var msg  models.Message
	var arr_msg []models.Message

	var query string = "SELECT * FROM Message WHERE chat_room_name=$1 ORDER BY sent_time ASC"
	rows, err := conn.Query( context.Background(), query,   chat_room_name)

	if err!=nil {
		log.Println("hey there is error get message: ======> ",err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
			respString, _ := json.Marshal(
				Response{ "Message": "Error retrieve data message", "Status": 400 } )
		w.Write([]byte  (respString))
	   	return
	}

	defer rows.Close()
	count := 0
	for rows.Next() {
		if err := rows.Scan(&msg.Id, &msg.Message, & msg.SentTime, &msg.ReadTime,  &msg.UserId, &msg.ChatRoomName )
		err != nil {
			log.Println("Error Scan Messages list: ===> ",err.Error())
			 if count == 0 {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
					respString, _ := json.Marshal(
						Response{ "Message": "No message yet", "Status": 404, "error": err } )
			 }else{
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "application/json")
					respString, _ := json.Marshal(
						Response{ "Message": "Error Scan Data", "Status": 400, "error": err } )
			 }
			return
		} else {
			arr_msg = append(arr_msg, msg)
			count += 1 
		}
	}

	var respMsg ResponseMessages
	respMsg.Message = "success get messages"
	respMsg.Status = 200
	respMsg.Data = msg

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respMsg)

	
}


func UpdateMessagesRead(w http.ResponseWriter,r *http.Request){
	log.Println("updatesMessage")
	return
}