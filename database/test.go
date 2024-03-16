package main

// func main() {
// 	// Create a TableRegistry to store tables
// 	registry := &table.TableRegistry{
// 		Tables: make(map[string]*table.Table),
// 	}

// 	// Create the employee table
// 	employeeTable := table.CreateTable("employee",
// 		[]table.Column{
// 			{Name: "id", DataType: "int", PrimaryKey: true},
// 			{Name: "name", DataType: "string"},
// 			{Name: "department", DataType: "string"},
// 		},
// 		[]string{"id"},
// 		nil,
// 		nil,
// 		"",
// 		"Employee Table",
// 		"default_tablespace",
// 		nil,
// 		"admin",
// 		time.Now(),
// 	)
// 	registry.Tables["employee"] = employeeTable

// 	// Create the salary table
// 	salaryTable := table.CreateTable("salary",
// 		[]table.Column{
// 			{Name: "id", DataType: "int", PrimaryKey: true},
// 			{Name: "employee_id", DataType: "int", ForeignKey: true},
// 			{Name: "amount", DataType: "int"},
// 			{Name: "date", DataType: "date"},
// 		},
// 		[]string{"id"},
// 		[]table.ForeignKey{
// 			{
// 				Name:       "salary_employee_fk",
// 				Columns:    []string{"employee_id"},
// 				RefTable:   "employee",
// 				RefColumns: []string{"id"},
// 			},
// 		},
// 		nil,
// 		"",
// 		"Salary Table",
// 		"default_tablespace",
// 		nil,
// 		"admin",
// 		time.Now(),
// 	)
// 	registry.Tables["salary"] = salaryTable

// 	// Add employees to the employee table
// 	employeeTable.AddRow(map[string]interface{}{"id": 1, "name": "John Doe", "department": "Engineering"}, time.Now(), time.Now())
// 	employeeTable.AddRow(map[string]interface{}{"id": 2, "name": "Jane Smith", "department": "Engineering"}, time.Now(), time.Now())
// 	employeeTable.AddRow(map[string]interface{}{"id": 3, "name": "Alice Johnson", "department": "Finance"}, time.Now(), time.Now())

// 	// Add salaries to the salary table
// 	salaryTable.AddRow(map[string]interface{}{"id": 1, "employee_id": 1, "amount": 5000, "date": time.Now()}, time.Now(), time.Now())
// 	salaryTable.AddRow(map[string]interface{}{"id": 2, "employee_id": 2, "amount": 6000, "date": time.Now()}, time.Now(), time.Now())
// 	salaryTable.AddRow(map[string]interface{}{"id": 3, "employee_id": 3, "amount": 7000, "date": time.Now()}, time.Now(), time.Now())

// 	// Create a query with two conditions
// 	queryConditions := table.CreateQueryConditions(
// 		// map[string]interface{}{"columnName": "amount", "value": 5000},
// 		// map[string]interface{}{"columnName": "name", "value": "John Doe"},
// 	)

// 	// Query the employee table with the conditions
// 	// matchedEmployees, err := employeeTable.QueryRows(queryConditions)
// 	// if err != nil {
// 	// 	fmt.Println("Error querying employees:", err)
// 	// 	return
// 	// }
// 	matchedEmployees, err := salaryTable.QueryRowsWithForeignKeys(queryConditions,registry)
// 	if err != nil {
// 		fmt.Println("Error querying employees:", err)
// 		return
// 	}
// 	fmt.Println("Matching employees:")
// 	for _, employee := range matchedEmployees {
// 		fmt.Println(employee.Data)
// 	}
// }

// func main() {
// 	// Example usage
// 	table := t.CreateTable("employees", nil, nil, nil, nil, "", "Employee Table", "", nil, "admin", time.Now())
// 	err := table.CreateHashIndex("name_index")
// 	if err != nil {
// 		fmt.Println("Error creating hash index:", err)
// 		return
// 	}

// 	// Insert data into the index
// 	err = table.InsertIntoIndex("name_index", []interface{}{"John Doe"}, t.Row{Data: map[string]interface{}{"id": 1}})
// 	if err != nil {
// 		fmt.Println("Error inserting data into index:", err)
// 		return
// 	}

// 	// Query rows by index
// 	rows, err := table.QueryRowsByIndex("name_index", []interface{}{"John Doe"})
// 	if err != nil {
// 		fmt.Println("Error querying rows by index:", err)
// 		return
// 	}

// 	fmt.Println("Rows found:", rows)
// }

type Employee struct {
	ID         int
	Name       string
	Department string
}
