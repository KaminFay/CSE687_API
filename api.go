package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

/*
*	Endpoint for uploading the testable functions from the GUI.
*	This is called from the C# side in the GUI to be added to the database.
 */
func sendTestFunction(w http.ResponseWriter, r *http.Request) {
	var tf functionToTest
	err := json.NewDecoder(r.Body).Decode(&tf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := insertTestFunction(tf.FunctionName, tf.DllName, tf.DllPath)

	fmt.Fprintf(w, "%d", id)
}

/*
*	Endpoint for returning the functions that need testing.
* 	Also calls to truncate the table once we have pulled the data.
*	This is called from the C++ side to be remove from the database.
 */
func recieveTestFunction(w http.ResponseWriter, r *http.Request) {
	listOfFunctions := getAllTestFunctions()
	json.NewEncoder(w).Encode(listOfFunctions)

	truncateTestFunctionTable()

}

/*
*	Endpoint for uploading the results of tests.
*	This is called from the C++ side to be added to the database.
 */
func sendResults(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoing Hit: sendResult")
	var fr functionResult
	err := json.NewDecoder(r.Body).Decode(&fr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertFunctionResult(fr)

}

/*
*	Endpoint for returning the results from the database.
*	This is called from the C# side to display in GUI.
 */
func recieveResults(w http.ResponseWriter, r *http.Request) {
	var idList resultID
	err := json.NewDecoder(r.Body).Decode(&idList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listOfResults := getResults(idList)
	if listOfResults == (functionResult{}) {
		fmt.Fprintf(w, "nil")
	} else {
		json.NewEncoder(w).Encode(listOfResults)
	}
}
