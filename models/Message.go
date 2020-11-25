package models


import(
	"time"
	"database/sql"
)

type Message struct {	
	Id  		int32   		  		  `json:"id"` 
	Message 	string    		  		  `json:"message"` 
	SentTime	time.Time    	 		  `json:"senttime"` 
	ReadTime	sql.NullTime    		  `json:"readtime"` 
	ChatRoomId	int32					  `json:"chat_room_id"`
	UserId		int32					  `json:"user_id"`
}
