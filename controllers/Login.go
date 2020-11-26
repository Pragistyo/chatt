package controllers


import (
	"log"
	"net/http"
	"context"
	"encoding/json"
	db "github.com/Pragistyo/chatt/db"
)

type Response map[string]interface{}

func Login(w http.ResponseWriter,r *http.Request){
	
}

func CreateUser(w http.ResponseWriter,r *http.Request){
	conn:= db.Connect()
	defer conn.Close()

	
	err := r.ParseMultipartForm(4096) // max memory 4mb
	if err != nil {
		panic(err)
	}

	var name string = r.FormValue("username")
	var email string = r.FormValue("password")
	var user_id int32

	//email validation

	var sqlStatement string = `
					INSERT INTO UserChat (name, email)
					VALUES ($1, $2)
					RETURNING user_id
		`
	err = conn.QueryRow(context.Background(), sqlStatement, email, name).Scan(&user_id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		respString, _ := json.Marshal( Response{ "Message": "Failed to create user", "Status": 400 } )
		w.Write([]byte  (respString))
		return 
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	respString, _ := json.Marshal(Response{"Message": "user created", "status": 401, "new_id":user_id }
	)
	w.Write([]byte  (respString))
	return
}