package tinConn

import (
	tinPro "client/tinConn/protocol"
	"encoding/json"
	"fmt"
	"net"
)


type tinConnection struct {
	tinProtocol tinPro.TinProtocol 

}

func (tc *tinConnection) Access(command, path, secretKey, version string) *tinConnection {
	tc.tinProtocol.SetAccess(command, path, secretKey, version)
	return tc;
}

func (tc *tinConnection) Body(data interface{}) *tinConnection {
	tc.tinProtocol.SetBody(interface{}(data));
	return tc
}

func (tc *tinConnection) Tail(message string, description string) *tinConnection {
	tc.tinProtocol.SetTail(message, description);
	return tc
}

func CreateTinConnection() *tinConnection {
	tc := tinConnection{}
	return &tc
}

func (tc *tinConnection) Run() error {
	err := handleConnection(&tc.tinProtocol);
	if (err != nil) {
		return err
	}
	return nil;	
}


func handleConnection(protocol *tinPro.TinProtocol) error {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)		
		return err
	}

	defer conn.Close()

	JsonTail, err := marshalJsonTail(protocol);
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	JsonBody, err := marshalJsonBody(protocol);
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	JsonHeader, err := marshalJsonHeader(protocol);
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("try to send to server...")
	_, err = conn.Write(append(JsonHeader, '\n'))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	_, err = conn.Write(JsonBody)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	
	_, err = conn.Write(JsonTail)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Handle Response
	buffer := make([]byte, 1024)
    _, err = conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading response:", err.Error())
        return nil;
    }
	 response := string(buffer)
    fmt.Println("Response from server: \n", response)
	return nil
}

func marshalJsonHeader(protocol *tinPro.TinProtocol) ([]byte, error) {
	JsonHeader, err := json.Marshal(protocol.Header)
	if err != nil {
		fmt.Println("Error:", err)
		return nil,err
	}
	return JsonHeader, nil;
}

func marshalJsonBody(protocol *tinPro.TinProtocol) ([]byte, error) {
	JsonBody, err := json.Marshal(protocol.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	if len(JsonBody) != 0 {

    protocol.Header.BodyLength = int64(len(JsonBody))
    protocol.Header.BodyType = "text"

	}
	return JsonBody, nil;
}

func marshalJsonTail(protocol *tinPro.TinProtocol) ([]byte, error) {
	JsonTail, err := json.Marshal(protocol.Tail)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	if len(JsonTail) != 0 {

    protocol.Header.TailLength = int64(len(JsonTail))
	}
	return JsonTail, nil;
}