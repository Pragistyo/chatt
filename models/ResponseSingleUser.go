package models

type ResponseSingleUser struct {
	Message		string				`json:"message"` 
	Status		int32				`json:"status"` 
	Data		User				`json:"Users"` 
}