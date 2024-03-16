package main

import (
	"database/db"
	"database/db/query"
	"fmt"
)

func main() {
	database,err := db.GetDataBase("CompanyA")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println("-----------------")
	fmt.Println("Database's name:",database.Name)
	fmt.Println("Table's name:",database.Tables[0].Metadata.Name)
	fmt.Println("----------------")
	conditions := []query.Condition{
		query.Condition{ColumnName: "department", Operator: "=", Value: "Engineering"},
		query.Condition{ColumnName: "salary", Operator: ">=", Value: 75000},
	}
	fmt.Println("Query conditions:",conditions)
	queriedData,err := query.QueryRows(database.Tables[0],conditions);
	if (err != nil) {
		fmt.Println(err)
		return
	}
	for _, row := range queriedData {
		fmt.Println("-",row.Data)
	}
}