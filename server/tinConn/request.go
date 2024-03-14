package tinConn

import (
	tinPro "server/tinConn/protocol"
	"fmt"
	"net"
)

func sendRequest(conn net.Conn,protocol *tinPro.TinReqProtocol) error {
	//Marshal the Protocol
	JsonHeader, JsonBody, JsonTail, err := tinPro.MarshalProtocol(protocol);
	if (err != nil) {
		return err
	}

	fmt.Println("try to send to server...")

	//Write the connection
	_, err = conn.Write(append(JsonHeader, '\n'))
	if err != nil {
		return err
	}

	_, err = conn.Write(JsonBody)
	if err != nil {
		return err
	}

	_, err = conn.Write(JsonTail)
	if err != nil {
		return err
	}

	return nil;
}

