package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "classPassword"
	dbname   = "cse687_db"
)

func establishDBConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected all right!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoing Hit: Homepage")
}

func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/cse687/sendFunctions", sendTestFunctions)
	router.HandleFunc("/cse687/recieveFunctions", recieveTestFunction)
	router.HandleFunc("/cse687/results", recieveResults).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Print("Operating System we are running on: ")
	fmt.Println(runtime.GOOS)
	establishDBConnection()
	requestHandler()
}
