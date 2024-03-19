package main

import (
	"database/db"
	"database/db/table"
	"fmt"
)

func main() {
	database, err := db.GetDataBase("A_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Database's name:", database.Name)
	fmt.Println("-----------------")
	userTable := database.Tables[0]
	userTable.PrintAsTable()
	conditions := []table.Condition{
		table.Condition{ColumnName: "age", Operator: "<", Value: 50},
		table.Condition{ColumnName: "gender", Operator: "!=", Value: "female"},
		// query.Condition{ColumnName: "fullName", Operator: "=", Value: "Isabella Clark"},
	}
	fmt.Println("-----------------")
	fmt.Println("Query conditions:", conditions)
	queriedData, err := userTable.QueryRows(conditions)
	// indexedData,err := query.QueryRowByIndex(userTable.IndexTable["email"],"alexander@example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	// for _, row := range queriedData {
	// 	fmt.Println("-", row.Data)
	// }
	table.PrintAsTable(userTable.Metadata.Columns, queriedData)
	// fmt.Println(indexedData.Data)
}
