package models



type ResponseListConversation struct {
	User  				User					`json:"id"` 
	ListConversation 	[]ConversationCard		`json:"list_conversation"` 
	ChatRoomId			int32					`json:"chat_room_id"`
	ChatRoomName		string					`json:"chat_room_name"`
}