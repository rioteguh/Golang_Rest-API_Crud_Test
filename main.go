package main

import (
	"api/nama"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/api/nama/data", nama.FindAll).Methods("GET")
	router.HandleFunc("/api/nama/search/{keyword}", nama.Search).Methods("GET")
	router.HandleFunc("/api/nama/create", nama.Create).Methods("POST")
	router.HandleFunc("/api/nama/update", nama.Update).Methods("PUT")
	router.HandleFunc("/api/nama/delete/{id}", nama.Delete).Methods("DELETE")

	err := http.ListenAndServe(":1945",router)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Success. API ini berjalan di port 1945")
}