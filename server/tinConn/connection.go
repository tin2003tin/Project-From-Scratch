package tinConn

import (
	"fmt"
	"net"
	tinPro "server/tinConn/protocol"
)


type tinConnection struct {
	port string
	tinProtocol tinPro.TinReqProtocol
	errorHandler func(error)
	response Response
}

func CreateTinConnection(port string) *tinConnection {
	tc := tinConnection{port: port}
	return &tc
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

func (tc *tinConnection) Run() *Response {
	if !tc.isErrorhandlerExisted() {
		tc.errorHandler = handleError
	}
	tc.handleConnection(&tc.tinProtocol,tc.errorHandler);
	return &tc.response
}

func (tc *tinConnection) Error(callback func(error)) {
	tc.errorHandler = callback;
}

func (tc *tinConnection) isErrorhandlerExisted() bool{
	return tc.errorHandler != nil;
}

func (tc *tinConnection) handleConnection(protocol *tinPro.TinReqProtocol, callback func(error)) {
	//Check protocol valid 
	err := protocol.IsValid()
	if (err != nil) {
		callback(err)
		return
	}

	//Open connection
	domain := ":" + tc.port;
	conn, err := net.Dial("tcp", domain)
	if err != nil {	
		callback(err)
		return 
	}
	defer conn.Close()

	//Send Request to server
	err = sendRequest(conn, protocol);
	if (err != nil) {
		callback(err)
		return
	}

	// Handle Response
	err = tc.response.readReponse(conn);
	if (err != nil) {
		callback(err)
		return 
	}
}

func handleError(err error) {
	fmt.Println("(Client site) Request Error", err)
}