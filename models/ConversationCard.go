package models


type ConversationCard struct {
	Name  		    string   		  `json:"id"`
	UnreadCount 	int32    		  `json:"unread_count"`
	LastMsg			string    	 	  `json:"last_msg"`
}