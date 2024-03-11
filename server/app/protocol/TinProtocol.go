package tinPro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)

type TinProtocol struct {
	Header tpHeader
	Body   tpBody
	Tail   tpTail
}

type tpHeader struct {
	Command    string    `json:"Command"`
	Path       string    `json:"Path"`
	Version    string    `json:"Version"`
	SecretKey  string    `json:"SecretKey"`
	MasterKey  string    `json:"MasterKey"`
	Date       time.Time `json:"Date"`
	BodyType   string    `json:"BodyType"`
	BodyLength int64     `json:"BodyLength"`
	TailLength int64     `json:"TailLength"`
}
type tpBody struct {
    Data interface{} `json:"Data"`
}
type tpTail struct {
	Message string `json:"Message"`
	Description string `json:"Description"`
}

func Middleware(header *tpHeader) error {
	if (header.Command == "" ) {
		return fmt.Errorf("the command is required")
	}
	if (header.SecretKey != "1234") {
		return fmt.Errorf("the SecretKey is incorrect")
	}
	if (header.Command == "TINY" && header.MasterKey == "master") {
		return fmt.Errorf("the TINY command need correctly masterKey")
	}
	return nil;
}

func ReadHeader(conn *net.Conn,	protocol *TinProtocol) error {
	return readAndUnmarshal(conn, &protocol.Header, 1024, true)
}

func ReadBody(conn *net.Conn, protocol *TinProtocol) error {
	return readAndUnmarshal(conn, &protocol.Body, int(protocol.Header.BodyLength), false)
}

func ReadTail(conn *net.Conn, protocol *TinProtocol) error {
	return readAndUnmarshal(conn, &protocol.Tail, int(protocol.Header.TailLength), false)
}

func readAndUnmarshal(conn *net.Conn, target interface{}, length int, head bool) error {
	var buffer bytes.Buffer
	index := length;
	for {
		tempBuffer := make([]byte, length)
		n, err := (*conn).Read(tempBuffer)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		buffer.Write(tempBuffer[:n])
		
		if head {
			index = bytes.Index(tempBuffer[:n], []byte("\n"))
			if index == -1 {
				index = n
			}
		}
		if index > buffer.Len() {
			index = buffer.Len()
		}
		break;
	}
	if err := json.Unmarshal(buffer.Bytes()[:index], target); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}