package main

import (
	"database/db"
	"database/db/query"
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
	userTable := database.Tables[0]; 
	fmt.Println(userTable.IndexTable["email"])
	userTable.PrintAsTable();
	conditions := []query.Condition{
		query.Condition{ColumnName: "age", Operator: ">=", Value: 50},
		query.Condition{ColumnName: "gender", Operator: "=", Value: "female"},
		// query.Condition{ColumnName: "fullName", Operator: "=", Value: "Isabella Clark"},
	}
	fmt.Println("-----------------")
	fmt.Println("Query conditions:", conditions)
	queriedData, err := query.QueryRows(userTable, conditions)
	// indexedData,err := query.QueryRowByIndex(userTable.IndexTable["email"],"alexander@example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	// for _, row := range queriedData {
	// 	fmt.Println("-", row.Data)
	// }
	table.PrintAsTable(userTable.Metadata.Columns,queriedData);
	// fmt.Println(indexedData.Data)
}