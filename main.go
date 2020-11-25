package main

import(
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	controllers "github.com/pragistyo/chatt/controllers"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r:= mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/user/",controllers.GetAllUser).Methods(http.MethodGet)
}

	