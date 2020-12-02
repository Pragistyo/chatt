package models

import( 
	"database/sql"
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
	Date_sent 						sql.NullTime
	User_last_message_id			sql.NullInt32
}
