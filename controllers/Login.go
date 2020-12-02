package controllers


import (
	"log"
	"net/http"
	"context"
	"encoding/json"
	"regexp"
	"fmt"
	db "chatt/db"
	models "chatt/models"
)

// type Response map[string]interface{}

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
		ResponseError(w, "user not found ", 404, http.StatusNotFound, err)
		return
	}

	var resp models.ResponseSingleUser
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
	if !isEmailValid(email) {
		log.Println("not a valid email")
		ResponseError(w, "Email not valid", 400, http.StatusBadRequest, fmt.Errorf( "email not valid") )
		return
	}

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
	respString, _ := json.Marshal(Response{"Message": "user created", "status": 201, "new_id":user_id })
	w.Write([]byte  (respString))
	return
}

func isEmailValid(email string) bool {
	log.Println("huahauhuahu")
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}