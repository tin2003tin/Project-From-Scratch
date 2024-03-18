package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.GetDataBase("E_university")
	if err != nil {
		fmt.Println(err)
		return
	}
	majorTable := database.Registry.Tables["major"]
	studentTable := database.Registry.Tables["student"]
	majorTable.AddRow(map[string]interface{}{"id": 304, "name": "English", "head": "Dr. Martin"})
	majorTable.PrintAsTable()
	studentTable.AddRow(map[string]interface{}{"id": 505, "name": "Nut Not Jake", "email": "nut@example.com", "major_id": 301})
	err = studentTable.AddRow(map[string]interface{}{"id": 506, "name": "Gaming Gear", "email": "game@example.com", "major_id": 302})
	studentTable.AddRow(map[string]interface{}{"id": 507, "name": "Tony Stark", "email": "stark@example.com", "major_id": 303})
	err = studentTable.AddRow(map[string]interface{}{"id": 508, "name": "Tin Siriwid", "email": "tin@example.com", "major_id": 304})
	studentTable.PrintAsTable()
	if err != nil {
		fmt.Println(err)
	}

}
