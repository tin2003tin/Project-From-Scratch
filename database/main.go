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
	employeeTable.AddIdColumn() // Add id column
	employeeTable.AddColumn("name", "string", 0, 0, 0, false, false, false, false, nil, "", "")
	employeeTable.AddColumn("gmail", "string", 0, 0, 0, false, false, true, false, nil, "", "")
	employeeTable.AddColumn("gender", "string", 0, 0, 0, false, false, false, false, nil, "", "")
	// Add columns to the salary table
	salaryTable.AddIdColumn()
	salaryTable.AddColumn("name", "string", 0, 0, 0, false, false, false, false, nil, "", "")
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
			fmt.Println("    Index Unique:", index.Unique)
			fmt.Println("    Index Using:", index.Using)
			fmt.Println("    Index Comment:", index.Comment)
			fmt.Println("    Index Tablespace:", index.Tablespace)
			fmt.Println("    Index Include Columns:", index.Include)
			fmt.Println("    Index Predicate:", index.Predicate)
			fmt.Println("    Index Fill Factor:", index.FillFactor)
		}
		fmt.Println("----------------------")
	}
}