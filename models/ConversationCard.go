package models


type ConversationCard struct {
	Id  		    string   		  `json:"id"`
	Name  		    string   		  `json:"name"`
	UnreadCount 	int32    		  `json:"unread_count"`
	LastMsg			string    	 	  `json:"last_msg"`
}