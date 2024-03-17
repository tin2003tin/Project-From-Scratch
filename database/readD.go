package main

import (
	"database/db"
	"fmt"
)

func main() {
	database,err :=db.GetDataBase("D_product")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	database.Tables[0].PrintAsTable()
	// conditions := []del.Condition{
	// 	{ColumnName: "name", Operator: "=", Value: "Banana"},
	// }
	fmt.Println((database.Tables[0].IndexTable["id"]).Rows)

	// if err := del.DeleteRow(database.Tables[0], conditions); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// conditions = []del.Condition{
	// 	{ColumnName: "name", Operator: "=", Value: "Pear"},
	// }

	// if err := del.DeleteRow(database.Tables[0], conditions); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// database.Tables[0].PrintAsTable()
	// fmt.Println((database.Tables[0].IndexTable["id"]).Rows)
	// database.Tables[0].SerializeRows()
}