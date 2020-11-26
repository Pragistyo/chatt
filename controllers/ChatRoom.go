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

	var roomName string = r.FormValue("roomName")
	var user_id string = r.FormValue("user_id")
	var oppose_user_id string = r.FormValue("oppose_id")


	//create chat room, then create junction dengan user1, then create dengan user 2
	// transaction query
	// kalo room udah ada , return response udah ada room,
	//		front end langsung redirect screen chat room


	sqlStatement = `
	`
	



}