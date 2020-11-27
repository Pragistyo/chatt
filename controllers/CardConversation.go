package controllers


import (
	"log"
	"net/http"
	db "github.com/Pragistyo/chatt/db"
	models "github.com/Pragistyo/chatt/models"
)

func CardConversation(w http.ResponseWriter,r *http.Request){
	log.Println( "======= Card Conversation ")
	conn := db.Connect()
	defer conn.Close()

	query=`
	`

	rows, err := conn.Query( context.Background(), query,   chat_room_name)

	if err!=nil {
		log.Println("hey there is error get message: ======> ",err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
			respString, _ := json.Marshal(
				Response{ "Message": "Error retrieve data message", "Status": 400, "error":err } )
		w.Write([]byte  (respString))
	   	return
	}
	// needed
	// opposite user data
		// id, email, nama
	// last message
		// select distinct message (topmost)
		// where user_id = opposite, chat_room = chat_room
		// order by desc sent_time

	// select * chatroom where user_id_1 = user_id OR user_id_2 = user_id
	return user_id != user_id

}