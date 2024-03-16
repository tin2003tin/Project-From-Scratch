package main

import (
	"database/db"
	"database/db/query"
	"fmt"
	"time"
)

func main() {
	startLoadDatabase := time.Now()
	database, err := db.GetDataBase("companyB")
	if (err != nil) {
		fmt.Println(err)
		return 
	}
	loadDuration := time.Since(startLoadDatabase)
	fmt.Println("Execution Time for Load the Database:",loadDuration)
	employeeTable,err := database.GetTable("employee");
	id := 80000
	// Define conditions for querying rows
	conditions := []query.Condition{
		{ColumnName: "id", Operator: "=", Value: id}, // Query for the last row
	}

	// Measure the execution time for QueryRows
	startQueryRows := time.Now()
	found, err := query.QueryRows(employeeTable, conditions)
	if err != nil {
		fmt.Println(err)
		return
	}
	queryRowsDuration := time.Since(startQueryRows)
	fmt.Println(found[0])
	// Measure the execution time for QueryRowByIndex
	startQueryRowByIndex := time.Now()
	ifound, err := query.QueryRowByIndex(employeeTable.IndexTable["id"], id) // Query using the index for the last row
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