package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.CreateDatabase("G_type")
	if err != nil {
		fmt.Println(err)
		return
	}
	mixedTable, _ := database.CreateTable("mixed")
	mixedTable.AddIdColumn()
	mixedTable.AddColumn("string", "string", 0, false, false, false, nil)
	mixedTable.AddColumn("varchar", "varchar", 10, false, false, false, nil)
	mixedTable.AddColumn("bool", "bool", 0, false, false, false, nil)
	mixedTable.AddColumn("int", "int", 0, false, false, false, nil)
	mixedTable.AddColumn("float", "float", 0, false, false, false, nil)
	mixedTable.AddColumn("byte", "byte", 10, false, false, false, nil)
	mixedTable.AddColumn("date", "date", 0, false, false, false, nil)
	mixedTable.AddColumn("intArray", "intArray", 0, false, false, false, nil)
	mixedTable.AddColumn("stringArray", "stringArray", 0, false, false, false, nil)
	mixedTable.AddColumn("floatArray", "floatArray", 0, false, false, false, nil)
	mixedTable.AddColumn("varcharArray", "varcharArray", 255, false, false, false, nil)
	mixedTable.AddColumn("boolArray", "boolArray", 0, false, false, false, nil)
	mixedTable.SerializeRows()
	fmt.Println("Write G_type completed")
}
