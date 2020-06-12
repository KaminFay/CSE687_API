package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoing Hit: Homepage")
}

func requestHandler() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Print("Operating System we are running on: ")
	fmt.Println(runtime.GOOS)
	requestHandler()
}
