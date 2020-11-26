package models


import(
	"time"
	"database/sql"
)

type Message struct {	
	Id  			int32   		  	  `json:"id"` 
	Message 		string    		  	  `json:"message"` 
	SentTime		time.Time    	 	  `json:"senttime"` 
	ReadTime		sql.NullTime    	  `json:"readtime"` 
	ChatRoomName	string				  `json:"chat_room_name"`
	UserId			int32				  `json:"user_id"`
}
