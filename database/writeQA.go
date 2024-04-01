package main

import (
	"database/db"
	"database/db/queryProcessor"
	"fmt"
)

func main() {
	database, err := db.CreateDatabase("A_user","tin","1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	userTable, err := db.CreateTable(database, "user")
	if err != nil {
		fmt.Println(err)
		return
	}
	userQueryManager := queryProcessor.NewQueryManager(userTable)

	if err := userQueryManager.AddIdColumn(); err != nil {
		fmt.Println(err)
		return
	}
	if err := userQueryManager.AddColumn("fullName", "string", 255, false, false, false, nil); err != nil {
		fmt.Println(err)
		return
	}
	if err := userQueryManager.AddColumn("age", "int", 0, false, false, false, nil); err != nil {
		fmt.Println(err)
		return
	}
	if err := userQueryManager.AddColumn("gender", "string", 1, false, false, false, nil); err != nil {
		fmt.Println(err)
		return
	}
	if err := userQueryManager.AddColumn("email", "string", 255, false, true, false, nil); err != nil {
		fmt.Println(err)
		return
	}
	userQueryManager.AddRow([]interface{}{1, "John Doe", 16, "male", "john@example.com"})
	userQueryManager.AddRow([]interface{}{2, "Jane Smith", 18, "female", "jane@example.com"})
	userQueryManager.AddRow([]interface{}{3, "Michael Johnson", 20, "male", "michael@example.com"})
	userQueryManager.AddRow([]interface{}{4, "Emily Williams", 22, "female", "emily@example.com"})
	userQueryManager.AddRow([]interface{}{5, "Daniel Brown", 24, "male", "daniel@example.com"})
	userQueryManager.AddRow([]interface{}{6, "Olivia Jones", 26, "female", "olivia@example.com"})
	userQueryManager.AddRow([]interface{}{7, "Matthew Davis", 28, "male", "matthew@example.com"})
	userQueryManager.AddRow([]interface{}{8, "Sophia Miller", 30, "female", "sophia@example.com"})
	userQueryManager.AddRow([]interface{}{9, "William Wilson", 32, "male", "william@example.com"})
	userQueryManager.AddRow([]interface{}{10, "Ava Taylor", 34, "female", "ava@example.com"})
	userQueryManager.AddRow([]interface{}{11, "James Anderson", 36, "male", "james@example.com"})
	userQueryManager.AddRow([]interface{}{12, "Emma Martinez", 38, "female", "emma@example.com"})
	userQueryManager.AddRow([]interface{}{13, "Benjamin Hernandez", 40, "male", "benjamin@example.com"})
	userQueryManager.AddRow([]interface{}{14, "Mia Nelson", 42, "female", "mia1@example.com"})
	userQueryManager.AddRow([]interface{}{15, "Alexander Wright", 44, "male", "alexander@example.com"})
	userQueryManager.AddRow([]interface{}{16, "Charlotte Adams", 46, "female", "charlotte@example.com"})
	userQueryManager.AddRow([]interface{}{17, "David Lewis", 48, "male", "david@example.com"})
	userQueryManager.AddRow([]interface{}{18, "Isabella Clark", 50, "female", "isabella@example.com"})
	userQueryManager.AddRow([]interface{}{19, "Ethan Lee", 52, "male", "ethan@example.com"})
	userQueryManager.AddRow([]interface{}{20, "Amelia Baker", 54, "female", "amelia@example.com"})
	userQueryManager.AddRow([]interface{}{21, "Jacob Hall", 56, "male", "jacob@example.com"})
	userQueryManager.AddRow([]interface{}{22, "Ella Green", 58, "female", "ella@example.com"})
	userQueryManager.AddRow([]interface{}{23, "Michaela Hill", 60, "female", "michaela@example.com"})
	userQueryManager.AddRow([]interface{}{24, "Logan Hughes", 62, "male", "logan@example.com"})
	userQueryManager.AddRow([]interface{}{25, "Grace Carter", 64, "female", "grace@example.com"})
	userQueryManager.AddRow([]interface{}{26, "William Ward", 66, "male", "williamw@example.com"})
	userQueryManager.AddRow([]interface{}{27, "Sofia Turner", 68, "female", "sofia@example.com"})
	userQueryManager.AddRow([]interface{}{28, "Christopher Collins", 70, "male", "christopher@example.com"})
	userQueryManager.AddRow([]interface{}{29, "Madison Parker", 72, "female", "madison@example.com"})
	userQueryManager.AddRow([]interface{}{30, "Andrew Hall", 74, "male", "andrew@example.com"})
	userQueryManager.AddRow([]interface{}{31, "Scarlett Evans", 76, "female", "scarlett@example.com"})
	userQueryManager.AddRow([]interface{}{32, "Jameson Foster", 78, "male", "jameson@example.com"})
	userQueryManager.AddRow([]interface{}{33, "Lily Cooper", 80, "female", "lily@example.com"})
	userQueryManager.AddRow([]interface{}{34, "Sebastian Hughes", 82, "male", "sebastian@example.com"})
	userQueryManager.AddRow([]interface{}{35, "Natalie Griffin", 84, "female", "natalie@example.com"})
	userQueryManager.AddRow([]interface{}{36, "Josephine Mitchell", 86, "female", "josephine@example.com"})
	userQueryManager.AddRow([]interface{}{37, "Gabriel Johnson", 88, "male", "gabriel@example.com"})
	userQueryManager.AddRow([]interface{}{38, "Victoria Martin", 90, "female", "victoria@example.com"})
	userQueryManager.AddRow([]interface{}{39, "Nicholas Martinez", 92, "male", "nicholas@example.com"})
	userQueryManager.AddRow([]interface{}{40, "Hannah Butler", 94, "female", "hannah@example.com"})
	userQueryManager.AddRow([]interface{}{41, "David Richardson", 96, "male", "davidr@example.com"})
	userQueryManager.AddRow([]interface{}{42, "Zoe Stewart", 98, "female", "zoe@example.com"})
	userQueryManager.AddRow([]interface{}{43, "Christopher Harris", 100, "male", "christopherh@example.com"})
	userQueryManager.AddRow([]interface{}{44, "Mia Garcia", 102, "female", "mia@example.com"})
	userQueryManager.AddRow([]interface{}{45, "Samuel King", 104, "male", "samuel@example.com"})
	userQueryManager.AddRow([]interface{}{46, "Ella Young", 106, "female", "ella2@example.com"})
	userQueryManager.AddRow([]interface{}{47, "Matthew Perez", 108, "male", "matthewp@example.com"})
	userQueryManager.AddRow([]interface{}{48, "Avery Scott", 110, "female", "avery@example.com"})
	userQueryManager.AddRow([]interface{}{49, "Jackson Gonzalez", 112, "male", "jackson@example.com"})
	userQueryManager.AddRow([]interface{}{50, "Abigail Walker", 114, "female", "abigail@example.com"})
	userQueryManager.Commit()

	salaryTable, err := db.CreateTable(database, "salary")
	if err != nil {
		fmt.Println(err)
		return
	}
	salaryQueryManager := queryProcessor.NewQueryManager(salaryTable)

	if err := salaryQueryManager.AddIdColumn(); err != nil {
		fmt.Println(err)
		return
	}
	if err := salaryQueryManager.AddColumn("user_id", "int", 255, true, false, false, nil); err != nil {
		fmt.Println(err)
		return
	}
	if err := salaryQueryManager.AddColumn("deparment", "string", 0, false, false, false, nil); err != nil {
		fmt.Println(err)
		return
	}
	if err := salaryQueryManager.AddColumn("salary", "int", 1, false, false, false, nil); err != nil {
		fmt.Println(err)
		return
	}

	salaryQueryManager.AddRow([]interface{}{1, 1, "HR", 60000})
	salaryQueryManager.AddRow([]interface{}{2, 2, "Marketing", 40000})
	salaryQueryManager.AddRow([]interface{}{3, 3, "Engineer", 80000})
	salaryQueryManager.AddRow([]interface{}{4, 4, "Docter", 120000})
	salaryQueryManager.AddRow([]interface{}{5, 5, "CEO", 200000})
	salaryQueryManager.AddRow([]interface{}{6, 6, "Waiter", 30000})
	salaryQueryManager.AddRow([]interface{}{7, 7, "Reporter", 100000})
	salaryQueryManager.AddRow([]interface{}{8, 8, "Car Driver", 20000})
	salaryQueryManager.AddRow([]interface{}{9, 9, "Chief", 70000})
	salaryQueryManager.AddRow([]interface{}{10, 10, "Pilot", 130000})
	salaryQueryManager.AddRow([]interface{}{11, 11, "Musicist", 60000})
	salaryQueryManager.AddRow([]interface{}{12, 12, "HR", 40000})
	salaryQueryManager.AddRow([]interface{}{13, 13, "Marketing", 60000})
	salaryQueryManager.AddRow([]interface{}{14, 14, "Docter", 60000})
	salaryQueryManager.Commit()
	fmt.Println("Write A_user completed")
}
