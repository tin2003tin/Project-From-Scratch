package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.CreateDatabase("E_university")
	if err != nil {
		fmt.Print(err)
		return
	}

	// Create a table for majors
	majorTable, err := database.CreateTable("major")
	if err != nil {
		fmt.Println(err)
		return
	}
	majorTable.AddIdColumn()
	majorTable.AddColumn("name", "string", 0, false, false, false, nil)
	majorTable.AddColumn("head", "string", 0, false, true, false, nil)
	majorColumnNames := []string{"id", "name", "head"}

	majorSampleData := [][]interface{}{
		{301, "Computer Science", "Dr. Smith"},
		{302, "Physics", "Dr. Johnson"},
		{303, "Mathematics", "Dr. Williams"},
	}

	// Add rows to the major table
	for _, data := range majorSampleData {
		columnValues := make(map[string]interface{})
		for i := 0; i < len(majorColumnNames); i++ {
			columnValues[majorColumnNames[i]] = data[i]
		}

		// Add the row to the major table
		err := majorTable.AddRow(columnValues)
		if err != nil {
			fmt.Println("Error adding row:", err)
			return
		}
	}
	majorTable.SerializeRows()

	// Create a table for students
	studentTable, err := database.CreateTable("student")
	if err != nil {
		fmt.Println(err)
		return
	}
	studentTable.AddIdColumn()
	studentTable.AddColumn("name", "string", 0, false, false, false, nil)
	studentTable.AddColumn("email", "string", 0,false, true, false, nil)
	studentTable.AddColumn("major_id", "int", 0,  false, false, false, nil)
	err = studentTable.CreateForeignKey("student_major_foreignKey", "major_id", database.Registry.Tables["major"], "id")
	if err != nil {
		fmt.Println(err)
		return
	}
	studentColumnNames := []string{"id", "name", "email", "major_id"}

	studentSampleData := [][]interface{}{
		{501, "John Doe", "john.doe@example.com", 301},
		{502, "Jane Smith", "jane.smith@example.com", 302},
		{503, "Michael Johnson", "michael.johnson@example.com", 303},
	}

	// Add rows to the student table
	for _, data := range studentSampleData {
		columnValues := make(map[string]interface{})
		for i := 0; i < len(studentColumnNames); i++ {
			columnValues[studentColumnNames[i]] = data[i]
		}

		// Add the row to the student table
		err := studentTable.AddRow(columnValues)
		if err != nil {
			fmt.Println("Error adding row:", err)
			return
		}
	}
	studentTable.SerializeRows()

	fmt.Println("Write E_University completed")
}
