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

	message1 := `Delete From user
				Where id = 37
				$`
	message2 := `Insert Into user
				 Value ( 37,'Job',20,'male','job@gmail.com' )
				$`
	message3 := `Select fullName From user
				Where age < 50
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
