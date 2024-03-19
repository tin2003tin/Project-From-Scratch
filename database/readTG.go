package main

import (
	"database/db"
	"database/db/table"
	"fmt"
	"time"
)

func main() {
	database, _ := db.GetDataBase("G_type")
	mixedTable, err := database.GetTable("mixed")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mixedTable.ListColumn())
	fmt.Println("1")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "This is the most longest words",
		"bool": "bool_1", "int": "int_1", "float": "float_1",
		"byte": "", "date": "", "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("2")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": "bool_1", "int": "int_1", "float": "float_1",
		"byte": "", "date": "", "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("3")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": "int_1", "float": "float_1",
		"byte": "", "date": "", "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("4")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": "float_1",
		"byte": "", "date": "", "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("5")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": "", "date": "", "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("6")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": []byte("This is the longest words"), "date": "", "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("7")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": []byte("Hello"), "date": "", "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("8")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": []byte("Hello"), "date": time.Now(), "intArray": "", "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("9")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": []byte("Hello"), "date": time.Now(), "intArray": []int{1, 2, 3, 4, 5}, "stringArray": "",
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("10")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": []byte("Hello"), "date": time.Now(),
		"intArray": []int{1, 2, 3, 4, 5}, "stringArray": []string{"Banana", "Papaya", "Apple"},
		"floatArray": "", "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("11")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": []byte("Hello"), "date": time.Now(),
		"intArray": []int{1, 2, 3, 4, 5}, "stringArray": []string{"Banana", "Papaya", "Apple"},
		"floatArray": []float64{1.1, 10.99, 0.5, 10}, "varcharArray": "", "boolArray": ""})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("12")
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 1, "string": "string_1", "varchar": "test",
		"bool": true, "int": 1111, "float": 10.05,
		"byte": []byte("Hello"), "date": time.Now(),
		"intArray": []int{1, 2, 3, 4, 5}, "stringArray": []string{"Banana", "Papaya", "Apple"},
		"floatArray": []float64{1.1, 10.99, 0.5, 10}, "varcharArray": []string{"This is the longes Words", "hi"},
		"boolArray": []bool{true, false, true, true}})
	if err != nil {
		fmt.Println(err)
	}
	err = mixedTable.AddRow(map[string]interface{}{
		"id": 2, "string": "Tin5", "varchar": "character",
		"bool": false, "int": 52, "float": 30.05,
		"byte": []byte("Tin"), "date": time.Now(),
		"intArray": []int{10, 12, 31, 24, 15}, "stringArray": []string{"Google", "Facebook", "Microsoft"},
		"floatArray": []float64{41.1, 1, 2.2, 1.01}, "varcharArray": []string{"Book", "Pen", "Pencil", "Paper"},
		"boolArray": []bool{false, false, false, false}})
	if err != nil {
		fmt.Println(err)
	}
	table.PrintAsTable(mixedTable.Metadata.Columns[:8], mixedTable.Metadata.Rows)
	table.PrintAsTable(mixedTable.Metadata.Columns[8:], mixedTable.Metadata.Rows)
}
