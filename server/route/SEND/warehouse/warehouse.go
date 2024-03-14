package warehouse

import (
	"encoding/base64"
	"errors"
	"fmt"
	"server/app"
	"server/tinConn"
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

func SaveFileToWH(req app.Request,res app.Response) error {
	tc := tinConn.CreateTinConnection("8000");
	tc.Access(req.Protocol.GetHeader().Command,req.Protocol.GetHeader().Path,req.Protocol.GetHeader().SecretKey,req.Protocol.GetHeader().Version)
	tc.Body((MappingFile(req.Protocol.GetBody().Data)))
	response := tc.Run()
	if (response.GetResponse().StatusCode == 400) {
		app.ErrorToClient(*res.Conn, res.ReqProtocol, errors.New(response.GetResponse().Message))
		return nil
	}
	err := res.SetMessage(response.GetResponse().Message).SetBody(response.GetData().Data).Send()
	if (err != nil) {
		return err
	}
	return nil;
}