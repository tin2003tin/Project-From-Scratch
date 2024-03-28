package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.CreateDatabase("A_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	userTable, err := database.CreateTable("user")
	if err != nil {
		fmt.Println(err)
		return
	}
	userTable.AddIdColumn()
	userTable.AddColumn("fullName", "string", 255, false, false, false, nil)
	userTable.AddColumn("age", "int", 0, false, false, false, nil)
	userTable.AddColumn("gender", "string", 1, false, false, false, nil)
	userTable.AddColumn("email", "string", 255, false, true, false, nil)

	userTable.AddRow(map[string]interface{}{"id": 1, "fullName": "John Doe", "age": 16, "gender": "male", "email": "john@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 2, "fullName": "Jane Smith", "age": 18, "gender": "female", "email": "jane@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 3, "fullName": "Michael Johnson", "age": 20, "gender": "male", "email": "michael@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 4, "fullName": "Emily Williams", "age": 22, "gender": "female", "email": "emily@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 5, "fullName": "Daniel Brown", "age": 24, "gender": "male", "email": "daniel@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 6, "fullName": "Olivia Jones", "age": 26, "gender": "female", "email": "olivia@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 7, "fullName": "Matthew Davis", "age": 28, "gender": "male", "email": "matthew@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 8, "fullName": "Sophia Miller", "age": 30, "gender": "female", "email": "sophia@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 9, "fullName": "William Wilson", "age": 32, "gender": "male", "email": "william@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 10, "fullName": "Ava Taylor", "age": 34, "gender": "female", "email": "ava@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 11, "fullName": "James Anderson", "age": 36, "gender": "male", "email": "james@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 12, "fullName": "Emma Martinez", "age": 38, "gender": "female", "email": "emma@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 13, "fullName": "Benjamin Hernandez", "age": 40, "gender": "male", "email": "benjamin@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 14, "fullName": "Mia Nelson", "age": 42, "gender": "female", "email": "mia@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 15, "fullName": "Alexander Wright", "age": 44, "gender": "male", "email": "alexander@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 16, "fullName": "Charlotte Adams", "age": 46, "gender": "female", "email": "charlotte@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 17, "fullName": "David Lewis", "age": 48, "gender": "male", "email": "david@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 18, "fullName": "Isabella Clark", "age": 50, "gender": "female", "email": "isabella@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 19, "fullName": "Ethan Lee", "age": 52, "gender": "male", "email": "ethan@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 20, "fullName": "Amelia Baker", "age": 54, "gender": "female", "email": "amelia@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 21, "fullName": "Jacob Hall", "age": 56, "gender": "male", "email": "jacob@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 22, "fullName": "Ella Green", "age": 58, "gender": "female", "email": "ella@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 23, "fullName": "Michaela Hill", "age": 60, "gender": "female", "email": "michaela@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 24, "fullName": "Logan Hughes", "age": 62, "gender": "male", "email": "logan@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 25, "fullName": "Grace Carter", "age": 64, "gender": "female", "email": "grace@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 26, "fullName": "William Ward", "age": 66, "gender": "male", "email": "williamw@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 27, "fullName": "Sofia Turner", "age": 68, "gender": "female", "email": "sofia@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 28, "fullName": "Christopher Collins", "age": 70, "gender": "male", "email": "christopher@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 29, "fullName": "Madison Parker", "age": 72, "gender": "female", "email": "madison@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 30, "fullName": "Andrew Hall", "age": 74, "gender": "male", "email": "andrew@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 31, "fullName": "Scarlett Evans", "age": 76, "gender": "female", "email": "scarlett@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 32, "fullName": "Jameson Foster", "age": 78, "gender": "male", "email": "jameson@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 33, "fullName": "Lily Cooper", "age": 80, "gender": "female", "email": "lily@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 34, "fullName": "Sebastian Hughes", "age": 82, "gender": "male", "email": "sebastian@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 35, "fullName": "Natalie Griffin", "age": 84, "gender": "female", "email": "natalie@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 36, "fullName": "Josephine Mitchell", "age": 86, "gender": "female", "email": "josephine@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 37, "fullName": "Gabriel Johnson", "age": 88, "gender": "male", "email": "gabriel@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 38, "fullName": "Victoria Martin", "age": 90, "gender": "female", "email": "victoria@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 39, "fullName": "Nicholas Martinez", "age": 92, "gender": "male", "email": "nicholas@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 40, "fullName": "Hannah Butler", "age": 94, "gender": "female", "email": "hannah@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 41, "fullName": "David Richardson", "age": 96, "gender": "male", "email": "davidr@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 42, "fullName": "Zoe Stewart", "age": 98, "gender": "female", "email": "zoe@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 43, "fullName": "Christopher Harris", "age": 100, "gender": "male", "email": "christopherh@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 44, "fullName": "Mia Garcia", "age": 102, "gender": "female", "email": "mia@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 45, "fullName": "Samuel King", "age": 104, "gender": "male", "email": "samuel@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 46, "fullName": "Ella Young", "age": 106, "gender": "female", "email": "ella@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 47, "fullName": "Matthew Perez", "age": 108, "gender": "male", "email": "matthewp@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 48, "fullName": "Avery Scott", "age": 110, "gender": "female", "email": "avery@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 49, "fullName": "Jackson Gonzalez", "age": 112, "gender": "male", "email": "jackson@example.com"})
	userTable.AddRow(map[string]interface{}{"id": 50, "fullName": "Abigail Walker", "age": 114, "gender": "female", "email": "abigail@example.com"})

	userTable.SerializeRows()
	fmt.Println("Write A_school completed")
}
