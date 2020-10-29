package main

import (
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
	fmt.Print("server running at port 8001...")
	http.ListenAndServe(":8001", handler)

}
func InitAllController(r *mux.Router) {

	r.HandleFunc("/UploadFile", Read.ReadFileCSV).Methods("POST")

	//r.HandleFunc("/GetRecommend",Read.ReadFileCSV).Methods("GET")
}
