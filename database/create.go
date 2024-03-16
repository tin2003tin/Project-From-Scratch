package main

import (
	"database/db"
	"fmt"
)

func main() {
	database,err := db.CreateDatabase("companyA")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(database)
}

