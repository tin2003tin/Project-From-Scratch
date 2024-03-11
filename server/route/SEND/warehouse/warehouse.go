package warehouse

import (
	"encoding/base64"
	"fmt"
	"os"
	"server/app"
)
type File struct {
	Content []byte
	Size    float64
	Name    string
}

func SaveFileToWH(req app.Request) {
	dataMap, ok := req.Protocol.Body.Data.(map[string]interface{})
	if !ok {
		fmt.Println("Error: The Body is not file")
		return
	}
	contentBase64, ok := dataMap["Content"].(string)
	if !ok {
		fmt.Println("Error: Content not found or not a string in protocol.Body.Data")
		return
	}
	content, err := base64.StdEncoding.DecodeString(contentBase64)
	if err != nil {
		fmt.Println("Error decoding Content:", err)
		return
	}

	size, ok := dataMap["Size"].(float64)
	if !ok {
		fmt.Println("Error: Size not found or not a float64 in protocol.Body.Data")
		return
	}

	name, ok := dataMap["Name"].(string)
	if !ok {
		fmt.Println("Error: Name not found or not a string in protocol.Body.Data")
		return
	}

	file := &File{
		Content: content,
		Size:    size,
		Name:    name,
	}
	 f, err := os.Create("warehouse/"+file.Name)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer f.Close()

    // Write data to the file
    _, err =  f.Write(file.Content)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return	
    }

	fmt.Println("File saved successfully:", file.Name)
}