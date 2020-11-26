package main

import (
	"net/http"
	"os"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	controllers "github.com/pragistyo/chatt/controllers"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r:= mux.NewRouter()
	api := r.PathPrefix("/go-chat/api/v1").Subrouter()


	//POST
	api.HandleFunc("/login/",controllers.Login).Methods(http.MethodPost)
	api.HandleFunc("/create-user/",controllers.CreateUser).Methods(http.MethodPost)
	api.HandleFunc("/chat-room/",controllers.CreateChatRoom).Methods(http.MethodPost)
	api.HandleFunc("/message/",controllers.PostMessage).Methods(http.MethodPost)

	//GET
	api.HandleFunc("/message-chat-room/{chat_room_name}/",controllers.GetMessagesChatRoom).Methods(http.MethodGet)
	api.HandleFunc("/conversation-card/",controllers.CardConversation).Methods(http.MethodGet)

	//PATCH
	api.HandleFunc("/update-read-message/{chat_room_name}/{user_id}/",controllers.UpdateMessagesRead).Methods(http.MethodPatch)

	PORT := os.Getenv( "PORT" )
	log.Fatal(http.ListenAndServe(PORT, r))
}

	