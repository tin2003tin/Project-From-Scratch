package main

import (
	"database/db/queryProcessor"
	buffermanager "database/db/storageManager/bufferManager"
	"fmt"
)

func main() {
	database, err := buffermanager.LoadDatabaseMetadata("A_user")
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := buffermanager.LoadTableMetadata(database, "user")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = buffermanager.LoadIndex(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = buffermanager.LoadRawData(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = buffermanager.BuildIndex(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	s, err := buffermanager.LoadTableMetadata(database, "salary")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = buffermanager.LoadIndex(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = buffermanager.LoadRawData(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = buffermanager.BuildIndex(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	userTableQ := queryProcessor.NewQueryManager(t)
	salaryTalbleQ := queryProcessor.NewQueryManager(s)
	userTableQ.PrintAsTable()
	salaryTalbleQ.PrintAsTable()
	// userTableQ.UpdateRow([]structure.Condition{structure.Condition{ColumnName: "gender", Operator: "=", Value: "male"}}, []queryProcessor.Set{queryProcessor.Set{ColumnName: "fullName", Value: "Siriwid"}})
	// err = userTableQ.DeleteRow([]structure.Condition{structure.Condition{ColumnName: "id", Operator: "=", Value: 50}})
	// userTableQ.Commit()
	// userTableQ.PrintAsTable()
	// // _, err = userTableQ.WhereWithIndex(&[]structure.Condition{structure.Condition{ColumnName: "id", Operator: "=", Value: 10}})
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	return
	// // }
	// userTableQ.ResetCurrent()
	// userTableQ.PrintAsTable()
	ss, _ := userTableQ.Select(true, []string{"age", "fullName"})
	fmt.Println(ss)
	// userTableQ.Commit()
	// userTableQ.Join(s, structure.InnerJoin, structure.On{Self: "id", Operator: "=", Another: "user_id"})
	// userTableQ.Where(&[]structure.Condition{structure.Condition{ColumnName: "user_id", Operator: "=", Value: 10}}) // userTableQ.Where(&[]structure.Condition{structure.Condition{ColumnName: "gender", Operator: "=", Value: "male"}})
	// userTableQ.PrintAsTable()
	// userTableQ.PrintAsTable()
	// salaryTalbleQ.AddRow([]interface{}{15,15,"HR",50000})
	// userTableQ.Join(s, structure.InnerJoin, structure.On{Self: "id", Operator: "=", Another: "user_id"})
	// fmt.Println(userTableQ.CurrentColumns)
	// fmt.Println(len(userTableQ.CurrentRows))
	// userTableQ.PrintAsTable()
	// userTableQ.Where(&[]structure.Condition{structure.Condition{ColumnName: "salary",Operator: ">", Value: 50000}})
	// userTableQ.PrintAsTable()
	// userTableQ.PrintAsTable()
	// // err = userTableQ.AddRow([]interface{}{51, "Siriwid", 20, "male", "avery1@example.com", "Test", "Testestsetst", "Tet1231"})
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	return
	// // }
	// // userTableQ.Where(&[]structure.Condition{structure.Condition{ColumnName: "age", Operator: ">", Value: 40}})
	// // fmt.Println(len(userTableQ.Table.Rows))
	// // fmt.Println(len(userTableQ.CurrentRows))
	// // userTableQ.PrintAsTable()
	// // userTableQ.Commit()
	// fmt.Println(userTableQ.CurrentRows)
	// fmt.Println(userTableQ.Table.Rows)
	// userTableQ.PrintAsTable()
	// userTableQ.ResetCurrent()

	// userTableQ.Where(&[]structure.Condition{structure.Condition{ColumnName: "gender", Operator: "=", Value: "female"},
	// 										structure.Condition{ColumnName: "age",Operator: ">",Value: 50}})
	// fmt.Println(userTableQ.CurrentRows)
	// userTableQ.PrintAsTable()
}
