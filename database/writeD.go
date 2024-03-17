package main

import (
	"database/db"
	"fmt"
)

func main() {
	database,err := db.CreateDatabase("D_product")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	foodTable, err := database.CreateTable("fruit");
		if (err != nil) {
		fmt.Println(err)
		return
	}
	foodTable.AddIdColumn()
	foodTable.AddColumn("name", "string", 255, 0, 0, false, false, false, false, nil, "", "")
	foodTable.AddColumn("price", "int", 0, 0, 0, false, false, false, false, nil, "", "")
	foodTable.AddColumn("stock", "int", 1, 0, 0, false, false, false, false, nil, "", "")
	ColumnNames := []string{"id", "name", "price","stock"}

	examples := [][]interface{}{
		{1, "Apple", 100, 50},
		{2, "Banana", 80, 30},
		{3, "Orange", 120, 40},
		{4, "Grapes", 150, 20},
		{5, "Watermelon", 200, 15},
		{6, "Pineapple", 180, 25},
		{7, "Strawberry", 90, 35},
		{8, "Mango", 130, 45},
		{9, "Peach", 110, 55},
		{10, "Pear", 95, 60},
	}

	for _, data := range examples {
		columnValues := make(map[string]interface{})
		for i := 0; i < len(ColumnNames); i++ {
			columnValues[ColumnNames[i]] = data[i]
		}

		// Add the row to the table
		err := foodTable.AddRow(columnValues)
		if err != nil {
			fmt.Println("Error adding row:", err)
			return
		}
	}
	foodTable.SerializeRows()

	foodTable.PrintAsTable()

}