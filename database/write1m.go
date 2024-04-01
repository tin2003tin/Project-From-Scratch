package main

import (
	"database/db"
	"database/db/queryProcessor"
	"fmt"
)

//
//	Write 100k employee
//

func main() {
	// Create a new database named "Company"
	database, err := db.CreateDatabase("1m", "tin", "1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a table named "employee" in the database
	employeeTable, err := db.CreateTable(database, "employee")
	if err != nil {
		fmt.Println(err)
		return
	}
	employeeQuery := queryProcessor.NewQueryManager(employeeTable)

	// Add columns to the employee table
	employeeQuery.AddIdColumn()
	employeeQuery.AddColumn("name", "string", 0, false, false, false, nil)
	employeeQuery.AddColumn("department", "string", 0, false, false, false, nil)
	employeeQuery.AddColumn("age", "int", 0, false, false, false, nil)

	// Add rows to the employee table
	employeeQuery.AddRow([]interface{}{1, "John Doe", "IT", 30})
	employeeQuery.AddRow([]interface{}{2, "Jane Smith", "HR", 35})
	employeeQuery.AddRow([]interface{}{3, "Mike Johnson", "Finance", 40})
	employeeQuery.AddRow([]interface{}{4, "Emily Brown", "Marketing", 28})
	employeeQuery.AddRow([]interface{}{5, "Chris Wilson", "IT", 32})
	for i := 6; i <= 1000000; i++ {
		employeeQuery.AddRow([]interface{}{i, fmt.Sprintf("Employee%d", i), "Finance", 40})
	}
	employeeQuery.Commit()

	salaryTable, err := db.CreateTable(database, "salary")
	if err != nil {
		fmt.Println(err)
		return
	}
	salaryQuery := queryProcessor.NewQueryManager(salaryTable)
	salaryQuery.AddIdColumn()
	salaryQuery.AddColumn("employee_id", "int", 0, true, false, false, nil)
	salaryQuery.AddColumn("salary", "int", 0, false, false, false, nil)
	for i := 1; i <= 1000000; i++ {
		salaryQuery.AddRow([]interface{}{i, i, 10000 * (i%10 + 1)})
	}
	salaryQuery.Commit()
	fmt.Println("Write B_Company completed")
}
