package main

import (
	"database/db"
	"fmt"
)

//
//	Write employee
//	Write salary
//

func main() {
	database,err := db.CreateDatabase("C_company")
	if (err != nil ) {
		fmt.Print(err)
		return ;
	}
	employeeTable,err := database.CreateTable("employee")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	employeeTable.AddIdColumn()
	employeeTable.AddColumn("name", "string", 0, 0, 0, false, false, false, false, nil, "", "")
	employeeTable.AddColumn("email", "string", 0, 0, 0, false, false, true, false, nil, "", "")
	employeeTable.AddColumn("department", "string", 0, 0, 0, false, false, false, false, nil, "", "")
	columnNames := []string{"id", "name", "email", "department"}

	sampleData := [][]interface{}{
		{201, "John Doe", "john.doe@example.com", "Engineering"},
		{202, "Jane Smith", "jane.smith@example.com", "Marketing"},
		{203, "Michael Johnson", "michael.johnson@example.com", "Finance"},
		{204, "Emily Williams", "emily.williams@example.com", "HR"},
		{205, "Daniel Brown", "daniel.brown@example.com", "Engineering"},
		{206, "Olivia Jones", "olivia.jones@example.com", "Marketing"},
		{207, "Matthew Davis", "matthew.davis@example.com", "Finance"},
		{208, "Sophia Miller", "sophia.miller@example.com", "HR"},
		{209, "William Wilson", "william.wilson@example.com", "Engineering"},
		{210, "Ava Taylor", "ava.taylor@example.com", "Marketing"},
	}

	// Add rows to the table
	for _, data := range sampleData {
		columnValues := make(map[string]interface{})
		for i := 0; i < len(columnNames); i++ {
			columnValues[columnNames[i]] = data[i]
		}

		// Add the row to the table
		err := employeeTable.AddRow(columnValues)
		if err != nil {
			fmt.Println("Error adding row:", err)
			return
		}
	}
	employeeTable.SerializeRows()

	salaryTable,err := database.CreateTable("salary")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	salaryTable.AddIdColumn()
	salaryTable.AddColumn("employee_id", "string", 0, 0, 0, true,false, false, false, nil, "", "")
	salaryTable.AddColumn("salary", "int", 0, 0, 0, false,false, false, false, nil, "", "")
	salaryColumnNames := []string{"id", "employee_id", "salary"}
	salarySampleData := [][]interface{}{
			{1, 201, 65000},
			{2, 202, 75000},
			{3, 203, 80000},
			{4, 204, 70000},
			{5, 205, 72000},
			{6, 206, 68000},
			{7, 207, 76000},
			{8, 208, 78000},
			{9, 209, 79000},
			{10, 210, 77000},
		}
	// Add rows to the table
	for _, data := range salarySampleData {
		columnValues := make(map[string]interface{})
		for i := 0; i < len(salaryColumnNames); i++ {
			columnValues[salaryColumnNames[i]] = data[i]
		}

		// Add the row to the table
		err := salaryTable.AddRow(columnValues)
		if err != nil {
			fmt.Println("Error adding row:", err)
			return
		}
	}
	salaryTable.SerializeRows()


	fmt.Println("Write C_Company completed")
}