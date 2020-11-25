package models


import (

)

type ResponseListMessage struct {
	Messages 			[]Message				`json:"Messages"`
	ChatRoomId			int32					`json:"chat_room_id"`
	ChatRoomName		string					`json:"chat_room_name"`
}