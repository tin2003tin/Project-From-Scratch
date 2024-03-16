package main

import (
	"database/db"
	"fmt"
)

func main() {
	database,err := db.CreateDatabase("CompanyA")
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
	employeeTable.AddColumn("salary", "int", 0, 0, 0, false, false, false, false, nil, "", "")
	columnNames := []string{"id", "name", "email", "department", "salary"}

	sampleData := [][]interface{}{
		{1, "John Doe", "john.doe@example.com", "Engineering", 70000},
		{2, "Jane Smith", "jane.smith@example.com", "Marketing", 60000},
		{3, "Michael Johnson", "michael.johnson@example.com", "Finance", 80000},
		{4, "Emily Williams", "emily.williams@example.com", "HR", 55000},
		{5, "Daniel Brown", "daniel.brown@example.com", "Engineering", 75000},
		{6, "Olivia Jones", "olivia.jones@example.com", "Marketing", 65000},
		{7, "Matthew Davis", "matthew.davis@example.com", "Finance", 82000},
		{8, "Sophia Miller", "sophia.miller@example.com", "HR", 56000},
		{9, "William Wilson", "william.wilson@example.com", "Engineering", 76000},
		{10, "Ava Taylor", "ava.taylor@example.com", "Marketing", 68000},
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
	fmt.Println("Write CompanyA completed")
}