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
	err = employeeTable.AddRow(map[string]interface{}{
	"id":     3,
	"name":   "Nut not Jake",
	"gmail":  "nutnotjake@gmail.com",
	"gender": "Male",
	})
	if err != nil {
		fmt.Println("Error adding row to employee table:", err)
		return
	}
	err = employeeTable.AddRow(map[string]interface{}{
	"id":     4,
	"name":   "Nia",
	"gmail":  "niania@gmail.com",
	"gender": "Female",
	})
	if err != nil {
		fmt.Println("Error adding row to employee table:", err)
		return
	}
	err = employeeTable.AddRow(map[string]interface{}{
	"id":     5,
	"name":   "Martin",
	"gmail":  "martin@gmail.com",
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

}