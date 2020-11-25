package models


import(

)

type User struct {	
	Id  		int32   		 	 `json:"id"` //,omitempty
	Email 	string    		 	 	`json:"Email"` //,omitempty
	Name	string    		  		`json:"Name"` //,omitempty
}
