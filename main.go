package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoing Hit: Homepage")
}

func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Print("Operating System we are running on: ")
	fmt.Println(runtime.GOOS)
	requestHandler()
}
