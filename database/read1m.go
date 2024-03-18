package main

import (
	"database/db"
	"database/db/table"
	"fmt"
	"time"
)

func main() {
	fmt.Println("try to Read...:", "1m")
	startLoadDatabase := time.Now()
	database, err := db.GetDataBase("1m")
	if err != nil {
		fmt.Println(err)
		return
	}
	loadDuration := time.Since(startLoadDatabase)
	fmt.Println("Execution Time for Load the Database:", loadDuration)
	employeeTable, err := database.GetTable("employee")
	id := 500000
	// Define conditions for querying rows
	conditions := []table.Condition{
		{ColumnName: "id", Operator: "=", Value: id}, // Query for the last row
	}

	// Measure the execution time for QueryRows
	startQueryRows := time.Now()
	found, err := employeeTable.QueryRows(conditions)
	if err != nil {
		fmt.Println(err)
		return
	}
	queryRowsDuration := time.Since(startQueryRows)
	fmt.Println(found[0])
	// Measure the execution time for QueryRowByIndex
	startQueryRowByIndex := time.Now()
	ifound, err := employeeTable.QueryRowByIndex("id", id) // Query using the index for the last row
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*ifound)
	queryRowByIndexDuration := time.Since(startQueryRowByIndex)

	// Print the execution times
	fmt.Println("Execution Time for QueryRows:", queryRowsDuration)
	fmt.Println("Execution Time for QueryRowByIndex:", queryRowByIndexDuration)
}
