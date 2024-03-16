package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.LoadDatabase("companyDatabase.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(database.Registry.Tables["employee"].IndexTable["gmail"].Rows["map[0:martin@gmail.com]"])
	fmt.Println(database.Tables[0].Metadata.Rows[4])
}