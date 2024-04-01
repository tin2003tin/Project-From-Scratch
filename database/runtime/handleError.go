package runtime

import (
	"database/betaSql/lib"
	"fmt"
)

func handleError(perr error) []byte {
	errorMap := map[string]string{}
	errorMap["error"] = string(perr.Error())
	jsonData, err := lib.ConvertToJson(errorMap)
	if err != nil {
		fmt.Printf("Error converting error to JSON: %s", err)
		return nil
	}
	return jsonData
}
