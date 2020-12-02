package controllers

import(
	"net/http"
	"encoding/json"
	"log"
)

func ResponseError(w http.ResponseWriter ,msg string, status int, statusHttp int, err error){
	log.Println("==== Error ====== ",err)

	w.WriteHeader(statusHttp)
	w.Header().Set("Content-Type", "application/json")
	respString, _ := json.Marshal(Response{ "Message":msg, "Status": status, "error":err.Error() } )
	
	w.Write([]byte  (respString))
}