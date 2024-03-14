package tinPro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func ReadHeader(conn net.Conn, protocol *TinResProtocol) error {
	bytes, err := readOnly(conn, 1024, true)
	if err != nil {
		return err
	}
	err = unmarshal(protocol.GetHeader(), bytes)
	if err != nil {
		return err
	}
	return nil
}

func ReadBody(conn net.Conn, protocol *TinResProtocol) error {
	bytes, err := readOnly(conn, int(protocol.GetHeader().BodyLength), false)
	if err != nil {
		return err
	}
	err = unmarshal(protocol.GetBody(), bytes)
	if err != nil {
		return err
	}
	return nil
}

func ReadTail(conn net.Conn, protocol *TinResProtocol) error {
	bytes, err := readOnly(conn, int(protocol.GetHeader().TailLength), false)
	if err != nil {
		return err
	}
	err = unmarshal(protocol.GetTail(), bytes)
	if err != nil {
		return err
	}
	return nil
}
func readOnly(conn net.Conn, length int, header bool) ([]byte, error) {
	var buffer bytes.Buffer
	index := length
	for {
		tempBuffer := make([]byte, length)
		n, err := conn.Read(tempBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		buffer.Write(tempBuffer[:n])
		if header {
			index = bytes.Index(tempBuffer[:n], []byte("\n"))
			if index == -1 {
				index = n
			}
		}
		if index > buffer.Len() {
			index = buffer.Len()
		}
		break
	}
	return buffer.Bytes()[:index], nil
}

func unmarshal(target interface{}, bytes []byte) error {
	if err := json.Unmarshal(bytes, target); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
