package internal

type Result struct {
	Code        string      `json:"code"`
	QueryResult QueryResult `json:"queryResult"`
}

type QueryResult struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message"`
	RequestId string      `json:"requestId"`
}
