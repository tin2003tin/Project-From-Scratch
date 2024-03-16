package main

import (
	"database/db"
	"fmt"
)
func main() {
	database,err := db.LoadDatabase("companyDatabase.bin");
	if (err != nil) {
		fmt.Println(err)
		return
	}
	// Print the database structure
	fmt.Println("----------------------")
	fmt.Println("Database Name:", database.Name)
	fmt.Println("----------------------")
	for _, table := range database.Tables {
		fmt.Println("Table Name:", table.Metadata.Name)
		fmt.Println("Table Columns:", table.Metadata.Columns)
		fmt.Println("Table Primary Keys:", table.Metadata.PrimaryKeys)
		fmt.Println("Table Foreign Keys:", table.Metadata.ForeignKeys)
		fmt.Println("Table Indexes:")
		for indexName, index := range table.IndexTable {
			fmt.Println("  - Index Name:", indexName)
			fmt.Println("    Index Columns:", index.Columns)
			fmt.Println("    Index Rows:", index.Rows)
			fmt.Println("    Index Unique:", index.Unique)
			fmt.Println("    Index Using:", index.Using)
			fmt.Println("    Index Comment:", index.Comment)
			fmt.Println("    Index Tablespace:", index.Tablespace)
			fmt.Println("    Index Include Columns:", index.Include)
			fmt.Println("    Index Predicate:", index.Predicate)
			fmt.Println("    Index Fill Factor:", index.FillFactor)
		}
		fmt.Println("Table Rows:")
		for _, row := range table.Metadata.Rows {
			fmt.Println("  - Row Data:", row.Data)
			fmt.Println("    Created At:", row.CreatedAt)
			fmt.Println("    Updated At:", row.UpdatedAt)
		}
		fmt.Println("----------------------")
	}
}