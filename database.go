package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "classPassword"
	dbname   = "cse687_db"
)

// DB is really dangerous, should not have global db connection but we're playing quick here
var DB *sql.DB

func establishDBConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected all right!")
	DB = db
}

func insertTestFunction(functionName string, dllName string, dllPath string) int {
	sqlStatment := `INSERT INTO test_functions (function_name, dll_name, dll_path)
					VALUES ($1, $2, $3) RETURNING id`

	id := 0
	err := DB.QueryRow(sqlStatment, functionName, dllName, dllPath).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

func insertFunctionResult(fr functionResult) {
	sqlStatment := `INSERT INTO finished_tests (id, function_name, dll_name, dll_path,
					bool_result, exception, start_time, end_time)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := DB.Exec(sqlStatment, fr.ID, fr.FunctionName, fr.DllName, fr.DllPath, fr.Result,
		fr.Exception, fr.StartTime, fr.EndTime)
	if err != nil {
		panic(err)
	}
}

func getAllTestFunctions() []functionToTest {

	var listOfFunctions []functionToTest

	sqlStatment := `SELECT id, function_name, dll_name, dll_path FROM test_functions`
	rows, err := DB.Query(sqlStatment)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var testFunc functionToTest

		err = rows.Scan(&testFunc.ID, &testFunc.FunctionName, &testFunc.DllName, &testFunc.DllPath)
		if err != nil {
			panic(err)
		}
		listOfFunctions = append(listOfFunctions, testFunc)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return listOfFunctions
}

func getResults(idList []resultID) []functionResult {

	var listOfResults []functionResult
	sqlStatment := `SELECT id, function_name, dll_name, dll_path, bool_result,
	exception, start_time, end_time FROM finished_tests WHERE id = $1`

	for _, result := range idList {
		var funcResult functionResult

		err := DB.QueryRow(sqlStatment, result.ID).Scan(&funcResult.ID, &funcResult.FunctionName,
			&funcResult.DllName, &funcResult.DllPath, &funcResult.Result, &funcResult.Exception,
			&funcResult.StartTime, &funcResult.EndTime)
		if err != nil {
			panic(err)
		}

		listOfResults = append(listOfResults, funcResult)
	}

	return listOfResults
}

func truncateTestFunctionTable() {
	sqlStatment := `TRUNCATE test_functions`

	_, err := DB.Exec(sqlStatment)
	if err != nil {
		panic(err)
	}
}

func closeDatabaseConnection() {
	DB.Close()
}
