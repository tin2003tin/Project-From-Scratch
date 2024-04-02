package main

import (
	db_constant "database/db/constant"
	"database/db/structure"
	"database/runtime"
	"fmt"
)

func main() {
	app := runtime.CreateApp()
	app.LoadMetaDatabase(db_constant.DATABASE_PATH)
	app.LoadTableMetadata()

	fmt.Println("Database")
	for i := range app.LoadedDatabase {
		fmt.Println("[", app.LoadedDatabase[i].Name, "]")
		for _, table := range app.LoadedDatabase[i].TableNames {
			fmt.Println("-", table)
			for _, col := range app.LoadedDatabase[i].Registry.Tables[table].Columns {
				fmt.Print(col.Name + " ")
			}
			fmt.Println()
		}
		fmt.Println()
		app.LoadedDatabase[i].Tables = []*structure.Table{}
		app.LoadedDatabase[i].Registry = &structure.TableRegistry{make(map[string]*structure.Table)}
	}
	fmt.Println()

	app.InitCompiler()

	fmt.Println("Action Table")
	for _, state := range app.SqlCompliler.Compiler.LRTable.States {
		fmt.Println(state)
	}
	fmt.Println()
	
	app.OpenServer("5432", HandleError)
}

func HandleError(err error) {
	fmt.Println(err)
	return
}
