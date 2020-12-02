package models

import( 
	"database/sql"
	"reflect"
	"log"
	"context"
	db "github.com/Pragistyo/chatt/db"
)


type ConversationCard struct {
	Id          	int32        		`json:"id"`
	Name        	string       	 	`json:"name"`
	ChatRoomName 	string				`json:"chat_room_name"`
	UnreadCount 	sql.NullInt64 		`json:"unread_count"`
	LastMsg     	sql.NullString      `json:"last_msg"`
}

type ConversationCardRaw struct { 
	Distinct  		    			string   		           
	Chat_room_name 		    		string   		         
	User_id_1 						int32    		  
	User_id_2						int32
	Name1							string
	Name2							string
	Msg								sql.NullString
	Not_read_count					sql.NullInt64
	date_sent 						sql.NullTime
	User_last_message_id			sql.NullInt32
}

func GetConvListDB(user_id int32 ){
	conn := db.Connect()
	defer conn.Close()


	var getQueryCardConv string = getQueryCardConv()
	rows, err := conn.Query( context.Background(), getQueryCardConv,   user_id)
	if err!= nil {
		log.Println("error in models conv card:", err)
		return
	}
	log.Println("==== type of rows =====:",reflect.TypeOf(rows))
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