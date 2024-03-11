package warehouse

import (
	"encoding/base64"
	"fmt"
	"server/app"
	"server/tinConn"
	"server/tinConn/lib"
)
type File struct {
	Content []byte
	Size    float64
	Name    string
}

func MappingFile (body interface{}) *File {
	dataMap, ok := body.(map[string]interface{})
	if !ok {
		fmt.Println("Error: The Body is not file")
		return nil
	}
	contentBase64, ok := dataMap["Content"].(string)
	if !ok {
		fmt.Println("Error: Content not found or not a string in protocol.Body.Data")
		return nil
	}
	content, err := base64.StdEncoding.DecodeString(contentBase64)
	if err != nil {
		fmt.Println("Error decoding Content:", err)
		return nil
	}

	size, ok := dataMap["Size"].(float64)
	if !ok {
		fmt.Println("Error: Size not found or not a float64 in protocol.Body.Data")
		return nil
	}

	name, ok := dataMap["Name"].(string)
	if !ok {
		fmt.Println("Error: Name not found or not a string in protocol.Body.Data")
		return nil
	}

	file := &File{
		Content: content,
		Size:    size,
		Name:    name,
	}
	return file;
}

func SaveFileToWH(req app.Request) {
	file := (MappingFile(req.Protocol.Body.Data))
	tinConn.CreateTinConnection("8000").Access(lib.Command.SEND,"/warehouse","1234",lib.VERSION_1_0).Body(file).Run();
}