package main

import (
	"fmt"
	dbconn "server/dbConn"
	"time"
)

type Access struct {
	Username string
	Password string
	Database string
}

func main() {
	start := time.Now()
	databaseConn, err := dbconn.Connect("5432", "tin", "1234", "1m")
	if err != nil {
		fmt.Println(err)
		return
	}
	message := `Select * From employee 
				Where id = 50000 
				$`
	err = databaseConn.Execute(message)
	if err != nil {
		fmt.Println(err)
	}
	message2 := `Select * From salary  
				Where employee_id = 50000 
				$`
	err = databaseConn.Execute(message2)
	if err != nil {
		fmt.Println(err)
	}
	databaseConn.Close()
	duration := time.Since(start)
	fmt.Println("Access database time: ", duration)

}
