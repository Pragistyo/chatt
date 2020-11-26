package controllers


import (
	"log"
	"net/http"
	db "github.com/Pragistyo/chatt/db"
)

func CreateChatRoom(w http.ResponseWriter,r *http.Request){
	conn := db.Connect()
	defer conn.Close()
	
	err := r.ParseMultipartForm(4096) // max memory 4mb
	if err != nil {
		panic(err)
	}

	var user_email string = r.FormValue("user_email")
	var user_id string = r.FormValue("user_id")
	var oppose_user_id string = r.FormValue("oppose_id")
	var oppose_user_email = r.FormValue("oppose_user_email")

	var chat_room_name_1 string = user_email + "-" + oppose_user_email
	var chat_room_name_2 string = oppose_user_email + "-" + user_email

	var possibleName [2]string
	possibleName[0] = chat_room_name_1
	possibleName[1] = chat_room_name_2
	// validation
	var chat_room_name string = checkChatRoomExist(possibleName)
	//create chat room, then create junction dengan user1, then create dengan user 2
	// transaction query
	// kalo room udah ada , return response udah ada room,
	//		front end langsung redirect screen chat room


	sqlStatement = `INSERT INTO ChatRoom (user_id, chat_room_name)
	VALUES ($1, $2)
	RETURNING id
	`
	


}

func checkChatRoomExist( possibleName [2]string) string {

}