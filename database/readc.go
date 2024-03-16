package main

import (
	"database/db"
	"fmt"
)

func main() {
	database,err := db.GetDataBase("companyA");
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(database.Tables)
	for _,table := range database.Tables {
		fmt.Println("=====")
		fmt.Println(table.Metadata.Name)
		fmt.Println(table.Metadata.Rows)
		fmt.Println(table.IndexTable)
		fmt.Println(table.IndexTable["email"].Rows)
	}

	// employeesTable := database.Registry.Tables["employee"];
	// sampleData := [][]interface{}{
	// 	{1, "John Doe", "john.doe@example.com", "Engineering", 70000},
	// 	{2, "Jane Smith", "jane.smith@example.com", "Marketing", 60000},
	// 	{3, "Michael Johnson", "michael.johnson@example.com", "Finance", 80000},
	// 	{4, "Emily Williams", "emily.williams@example.com", "HR", 55000},
	// 	{5, "Daniel Brown", "daniel.brown@example.com", "Engineering", 75000},
	// 	{6, "Olivia Jones", "olivia.jones@example.com", "Marketing", 65000},
	// 	{7, "Matthew Davis", "matthew.davis@example.com", "Finance", 82000},
	// 	{8, "Sophia Miller", "sophia.miller@example.com", "HR", 56000},
	// 	{9, "William Wilson", "william.wilson@example.com", "Engineering", 76000},
	// 	{10, "Ava Taylor", "ava.taylor@example.com", "Marketing", 68000},
	// }

	// // Add rows to the table
	// columnNames := []string{"id", "name", "email", "department", "salary"}
	// for _, data := range sampleData {
	// 	columnValues := make(map[string]interface{})
	// 	for i := 0; i < len(columnNames); i++ {
	// 		columnValues[columnNames[i]] = data[i]
	// 	}

	// 	// Add the row to the table
	// 	err := employeesTable.AddRow(columnValues)
	// 	if err != nil {
	// 		fmt.Println("Error adding row:", err)
	// 		return
	// 	}
	// }
	// fmt.Println("Rows added successfully.")
	// employeesTable.SerializeRows();
	// err = database.Registry.Tables["employee"].AddColumn("department","string",255,0,0,false,false,false,false,"","","")
	// if (err != nil) {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = database.Registry.Tables["employee"].AddColumn("salary","int",255,0,0,false,false,false,false,"","","")
	// if (err != nil) {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(database.Tables[0].IndexTable["email"])
	// database.Tables[0].AddRow()
	// if err := database.Tables[0].AddIdColumn(); err !=nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = database.Registry.Tables["employee"].AddColumn("name","string",255,0,0,false,false,false,false,"","","")
	// if (err != nil) {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(database.Registry.Tables["employee"].Metadata.Columns)
	// fmt.Println(database.Tables[0].Metadata.Name)
	// fmt.Println(database.Registry.Tables["employee"].Metadata.Name)
	// _, err = database.CreateTable("employee");
	// _, err = database.CreateTable("salary");
	// if (err != nil) {
	// 	fmt.Println(err)
	// 	return 
	// }
}