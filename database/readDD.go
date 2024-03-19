package main

import (
	"database/db"
	"database/db/table"
	"fmt"
)

func main() {
	database, err := db.GetDataBase("D_product")
	if err != nil {
		fmt.Println(err)
		return
	}
	database.Tables[0].PrintAsTable()
	// conditions := []table.Condition{
	// 	{ColumnName: "name", Operator: "=", Value: "Banana"},
	// }

	// if err := database.Tables[0].DeleteRow(conditions); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	conditions := []table.Condition{
		{ColumnName: "id", Operator: "=", Value: 10},
	}

	sets := []table.Set{
		{ColumnName: "id", Value: 100},
		// {ColumnName: "name",Value: "test"},
	}
	err = database.Tables[0].UpdateRow(conditions, sets)
	if err != nil {
		fmt.Println(err)
		return
	}
	// if err := database.Tables[0].DeleteRow(conditions); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	database.Tables[0].PrintAsTable()
	fmt.Println((database.Tables[0].IndexTable.Rows))
	// database.Tables[0].SerializeRows()
}
