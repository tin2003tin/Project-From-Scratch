package runtime

import (
	"bufio"
	"database/db/structure"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)

type Request struct {
	Conn     *net.Conn
	Database *structure.Database
	Access   Access
}

type Access struct {
	Username string
	Password string
	Database string
}

func (app *App) handleClient(conn net.Conn) {
	access, err := ReadAccess(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(access)
	var found = false
	var db *structure.Database
	for i := range app.LoadedDatabase {
		if app.LoadedDatabase[i].Name == access.Database {
			if err := Authentication(&app.LoadedDatabase[i], access); err != nil {
				json := handleError(err)
				conn.Write(json)
				return
			}
			db = &app.LoadedDatabase[i]
			found = true
		}
	}
	if !found {
		json := handleError(fmt.Errorf(access.Database + " database is not found"))
		conn.Write(json)
		return
	}
	fmt.Println(db)

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('$')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
				return
			}
			fmt.Println("Error reading from connection:", err)
			return
		}
		start := time.Now()
		json, err := app.SqlCompliler.Prase(db, message)
		if err != nil {
			json = handleError(err)

		}
		conn.Write(json)
		loadDuration := time.Since(start)
		fmt.Println("Execution Total Time:", loadDuration)
	}
}

func ReadAccess(conn net.Conn) (*Access, error) {
	// Read JSON data length
	jsonLenBuf := make([]byte, 4)
	_, err := io.ReadFull(conn, jsonLenBuf)
	if err != nil {
		return nil, err
	}
	jsonLen := binary.BigEndian.Uint32(jsonLenBuf)

	// Read JSON data
	jsonBuf := make([]byte, jsonLen)
	_, err = io.ReadFull(conn, jsonBuf)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into Access struct
	var access Access
	err = json.Unmarshal(jsonBuf, &access)
	if err != nil {
		return nil, err
	}

	return &access, nil
}

func Authentication(database *structure.Database, access *Access) error {
	if database.Username != access.Username {
		return fmt.Errorf("username is invalid")
	}
	if database.Password != access.Password {
		return fmt.Errorf("password is invalid")
	}
	return nil
}
