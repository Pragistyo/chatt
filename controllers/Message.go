package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	// "database/sql"
	// "reflect"
	"context"

	db "chatt/db"
	models "chatt/models"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)
type ResponseMessages struct {
	Message		string					`json:"message"` 
	Status		int32					`json:"status"` 
	UpdatedRead int64					`json:"updated_read_count"`
	Data		[]models.Message		`json:"Users"` 
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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respString, _ := json.Marshal( 
		Response{ "Message": "success post message", "Status": 201, "id_message": id } )
	w.Write([]byte  (respString))
	return

}

func GetMessagesChatRoom(w http.ResponseWriter,r *http.Request){
	log.Println("Hello")
	params := mux.Vars(r)
	chat_room_name := params["chat_room_name"]
	opposite_user_id := params["opposite_user_id"]
	log.Println("====== chat_room_name =======", chat_room_name) // should move to logger
	log.Println("====== opposite_user_id =======", opposite_user_id) // shoul mode to logger

	conn := db.Connect();
	defer conn.Close()

	var msg  models.Message
	var arr_msg []models.Message
	var countMsgRead int64

	countMsgRead = UpdateMessagesRead(conn, chat_room_name, opposite_user_id)

	var query string = "SELECT * FROM Message WHERE chat_room_name=$1 ORDER BY sent_time ASC"
	rows, err := conn.Query( context.Background(), query,   chat_room_name)

	if err!=nil {
		log.Println("hey there is error get message: ======> ",err)
		ResponseError(w, "Error retrieve data message", 400, http.StatusBadRequest, err)
	   	return
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&msg.Id, &msg.Message, & msg.SentTime, &msg.ReadTime,  &msg.UserId, &msg.ChatRoomName )
		err != nil {
			log.Println("Error Scan Messages list: ===> ",err.Error())
			ResponseError(w, "Error Scan Data", 400, http.StatusBadRequest, err)
			return
		} else {
			arr_msg = append(arr_msg, msg)
		}
	}
	err = rows.Err()
	if err != nil {
		ResponseError(w, "Error iteration message", 400, http.StatusBadRequest, err)
		return
	}

	if len(arr_msg) == 0 {
		w.WriteHeader(http.StatusNotFound)
			var respMsg ResponseMessages
			respMsg.Message = "No Message Yet"
			respMsg.Status = 404
			respMsg.Data = arr_msg
			respMsg.UpdatedRead = countMsgRead
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respMsg)
		return 
	}

	var respMsg ResponseMessages
	respMsg.Message = "success get messages"
	respMsg.Status = 200
	respMsg.Data = arr_msg
	respMsg.UpdatedRead = countMsgRead

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respMsg)
	return
}


func UpdateMessagesRead(conn *pgxpool.Pool, chat_room_name string, opposite_user_id string) int64 {
	log.Println("updatesMessage")
	var rowsAffected int64

	var sqlStatement string=`
	UPDATE Message 
	SET read_time =$1
	WHERE chat_room_name = $2 AND user_id = $3 AND read_time is NULL
	`
	log.Println("====== hello ====== ", opposite_user_id)
	resUpd, err := conn.Exec(context.Background(), sqlStatement, time.Now(), chat_room_name, opposite_user_id)
	if err!= nil {
		log.Println(" ===== Error update read message =====: ", err)
		return 0
	}
	rowsAffected = resUpd.RowsAffected()
	return rowsAffected
}