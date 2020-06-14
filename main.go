package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

/*
*	Unneccessary endpoint
 */
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoing Hit: Homepage")
}

/*
* Sets up the routs / endpoints that we can connect to
* Spins up the listener / server on port 8080
 */
func requestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/cse687/sendFunctions", sendTestFunction)
	router.HandleFunc("/cse687/recieveFunctions", recieveTestFunction).Methods("GET")
	router.HandleFunc("/cse687/sendResults", sendResults).Methods("POST")
	router.HandleFunc("/cse687/results", recieveResults).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}

func main() {

	// Setup for running this in the background
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s", s)
		serverShutdown()
		os.Exit(1)
	}()

	// Checking to make sure we are on linux
	fmt.Print("Operating System we are running on: ")
	fmt.Println(runtime.GOOS)

	establishDBConnection()
	requestHandler()
}

/*
*	If A signal is recieved to terminate *example a Ctrl-C*
*	or if the program crashes, this function is called to close
*	the DB connection.
 */
func serverShutdown() {
	log.Println("Need to clean up server.")
	DB.Close()
}
