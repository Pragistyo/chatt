package models



type ResponseListConversation struct {
	User  User								`json:"id"` 
	ListConversation ListConversation		`json:"list_conversation"` 
}