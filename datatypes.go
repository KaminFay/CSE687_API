package main

/*
* Datatype to be used for passing function data
 */
type functionToTest struct {
	ID           int    `json:"ID"`
	FunctionName string `json:"FuncName"`
	DllName      string `json:"DllName"`
	DllPath      string `json:"DllPath"`
}

/*
* Datatype to be used for passing result data
 */
type functionResult struct {
	ID           int    `json:"ID"`
	FunctionName string `json:"FuncName"`
	DllName      string `json:"DllName"`
	DllPath      string `json:"DllPath"`
	Result       bool   `json:"Result"`
	Exception    string `json:"Exception"`
	StartTime    string `json:"StartTime"`
	EndTime      string `json:"EndTime"`
}

/*
* This is used strictly to allow for us to read in JSON and marshal the data
* correctly.
 */
type resultID struct {
	ID int `json:"ID"`
}
