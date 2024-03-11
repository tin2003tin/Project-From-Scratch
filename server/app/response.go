package app

import (
	"encoding/json"
	"fmt"
	"net"
)

type Response struct {
	Message    string
	StatusCode int
}

const (
	SuccessStatusCode = 200
	ErrorStatusCode   = 404
)

func CreateResponseMessage(message string) ([]byte, error) {
    response := Response{
        Message:    message,
        StatusCode: SuccessStatusCode,
    }
    jsonBytes, err := json.Marshal(response)
    if (err != nil ) {
        return nil,err;
    }
    return jsonBytes, nil;
}

func HandleError(conn *net.Conn, err error) {
	if err != nil {
		errorMsg := struct{ Error string }{Error: err.Error()}
		jsonBytes, jsonErr := json.Marshal(errorMsg)
		if jsonErr != nil {
			fmt.Println("Error marshaling JSON:", jsonErr)
			return
		}
		_, writeErr := (*conn).Write(jsonBytes)
		if writeErr != nil {
			fmt.Println("Error writing error message:", writeErr)
		}
	}
}
