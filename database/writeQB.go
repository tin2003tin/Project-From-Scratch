package main

import (
	"database/db"
	"fmt"
)

//
//	Write 100k employee
//

func main() {
	// Create a new database named "Company"
	database, err := db.CreateDatabase("B_company")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a table named "employee" in the database
	_, err = database.CreateTable("employee")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Retrieve the "employee" table from the database
	employeeTable, err := database.GetTable("employee")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add columns to the employee table
	employeeTable.AddIdColumn()
	employeeTable.AddColumn("name", "string", 0, 0, 0, false, false, false, nil, "", "")
	employeeTable.AddColumn("department", "string", 0, 0, 0, false, false, false, nil, "", "")
	employeeTable.AddColumn("salary", "int", 0, 0, 0, false, false, false, nil, "", "")
	employeeTable.AddColumn("age", "int", 0, 0, 0, false, false, false, nil, "", "")

	// Add rows to the employee table
	employeeTable.AddRow(map[string]interface{}{"id": 1, "name": "John Doe", "department": "IT", "salary": 5000, "age": 30})
	employeeTable.AddRow(map[string]interface{}{"id": 2, "name": "Jane Smith", "department": "HR", "salary": 6000, "age": 35})
	employeeTable.AddRow(map[string]interface{}{"id": 3, "name": "Mike Johnson", "department": "Finance", "salary": 5500, "age": 40})
	employeeTable.AddRow(map[string]interface{}{"id": 4, "name": "Emily Brown", "department": "Marketing", "salary": 7000, "age": 28})
	employeeTable.AddRow(map[string]interface{}{"id": 5, "name": "Chris Wilson", "department": "IT", "salary": 5200, "age": 32})
	for i := 6; i <= 100000; i++ {
		employeeTable.AddRow(map[string]interface{}{"id": i, "name": fmt.Sprintf("Employee%d", i), "department": "Finance", "salary": 5500, "age": 40})
	}
	employeeTable.SerializeRows()
	fmt.Println("Write B_Company completed")
}
