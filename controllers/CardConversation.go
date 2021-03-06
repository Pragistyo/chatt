package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	db "chatt/db"
	models "chatt/models"
	"github.com/gorilla/mux"
)


type ResponseConvCard struct {
	Message		string					`json:"message"` 
	Status		int32					`json:"status"` 
	Data		[]models.ConversationCard		`json:"Data"` 
}

func CardConversation(w http.ResponseWriter,r *http.Request){
	log.Println( "======= Card Conversation ")
	params := mux.Vars(r)
	raw_user_id := params["user_id"]

	i, err := strconv.ParseInt(raw_user_id, 10, 32)
	if err != nil {
		log.Println(" ==== error parse params to int32 ========: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte  (" error parsing params "))
	}
	user_id := int32(i)

	conn := db.Connect()
	defer conn.Close()

	var rawConvCard models.ConversationCardRaw
	var arr_cardConvObj []models.ConversationCard

	var getQueryCardConv string = getQueryCardConv()
	rows, err := conn.Query( context.Background(), getQueryCardConv,   user_id)

	if err!=nil {
		log.Println("hey there is error get message: ======> ",err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
			respString, _ := json.Marshal(
				Response{ "Message": "Error retrieve conversations data", "Status": 400, "error":err } )
		w.Write([]byte  (respString))
	   	return
	}
	log.Println("=============== here rows: ", rows)
	log.Println( reflect.TypeOf( rows ))
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&rawConvCard.Distinct, &rawConvCard.Chat_room_name, & rawConvCard.User_id_1, 
			&rawConvCard.User_id_2,  &rawConvCard.Name1, &rawConvCard.Name2,
			&rawConvCard.Msg, &rawConvCard.Not_read_count, &rawConvCard.Date_sent, 
			&rawConvCard.User_last_message_id)
		err != nil {

			log.Println("Error Scan Messages list: ===> ",err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
				respString, _ := json.Marshal(
					Response{ "Message": "Error Scan Data", "Status": 400, "error": err } )
			w.Write([]byte  (respString))	 
			return

		} else {
			var cardConvObj models.ConversationCard
			log.Println(" ========= here raw conversation card ===========")
			if rawConvCard.User_id_1 == user_id{
				cardConvObj.Name = rawConvCard.Name2
				cardConvObj.Id = rawConvCard.User_id_2
			} else if rawConvCard.User_id_2 == user_id{
				cardConvObj.Name = rawConvCard.Name1
				cardConvObj.Id = rawConvCard.User_id_1
			}
			cardConvObj.LastMsg = rawConvCard.Msg
			cardConvObj.UnreadCount = rawConvCard.Not_read_count
			cardConvObj.ChatRoomName = rawConvCard.Chat_room_name
			arr_cardConvObj = append(arr_cardConvObj, cardConvObj)
		}
	}
	if len(arr_cardConvObj) == 0 {
		w.WriteHeader(http.StatusNotFound)
			var respMsg ResponseConvCard
			respMsg.Message = "No Conversation Yet"
			respMsg.Status = 404
			respMsg.Data = arr_cardConvObj
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respMsg)
		return 
	}

	var respMsg ResponseConvCard
	respMsg.Message = "success get list conversation card"
	respMsg.Status = 200
	respMsg.Data = arr_cardConvObj

	log.Println(reflect.TypeOf(http.StatusOK))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respMsg)
	return
	
}

func getQueryCardConv() string {

	var query string =`
	select 
		DISTINCT ON (cr.chat_room_name) cr.chat_room_name as distinct,
		cr.*, uc1.email as name1, uc2.email as name2,
		msg.message,
		msg.not_read_count,
		msg.date_sent,
		msg.user_id as user_last_message_id
	from
		chatroom cr
	Left join
		userchat uc1
		on cr.user_id_1 = uc1.user_id
	Left join
		userchat uc2
		on cr.user_id_2 = uc2.user_id
	Left join lateral
		(
			select message, chat_room_name, user_id, sent_time as date_sent, read_time, sent_time, count(*) over(),
			SUM(CASE WHEN 
				read_time is null and user_id !=$1
				then 1 else 0 end )OVER() AS not_read_count
			from message m
			where m.chat_room_name = cr.chat_room_name
			order by sent_time desc
		) msg on msg.chat_room_name = cr.chat_room_name

 	where user_id_1 = $1 OR user_id_2 = $1
	`

	return query
	// really this one trying to pull anything out of nowhere
}