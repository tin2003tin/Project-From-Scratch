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
	databaseConn, err := dbconn.Connect("5432", "tin", "1234", "A_user")
	if err != nil {
		fmt.Println(err)
		return
	}

	message1 := `Select * From user
				Where id > 34
				$`

	message2 := `Update user 
				Set fullName = 'nut', age = 100
				Where id > 34
				$`

	message3 := `Select * From user

				$`

	fmt.Println("====")
	err = databaseConn.Execute(message1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("====")
	err = databaseConn.Execute(message2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("====")
	err = databaseConn.Execute(message3)
	if err != nil {
		fmt.Println(err)
	}
	databaseConn.Close()
	duration := time.Since(start)
	fmt.Println("Access database time: ", duration)

}
