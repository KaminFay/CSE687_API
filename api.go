package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func sendTestFunctions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending Test Function")
}

func recieveTestFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Recieving Test Function")

}

func recieveResults(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Recieving Result")

}
