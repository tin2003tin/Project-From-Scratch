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
		return nil, fmt.Errorf("content not found or not a string in body")
	}

	content, err := base64.StdEncoding.DecodeString(contentBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding Content: %v", err)
	}

	sizeFloat, ok := bodyMap["Size"].(float64)
	if !ok {
		return nil, fmt.Errorf("size not found or not a float64 in body")
	}
	size := float64(sizeFloat)

	name, ok := bodyMap["Name"].(string)
	if !ok {
		return nil, fmt.Errorf("name not found or not a string in body")
	}

	// Creating a File struct
	file := &File{
		Content: content,
		Size:    size,
		Name:    name,
	}

	return file, nil
}

func SaveFileToWH(req app.Request, res app.Response) error {
	file, err := MappingFile(req.Protocol.GetBody().Data)
	if err != nil {
		return err
	}

	f, err := os.Create("warehouse/" + file.Name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer f.Close()

	// Write data to the file
	_, err = f.Write(file.Content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	text := "File saved successfully: " + file.Name
	err = res.SetMessage(text).Send()
	if err != nil {
		fmt.Println("Error sending response:", err)
		return err
	}

	fmt.Println("File saved successfully:", file.Name)
	return nil
}