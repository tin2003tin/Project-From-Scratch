package main

import (
	"database/db"
	"database/db/table"
	"fmt"
)

func main() {
	database, err := db.GetDataBase("F_market")
	if err != nil {
		fmt.Println(err)
		return
	}
	marketTable := database.Registry.Tables["market"]
	foodTable := database.Registry.Tables["food"]
	drinkTable := database.Registry.Tables["drink"]
	fmt.Println("Before\n\n")
	marketTable.PrintAsTable()
	foodTable.PrintAsTable()
	drinkTable.PrintAsTable()

	queriedRow, err := foodTable.QueryRows([]table.Condition{
		{ColumnName: "price", Operator: ">=", Value: float64(10)},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	table.PrintAsTable(foodTable.Metadata.Columns, queriedRow)
	mJf_Table, err := marketTable.Join(foodTable, table.InnerJoin, []table.On{{Self: "id", Operator: "=", Another: "market_id"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	mJf_Table.PrintAsTable()
	queriedRow, err = mJf_Table.QueryRows([]table.Condition{
		{ColumnName: "name", Operator: "=", Value: "Market A"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	table.PrintAsTable(mJf_Table.Metadata.Columns, queriedRow)
	mJd_Table, err := marketTable.Join(drinkTable, table.InnerJoin, []table.On{{Self: "id", Operator: "=", Another: "market_id"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	mJd_Table.PrintAsTable()
	queriedRow, err = mJf_Table.QueryRows([]table.Condition{
		{ColumnName: "name", Operator: "=", Value: "Market A"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	table.PrintAsTable(mJd_Table.Metadata.Columns, queriedRow)
}
