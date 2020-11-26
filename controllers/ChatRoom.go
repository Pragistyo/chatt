package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	db "github.com/Pragistyo/chatt/db"
)

type ResponseBadRequest map[string]interface{}

func CreateChatRoom(w http.ResponseWriter,r *http.Request){
	conn := db.Connect()
	defer conn.Close()
	
	err := r.ParseMultipartForm(4096) // max memory 4mb
	if err != nil {
		panic(err)
	}

	var user_id string = r.FormValue("user_id")
	var oppose_user_id string = r.FormValue("oppose_id")
	var user_email string = r.FormValue("user_email")
	var oppose_user_email = r.FormValue("oppose_user_email")

	var chat_room_name_1 string = user_email + "-" + oppose_user_email
	var chat_room_name_2 string = oppose_user_email + "-" + user_email

	var possibleName [2]string

	possibleName[0] = chat_room_name_1
	possibleName[1] = chat_room_name_2
	
	log.Println("============bom bon")
	// validation
	if !checkChatRoomExist( possibleName){
		w.WriteHeader(http.StatusBadRequest)
		respString, _ := json.Marshal( ResponseBadRequest{ "Message": "duplicate chat_room_name", "Status": 400 } )
		w.Write([]byte  (respString))
		return
	}
	var chat_room_name string = possibleName[0]


	var sqlStatement string = `INSERT INTO ChatRoom (user_id_1, user_id_2, chat_room_name)
	VALUES ($1, $2, $3)
	RETURNING chat_room_name
	`
	err = conn.QueryRow(context.Background(), sqlStatement, user_id, oppose_user_id, chat_room_name).Scan(&chat_room_name)
	if err!=nil {
			log.Println("Error create chat room: ", err)
			w.WriteHeader(http.StatusBadRequest)
			respString, _ := json.Marshal( ResponseBadRequest{ "Message": "error create chat room", "Status": 400, "error":err } )
			w.Write([]byte  (respString))
			return
	}


	w.WriteHeader(http.StatusBadRequest)
	respString, _ := json.Marshal( ResponseBadRequest{ "Message": "success create chat room", "Status": 201, "chat_room_name": chat_room_name } )
	w.Write([]byte  (respString))
	return

}

func checkChatRoomExist( possibleName [2]string) bool {
	conn := db.Connect()
	defer conn.Close()
	

	type ResponseChatRoom struct {
		user_id	int32
		chat_room_name	string
	}
	var cr  ResponseChatRoom
	queryCheck := "SELECT chat_room_name FROM ChatRoom WHERE chat_room_name=$1 "

	// ============ cek possibleName[0] ============
		// user_1
	row := conn.QueryRow( context.Background(),queryCheck , possibleName[0] )
	err := row.Scan(&cr.chat_room_name )

	if err!=nil {
		// if recorded, do not replicate insert
		log.Println(" should be not found 0 ", err)
		return true
	}


	// ========== check possibleName[1] ============
		//user_1
	row = conn.QueryRow( context.Background(),queryCheck, possibleName[1])
	err = row.Scan(&cr.chat_room_name )

	if err!=nil {
		// if recorded, do not replicate insert
		log.Println(" should be not found 1 ", err)
		return true
	}
	

	//row exist
	return false

	


}