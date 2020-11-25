package models


type ListConversation struct {
	Name  		    string   		  `json:"id"`
	UnreadCount 	int32    		  `json:"unread_count"`
	LastMsg			string    	 	  `json:"last_msg"`
}