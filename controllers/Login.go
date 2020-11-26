package controllers


import (
	"log"
	"net/http"
	"context"
	"encoding/json"
	db "github.com/Pragistyo/chatt/db"
	models "github.com/Pragistyo/chatt/models"
)

type Response map[string]interface{}
type ResponseSingleUser struct {
	Message		string					`json:"message"` 
	Status		int32					`json:"status"` 
	Data		models.User				`json:"Users"` 
}


func Login(w http.ResponseWriter,r *http.Request){
	
	conn := db.Connect();
	defer conn.Close()

	err := r.ParseMultipartForm(64) // max memory 64kb
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte  ("error parse input"))
		return
	}

	var email string = r.FormValue("email")
	var u  models.User // id, email, name
	var queryGetUser = "SELECT user_id, email, name FROM UserChat WHERE email=$1"

	row := conn.QueryRow( context.Background(), queryGetUser,   email)

	err = row.Scan(&u.Id, &u.Email, & u.Name)

	if err!=nil {
		log.Println(" ==== error login: ", err)
		w.WriteHeader(http.StatusNotFound)
		respString, _ := json.Marshal( Response{ "Message": "user not found ", "Status": 404 } )
		w.Write([]byte  (respString))
		return
	}

	var resp ResponseSingleUser
	resp.Message = "success"
	resp.Status = 200
	resp.Data = u

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func CreateUser(w http.ResponseWriter,r *http.Request){
	conn:= db.Connect()
	defer conn.Close()

	
	err := r.ParseMultipartForm(4096) // max memory 4mb
	if err != nil {
		panic(err)
	}

	var name string = r.FormValue("name")
	var email string = r.FormValue("email")
	var user_id int32

	//email validation

	var sqlStatement string = `
					INSERT INTO UserChat (name, email)
					VALUES ($1, $2)
					RETURNING user_id
		`
	err = conn.QueryRow(context.Background(), sqlStatement, name, email).Scan(&user_id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		respString, _ := json.Marshal( Response{ "Message": "Failed to create user", "Status": 400, "error": err } )
		w.Write([]byte  (respString))
		return 
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	respString, _ := json.Marshal(Response{"Message": "user created", "status": 401, "new_id":user_id })
	w.Write([]byte  (respString))
	return
}