package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.GetDataBase("H_market")
	if err != nil {
		fmt.Println(err)
		return
	}
	marketTable, err := database.GetTable("market")
	if err != nil {
		fmt.Println(err)
		return
	}
	marketTable.PrintAsTable()
	drinkTable, err := database.GetTable("drink")
	if err != nil {
		fmt.Println(err)
		return
	}
	drinkTable.PrintAsTable()
}
