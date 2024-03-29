package main

import (
	"database/db"
	"database/db/table"
	"fmt"
)

func main() {
	database, err := db.GetDataBase("C_company")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-----------------")
	fmt.Println("Database's name:", database.Name)
	fmt.Println("----------------")
	for _, table := range database.Tables {
		table.PrintAsTable()
	}
	on_condition := []table.On{
		{Self: "id", Operator: "=", Another: "employee_id"},
	}

	joinedTable, err := database.Registry.Tables["employee"].Join(database.Registry.Tables["salary"], table.InnerJoin, on_condition)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Joined Data:")
	joinedTable.PrintAsTable()
	conditions := []table.Condition{
		table.Condition{ColumnName: "department", Operator: "=", Value: "Marketing"},
		table.Condition{ColumnName: "salary", Operator: ">=", Value: 70000},
	}
	fmt.Println("Query conditions:", conditions)
	queriedData, err := joinedTable.QueryRows(conditions)
	if err != nil {
		fmt.Println(err)
		return
	}
	table.PrintAsTable(joinedTable.Metadata.Columns, queriedData)
}
