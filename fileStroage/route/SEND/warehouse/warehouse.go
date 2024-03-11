package warehouse

import (
	"encoding/base64"
	"fmt"
	"os"
	"system/app"
)
type File struct {
	Content []byte
	Size    float64
	Name    string
}
func MappingFile(body interface{}) (*File, error) {
	bodyMap, ok := body.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("body is not a map[string]interface{}")
	}

	contentBase64, ok := bodyMap["Content"].(string)
	if !ok {
		return nil, fmt.Errorf("Content not found or not a string in body")
	}

	content, err := base64.StdEncoding.DecodeString(contentBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding Content: %v", err)
	}

	sizeFloat, ok := bodyMap["Size"].(float64)
	if !ok {
		return nil, fmt.Errorf("Size not found or not a float64 in body")
	}
	size := float64(sizeFloat)

	name, ok := bodyMap["Name"].(string)
	if !ok {
		return nil, fmt.Errorf("Name not found or not a string in body")
	}

	// Creating a File struct
	file := &File{
		Content: content,
		Size:    size,
		Name:    name,
	}

	return file, nil
}

func SaveFileToWH(req app.Request) {
	file,err := MappingFile(req.Protocol.Body.Data);
	if (err !=nil ) {
		return
	}

	 f, err := os.Create("warehouse/"+file.Name)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer f.Close()

    _, err =  f.Write((*file).Content)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return	
    }
	fmt.Println("File saved successfully:", file.Name)
}