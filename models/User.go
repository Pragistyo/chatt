package models


import(

)

type User struct {	
	Id  	int32   		 	 	`json:"id"` 
	Email 	string    		 	 	`json:"email"` 
	Name	string    		  		`json:"name"` 
}

type ResponseSingleUser struct {
	Message		string				`json:"message"` 
	Status		int32				`json:"status"` 
	Data		User				`json:"Data"` 
}

type ResponseMultiUser struct {
	Message		string				`json:"message"` 
	Status		int32				`json:"status"` 
	Data		[]User				`json:"Users"` 
}




