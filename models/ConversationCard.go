package models

import "database/sql"


type ConversationCard struct {
	Id          	int32        		`json:"id"`
	Name        	string       	 	`json:"name"`
	ChatRoomName 	string				`json:"chat_room_name"`
	UnreadCount 	sql.NullInt64 		`json:"unread_count"`
	LastMsg     	sql.NullString      `json:"last_msg"`
}