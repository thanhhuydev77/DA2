package main

import (
	processing "Project2/Processing/ContentBaseFiltering"
	"Project2/Read"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	InitAllController(r)
	//allow all method CORS
	handler := cors.AllowAll().Handler(r)
	fmt.Print("server running at port 8001...\n")
	http.ListenAndServe(":8001", handler)

}
func InitAllController(r *mux.Router) {

	r.HandleFunc("/UploadFile", Read.ReadFileCSV).Methods("POST")
	r.HandleFunc("/GetRecommendContent",processing.GetContentRecommend).Methods("GET")

}
