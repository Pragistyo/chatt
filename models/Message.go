package models


import(
	"time"
	"database/sql"
)

type Message struct {	
	Id  		int32   		  		  `json:"id"` //,omitempty
	Message 	string    		  		  `json:"message"` //,omitempty
	SentTime	time.Time    	 		  `json:"senttime"` //,omitempty
	ReadTime	sql.NullTime    		  `json:"readtime"` //,omitempty
	ChatRoomId	int32					  `json:"chat_room_id"`
	UserId		int32					  `json:"user_id"`
}
