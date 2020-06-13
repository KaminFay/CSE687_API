package main

type functionToTest struct {
	ID           int    `json:"ID"`
	FunctionName string `json:"FuncName"`
	DllName      string `json:"DllName"`
	DllPath      string `json:"DllPath"`
}

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

type resultID struct {
	ID int `json:"ID"`
}
