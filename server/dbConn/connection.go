package dbconn

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
)

type Access struct {
	Username string
	Password string
	Database string
}

type DatabaseConnetion struct {
	conn net.Conn
}

func Connect(port string, username string, password string, database string) (*DatabaseConnetion, error) {
	serverAddr := ":" + port
	access := Access{
		Username: username,
		Password: password,
		Database: database,
	}
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return nil, err
	}
	if err := SendAccess(conn, access); err != nil {
		fmt.Println("Error sending access information:", err)
		return nil, err
	}
	return &DatabaseConnetion{conn: conn}, nil
}

func SendAccess(conn net.Conn, access Access) error {
	accessJSON, err := json.Marshal(access)
	if err != nil {
		return err
	}

	jsonLenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(jsonLenBuf, uint32(len(accessJSON)))
	if _, err := conn.Write(jsonLenBuf); err != nil {
		return err
	}

	if _, err := conn.Write(accessJSON); err != nil {
		return err
	}

	return nil
}

func (dc *DatabaseConnetion) Execute(sqltext string) error {
	if !(strings.HasSuffix(sqltext, "$")) {
		return errors.New("the SQL statement should end with $")
	}
	writer := bufio.NewWriter(dc.conn)

	_, err := writer.WriteString(sqltext + "\n")
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	response := make([]byte, 1024)
	n, err := dc.conn.Read(response)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}
	fmt.Println("Response from server:", string(response[:n]))

	return nil
}

func (dc *DatabaseConnetion) Close() {
	dc.conn.Close()
}
