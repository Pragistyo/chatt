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

func Login(w http.ResponseWriter,r *http.Request){
	
	conn := db.Connect();
	defer conn.Close()

	err := r.ParseMultipartForm(64) // max memory 64kb
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte  ("error parse input"))
	}

	var email string = r.FormValue("email")
	var u  models.User // id, email, name
	row := conn.QueryRow( context.Background(), "SELECT user_id, email, name FROM UserChat WHERE email=$1",   email)

	err = row.Scan(&u.Id, &u.Email, & u.Name)
	if err!=nil {
		log.Println(" ==== error login: ", err)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
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