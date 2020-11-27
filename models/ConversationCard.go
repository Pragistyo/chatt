package models

import "database/sql"

import(
	"database/sql"
)

type ConversationCard struct {
	Id          string        		`json:"id"`
	Name        string       	 	`json:"name"`
	UnreadCount sql.NullInt64 		`json:"unread_count"`
	LastMsg     sql.NullString     `json:"last_msg"`
}