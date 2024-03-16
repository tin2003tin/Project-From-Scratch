package main

import (
	"database/db"
	"fmt"
)

func main() {
	// Create a new database named "Company"
	database := db.CreateDatabase("company")
	// Create a table named "employee" in the database
	_, err := database.CreateTable("employee")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = database.CreateTable("salary")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Retrieve the "salary" table from the database
	salaryTable, err := database.GetTable("salary")
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
	err = employeeTable.AddIdColumn() // Add id column
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}	
	err = employeeTable.AddColumn("name", "string", 0, 0, 0, false, false, false, false, nil, "", "")
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}
	err = employeeTable.AddColumn("gmail", "string", 0, 0, 0, false, false, true, false, nil, "", "")
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}
	err = employeeTable.AddColumn("gender", "string", 0, 0, 0, false, false, false, false, nil, "", "")
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}

	// Add columns to the salary table
	err = salaryTable.AddIdColumn()
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}
	err = salaryTable.AddColumn("working_age", "int", 0, 0, 0, false, false, false, false, nil, "", "")
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}
	err = salaryTable.AddColumn("salary", "int", 0, 0, 0, false, false, false, false, nil, "", "")
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}

	// Add rows to the employee table
	err = employeeTable.AddRow(map[string]interface{}{
	"id":     1,
	"name":   "John Doe",
	"gmail":  "johndoe@example.com",
	"gender": "Male",
	})
	if err != nil {
		fmt.Println("Error adding row to employee table:", err)
		return
	}	
	err = employeeTable.AddRow(map[string]interface{}{
	"id":     2,
	"name":   "Tin Siriwid",
	"gmail":  "tin@gmail.com",
	"gender": "Male",
	})
	if err != nil {
		fmt.Println("Error adding row to employee table:", err)
		return
	}	

	// Add rows to the salary table
	err = salaryTable.AddRow(map[string]interface{}{
	"id":          1,
	"working_age": 10,
	"salary":      50000,
	})
	if err != nil {
		fmt.Println("Error adding row to salary table:", err)
		return
	}	
	err = salaryTable.AddRow(map[string]interface{}{
	"id":          2,
	"working_age": 15,
	"salary":      70000,
	})
	if err != nil {
		fmt.Println("Error adding row to salary table:", err)
		return
	}	

	// Print the database structure
	fmt.Println("Database Name:", database.Name)
	for _, table := range database.Tables {
		fmt.Println("Table Name:", table.Metadata.Name)
		fmt.Println("Table Columns:", table.Metadata.Columns)
		fmt.Println("Table Primary Keys:", table.Metadata.PrimaryKeys)
		fmt.Println("Table Foreign Keys:", table.Metadata.ForeignKeys)
		fmt.Println("Table Indexes:")
		for indexName, index := range table.IndexTable {
			fmt.Println("  - Index Name:", indexName)
			fmt.Println("    Index Columns:", index.Columns)
			fmt.Println("    Index Rows:", index.Rows)
			fmt.Println("    Index Unique:", index.Unique)
			fmt.Println("    Index Using:", index.Using)
			fmt.Println("    Index Comment:", index.Comment)
			fmt.Println("    Index Tablespace:", index.Tablespace)
			fmt.Println("    Index Include Columns:", index.Include)
			fmt.Println("    Index Predicate:", index.Predicate)
			fmt.Println("    Index Fill Factor:", index.FillFactor)
		}
		fmt.Println("Table Rows:")
		for _, row := range table.Metadata.Rows {
			fmt.Println("  - Row Data:", row.Data)
			fmt.Println("    Created At:", row.CreatedAt)
			fmt.Println("    Updated At:", row.UpdatedAt)
		}
		fmt.Println("----------------------")
	}
}